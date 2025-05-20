package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	// "github.com/labstack/echo/v5"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/plugins/jsvm"
	"github.com/pocketbase/pocketbase/plugins/migratecmd"
	"github.com/pocketbase/pocketbase/tools/hook"

	"database/sql"
	"fmt"

	"github.com/mattn/go-sqlite3"
	"github.com/pocketbase/dbx"
)

func init() {

	// initialize default PRAGMAs for each new connection
	sql.Register("pb_sqlite3",
		&sqlite3.SQLiteDriver{
			ConnectHook: func(conn *sqlite3.SQLiteConn) error {
				_, err := conn.Exec(`
                    PRAGMA busy_timeout       = 10000;
                    PRAGMA journal_mode       = WAL;
                    PRAGMA journal_size_limit = 200000000;
                    PRAGMA synchronous        = NORMAL;
                    PRAGMA foreign_keys       = ON;
                    PRAGMA temp_store         = MEMORY;
                    PRAGMA cache_size         = -16000;
                `, nil)

				return err
			},
		},
	)

	dbx.BuilderFuncMap["pb_sqlite3"] = dbx.BuilderFuncMap["sqlite3"]
}

func main() {

	app := pocketbase.NewWithConfig(pocketbase.Config{
		// DefaultDataDir: path,

		DBConnect: func(dbPath string) (*dbx.DB, error) {
			// key := "secretkey" // replace with your actual key
			key := GetEnvOrDefault("DB_KEY", "secretkey")
			dbname := fmt.Sprintf("%s?_cipher=sqlcipher&_legacy=4&_key=%s", dbPath, key)
			log.Println("--- db start --- " + dbname)
			return dbx.Open("pb_sqlite3", dbname)
		},
	})

	// ---------------------------------------------------------------
	// Optional plugin flags:
	// ---------------------------------------------------------------

	var hooksDir string
	app.RootCmd.PersistentFlags().StringVar(
		&hooksDir,
		"hooksDir",
		"",
		"the directory with the JS app hooks",
	)

	var hooksWatch bool
	app.RootCmd.PersistentFlags().BoolVar(
		&hooksWatch,
		"hooksWatch",
		true,
		"auto restart the app on pb_hooks file change; it has no effect on Windows",
	)

	var hooksPool int
	app.RootCmd.PersistentFlags().IntVar(
		&hooksPool,
		"hooksPool",
		15,
		"the total prewarm goja.Runtime instances for the JS app hooks execution",
	)

	var migrationsDir string
	app.RootCmd.PersistentFlags().StringVar(
		&migrationsDir,
		"migrationsDir",
		"",
		"the directory with the user defined migrations",
	)

	var automigrate bool
	app.RootCmd.PersistentFlags().BoolVar(
		&automigrate,
		"automigrate",
		true,
		"enable/disable auto migrations",
	)

	var publicDir string
	app.RootCmd.PersistentFlags().StringVar(
		&publicDir,
		"publicDir",
		defaultPublicDir(),
		"the directory to serve static files",
	)

	var indexFallback bool
	app.RootCmd.PersistentFlags().BoolVar(
		&indexFallback,
		"indexFallback",
		true,
		"fallback the request to index.html on missing static path, e.g. when pretty urls are used with SPA",
	)

	app.RootCmd.ParseFlags(os.Args[1:])

	// ---------------------------------------------------------------
	// Plugins and hooks:
	// ---------------------------------------------------------------

	// load jsvm (pb_hooks and pb_migrations)
	jsvm.MustRegister(app, jsvm.Config{
		MigrationsDir: migrationsDir,
		HooksDir:      hooksDir,
		HooksWatch:    hooksWatch,
		HooksPoolSize: hooksPool,
	})

	// migrate command (with js templates)
	migratecmd.MustRegister(app, app.RootCmd, migratecmd.Config{
		TemplateLang: migratecmd.TemplateLangJS,
		Automigrate:  automigrate,
		Dir:          migrationsDir,
	})

	// static route to serves files from the provided public dir
	// (if publicDir exists and the route path is not already defined)
	app.OnServe().Bind(&hook.Handler[*core.ServeEvent]{
		Func: func(e *core.ServeEvent) error {
			if !e.Router.HasRoute(http.MethodGet, "/{path...}") {
				e.Router.GET("/{path...}", apis.Static(os.DirFS(publicDir), indexFallback))
			}

			e.Router.GET("/hello/{name}", func(e *core.RequestEvent) error {

				name := e.Request.PathValue("name")
				envs := "envirenonments:\n"

				for _, env := range os.Environ() {
					envs += env + "\n"
				}

				return e.String(http.StatusOK, "Hello "+name+"\n"+envs)
			})

			e.Router.GET("/orders1", func(e *core.RequestEvent) error {
				// Auth kontrolü istiyorsan:

				// "order_infos" koleksiyonundan verileri çek
				records, err := app.FindRecordsByFilter("order_infos", "", "-created", 100, 0)
				if err != nil {
					return apis.NewBadRequestError("Liste alınamadı", err)
				}

				errs := app.ExpandRecords(records, []string{"customer", "address", "payments.order", "order_items.order"}, nil)
				if len(errs) > 0 {
					return fmt.Errorf("failed to expand: %v", errs)
				}

				return e.JSON(http.StatusOK, records)
			})

			e.Router.GET("/orders2", func(e *core.RequestEvent) error {
				// SQL sorgusunu kur
				query := app.DB().
					Select("order_infos.*").
					From("order_infos").
					LeftJoin("customers", dbx.NewExp("customers.id = order_infos.customer")).
					LeftJoin("address", dbx.NewExp("address.id = order_infos.address")).
					LeftJoin("payments", dbx.NewExp("payments.order = order_infos.id"))

				var result []map[string]any
				err := query.All(&result)
				if err != nil {
					return apis.NewBadRequestError("Sorgu hatası", err)
				}

				return e.JSON(http.StatusOK, result)
			})

			return e.Next()
		},
		Priority: 999, // execute as latest as possible to allow users to provide their own route
	})

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}

// the default pb_public dir location is relative to the executable
func defaultPublicDir() string {
	if strings.HasPrefix(os.Args[0], os.TempDir()) {
		// most likely ran with go run
		return "./pb_public"
	}

	return filepath.Join(os.Args[0], "../pb_public")
}

func GetEnvOrDefault(key string, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
