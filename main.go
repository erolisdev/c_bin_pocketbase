package main

import (
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
)

func main() {
	app := pocketbase.New()

	// serves static files from the provided public dir (if exists)
	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		e.Router.GET("/*", apis.StaticDirectoryHandler(os.DirFS("./pb_public"), false))
		e.Router.GET("/old/hello/:name", func(c echo.Context) error {
			name := c.PathParam("name")
			return c.JSON(http.StatusOK, map[string]string{"message": "Hello " + name})
		} /* optional middlewares */)
		return nil
	})

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}

	// app.OnServe().BindFunc(func(se *core.ServeEvent) error {
	//     // serves static files from the provided public dir (if exists)
	//     se.Router.GET("/{path...}", apis.Static(os.DirFS("./pb_public"), false))

	//     return se.Next()
	// })

	// if err := app.Start(); err != nil {
	//     log.Fatal(err)
	// }

}
