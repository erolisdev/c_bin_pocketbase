package main

import (
	"fmt"
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

	// "c_bin_pocketbase/generated"

	_ "c_bin_pocketbase/migrations" // migrations folder
	"c_bin_pocketbase/models"
	"c_bin_pocketbase/services"

	"github.com/mattn/go-sqlite3"
	"github.com/pocketbase/dbx"
)

// var _ core.RecordProxy = (*generated.LiveOrders)(nil)

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
			return dbx.Open("pb_sqlite3", dbname)
			// log.Println("--- db start --- test dbname: ")
			// return dbx.Open("pb_sqlite3", dbPath)
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

	migratecmd.MustRegister(app, app.RootCmd, migratecmd.Config{
		// enable auto creation of migration files when making collection changes in the Dashboard
		// (the isGoRun check is to enable it only during development)
		Automigrate: true,
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
				// SQL sorgusunu kur
				var orders []OrderResponse

				// 1. Tüm order_infos kayıtlarını çek
				// var rawOrders []models.Record

				rawOrders := []*core.Record{}

				err := app.RecordQuery("order_infos").
					Select("order_infos.*").
					OrderBy("created DESC").
					Limit(100).
					All(&rawOrders)

				if err != nil {
					return apis.NewBadRequestError("Order verileri alınamadı "+err.Error(), err)
				}

				// 2. Her order için ilişkili verileri ekle
				for _, rec := range rawOrders {
					order := OrderResponse{
						Id:         rec.GetString("id"),
						Created:    rec.GetString("created"),
						Updated:    rec.GetString("updated"),
						Status:     rec.GetString("status"),
						OrderNote:  rec.GetString("orderNote"),
						OrderTotal: rec.GetInt("orderTotal"),
					}

					// Customer
					customerRec, _ := app.FindRecordById("customers", rec.GetString("customer"))
					if customerRec != nil {
						order.Customer = customerRec.PublicExport()
					}

					// Address
					addressRec, _ := app.FindRecordById("shipping_address", rec.GetString("address"))
					if addressRec != nil {
						order.ShippingAddress = addressRec.PublicExport()
					}

					payments, _ := app.FindAllRecords("payments",
						dbx.NewExp("[order] = {:orderId}", dbx.Params{"orderId": rec.Id}),
					)

					for _, p := range payments {
						order.Payments = append(order.Payments, p.PublicExport())
					}

					// Order Items

					items, errItem := app.FindAllRecords("order_items",
						dbx.NewExp("[order] = {:orderId}", dbx.Params{"orderId": rec.Id}),
					)
					if errItem != nil {
						log.Println("Error: ", errItem)
					}
					// log.Println("items: ", items)
					for _, i := range items {
						log.Println("item: ", i.PublicExport())
						order.OrderItems = append(order.OrderItems, i.PublicExport())
					}

					orders = append(orders, order)
				}

				return e.JSON(http.StatusOK, orders)

			})

			e.Router.GET("/orders2", func(e *core.RequestEvent) error {
				// SQL sorgusunu kur
				var orders []OrderResponse

				// 1. Tüm order_infos kayıtlarını çek
				// var rawOrders []models.Record

				records := []*core.Record{}

				err := app.RecordQuery("order_infos").
					Select("order_infos.*").
					OrderBy("created DESC").
					Limit(100).
					All(&records)

				if err != nil {
					return apis.NewBadRequestError("Order verileri alınamadı "+err.Error(), err)
				}

				errs := app.ExpandRecords(records, []string{
					"customer",              // single relation
					"address",               // single relation
					"payments_via_order",    // multiple relation
					"order_items_via_order", // multiple relation
				}, nil)

				if len(errs) > 0 {
					return apis.NewBadRequestError("Order verileri alınamadı: %v", errs)
				}

				// 2. Her order için ilişkili verileri ekle
				for _, rec := range records {
					order := OrderResponse{
						Id:         rec.GetString("id"),
						Created:    rec.GetString("created"),
						Updated:    rec.GetString("updated"),
						Status:     rec.GetString("status"),
						OrderNote:  rec.GetString("orderNote"),
						OrderTotal: rec.GetInt("orderTotal"),
					}

					// log.Println("record: : ", rec.PublicExport())

					customer := rec.ExpandedOne("customer")
					order.Customer = customer.PublicExport()

					address := rec.ExpandedOne("address")
					order.ShippingAddress = address.PublicExport()

					payments := rec.ExpandedAll("payments_via_order")
					for _, p := range payments {
						order.Payments = append(order.Payments, p.PublicExport())
					}

					items := rec.ExpandedAll("order_items_via_order")
					for _, i := range items {
						order.OrderItems = append(order.OrderItems, i.PublicExport())
					}

					orders = append(orders, order)
				}

				return e.JSON(http.StatusOK, orders)

			})

			// ================
			e.Router.POST("/createOrder", func(e *core.RequestEvent) error {

				var order models.OrderModel
				if err := e.BindBody(&order); err != nil {
					return apis.NewBadRequestError("Invalid request body", err)
				}

				// Verify the order
				orderData, err := services.VerifyOrder(app, order)

				if err != nil {
					log.Printf("=== Order verification failed: %v. Order data: %+v ===", err, order)
					return apis.NewBadRequestError("Order verification failed: "+err.Error(), nil)
				}

				// Save order
				orderNumber, err := services.SaveOrder(app, orderData)

				if err != nil {
					// log.Printf("=== Order save failed: %v. Order number: %+v ===", err, orderNumber)
					return apis.NewBadRequestError("Order save failed: "+err.Error(), nil)
				}

				return e.JSON(http.StatusCreated, map[string]interface{}{
					"success":      true,
					"total":        order.OrderData.Total,
					"order_number": orderNumber,
				})

			})
			// .Bind(apis.RequireAuth()) // require auth and admin role

			e.Router.GET("/list-images/:folder", func(e *core.RequestEvent) error {
				return e.Next()
			})

			return e.Next()
		},
		Priority: 999, // execute as latest as possible to allow users to provide their own route
	})

	// eronorhooks.OrderNumberCreate(app)

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}

