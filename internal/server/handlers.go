package server

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func (s *Server) MapHandlers() error {
	v1 := s.echo.Group("/api/v1")
	health := v1.Group("/health")
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
