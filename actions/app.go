package actions

import (
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/buffalo/middleware"

	"github.com/BryanMoslo/QAReport/models"

	"github.com/gobuffalo/envy"
	"github.com/gobuffalo/packr"
)

// ENV is used to help switch settings based on where the
// application is being run. Default is "development".
var ENV = envy.Get("GO_ENV", "development")
var app *buffalo.App

// App is where all routes and middleware for buffalo
// should be defined. This is the nerve center of your
// application.
func App() *buffalo.App {
	if app == nil {
		app = buffalo.Automatic(buffalo.Options{
			Env:         ENV,
			SessionName: "_QAReport_session",
		})
		if ENV == "development" {
			app.Use(middleware.ParameterLogger)
		}

		app.Use(middleware.PopTransaction(models.DB))

		app.GET("/", HomeHandler)

		app.ServeFiles("/assets", packr.NewBox("../public/assets"))
		reports := app.Group("/reports")
		reports.Use(findReportMW)
		reports.GET("/", ReportsIndex)
		reports.GET("/{id}", ReportsShow)
		reports.DELETE("/{id}", ReportsDestroy)
		app.POST("/send_file", SendFile)
	}

	return app
}
