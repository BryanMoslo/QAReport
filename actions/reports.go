package actions

import (
	"github.com/BryanMoslo/QAReport/models"
	"github.com/gobuffalo/buffalo"
	"github.com/markbates/pop"
	"github.com/pkg/errors"
)

func findReportMW(h buffalo.Handler) buffalo.Handler {
	return func(c buffalo.Context) error {
		id := c.Param("id")

		if id != "" {
			report := &models.Report{}
			tx := c.Value("tx").(*pop.Connection)
			err := tx.Find(report, id)

			if err != nil {
				return c.Error(404, errors.WithStack(err))
			}

			c.Set("report", report)
		}

		return h(c)
	}
}

// ReportsIndex default implementation.
func ReportsIndex(c buffalo.Context) error {
	reports := &models.Reports{}
	tx := c.Get("tx").(*pop.Connection)
	err := tx.All(reports)
	if err != nil {
		return c.Error(404, errors.WithStack(err))
	}

	c.Set("reports", reports)

	return c.Render(200, r.HTML("reports/index.html"))
}

// ReportsShow default implementation.
func ReportsShow(c buffalo.Context) error {
	return c.Render(200, r.HTML("reports/show.html"))
}

func ReportsDestroy(c buffalo.Context) error {
	tx := c.Get("tx").(*pop.Connection)
	report := c.Get("report").(*models.Report)

	err := tx.Destroy(report)
	if err != nil {
		return errors.WithStack(err)
	}

	return c.Redirect(301, "/reports")
}
