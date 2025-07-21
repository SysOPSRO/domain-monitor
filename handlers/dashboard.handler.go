package handlers

import (
	"github.com/sysopsro/domain-monitor/views/dashboard"
	"github.com/labstack/echo/v4"
)

func HandlerRenderDashboard(c echo.Context) error {
	dashboard := dashboard.Dashboard()

	return View(c, dashboard)
}
