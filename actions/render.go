package actions

import (
	"fmt"

	"github.com/gobuffalo/buffalo/render"
	"github.com/gobuffalo/packr"
)

var r *render.Engine

func init() {
	r = render.New(render.Options{
		// HTML layout to be used for all HTML requests:
		HTMLLayout: "application.html",

		// Box containing all of the templates:
		TemplatesBox: packr.NewBox("../templates"),

		// Add template helpers here:
		Helpers: map[string]interface{}{
			"floatToString": func(val float64) string {
				return fmt.Sprintf("%.2f", val)
			},
			"milesFor": func(val float64) string {
				return fmt.Sprintf("%.2f", val*0.621371)
			},
		},
	})
}
