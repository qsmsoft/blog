package app

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/qsmsoft/blog/config"
	"github.com/qsmsoft/blog/internal/routes"
)

func SetupApp(cfg *config.Config) *echo.Echo {
	e := echo.New()

	// Middleware
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "\033[1;34m${time_rfc3339}\033[0m method=\033[1;32m${method}\033[0m, uri=\033[1;36m${uri}\033[0m, status=\033[1;33m${status}\033[0m\n",
	}))

	// Routes
	routes.RegisterRoutes(e, cfg)

	return e
}
