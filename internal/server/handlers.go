package server

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func (s *Server) MapHandlers(e *echo.Echo) error {

	// Init repositories
	// Init useCases
	// Init handlers

	v1 := e.Group("/api/v1")
	health := v1.Group("/health")
	authGroup := v1.Group("/auth")

	authHttp.MapAuthRoutes(authGroup, authHandlers, mw)

	health.GET("", func(c echo.Context) error {
		return c.JSON(http.StatusOK, struct {
			Message string `json:"message"`
			Status  int    `json:"status"`
		}{
			Message: "Healthy",
			Status:  http.StatusOK,
		})
	})

	return nil
}