func PtrValue[T any](ptr *T) T {
	if ptr != nil {
		return *ptr
	}
	var zero T
	return zero
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

type OrderResponse struct {
	Id              string           `json:"id"`
	Created         string           `json:"created"`
	Updated         string           `json:"updated"`
	Status          string           `json:"status"`
	OrderNote       string           `json:"orderNote"`
	OrderTotal      int              `json:"orderTotal"`
	Customer        map[string]any   `json:"customer"`
	ShippingAddress map[string]any   `json:"shipping_address"`
	Payments        []map[string]any `json:"payments"`
	OrderItems      []map[string]any `json:"order_items"`
}

// burasi dogrulama icin kullanilcak
type ClientOrder struct {
	TotalPriceHT  float64              `json:"total_price_ht"`
	TotalPriceTTC float64              `json:"total_price_ttc"`
	Products      []ClientOrderProduct `json:"products"`
}

type ClientOrderProduct struct {
	Id            string          `json:"product_id"`
	PriceHT       string          `json:"price_ht"`
	PriceTTC      string          `json:"price_ttc"`
	Quantity      int             `json:"quantity"`
	TotalPriceHT  float64         `json:"total_price_ht"`
	TotalPriceTTC float64         `json:"total_price_ttc"`
	Options       *[]ClientOption `json:"options"`
}

type ClientOption struct {
	Id            string              `json:"option_id"`
	FreeCount     string              `json:"free_count"`
	MaxCount      string              `json:"max_count"`
	OptionValaues []ClientOptionValue `json:"option_values"`
}

type ClientOptionValue struct {
	Id             string  `json:"option_value_id"`
	PriceHT        string  `json:"price_ht"`
	PriceTTC       string  `json:"price_ttc"`
	Quantity       int     `json:"quantity"`
	TotalPriceHT   float64 `json:"total_price_ht"`
	TotalPriceWTax float64 `json:"total_price_ttc"`
}
