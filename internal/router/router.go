package router

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func SetupRouter(e *echo.Echo) {
	// health check
	e.GET("/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "Health check OK")
	})
}
