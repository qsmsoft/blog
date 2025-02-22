package utils

import (
	"github.com/labstack/echo/v4"
	"github.com/qsmsoft/blog/pkg/logger"
)

// GetRequestID returns request id from echo context
func GetRequestID(c echo.Context) string {
	return c.Response().Header().Get(echo.HeaderXRequestID)
}

// GetConfigPath returns config path for local or docker
func GetConfigPath(configPath string) string {
	if configPath == "docker" {
		return "./config/config-docker"
	}
	return "./config/config-local"
}

// GetIPAddress returns user ip address
func GetIPAddress(c echo.Context) string {
	return c.Request().RemoteAddr
}

// ReadRequest returns request body and validate
func ReadRequest(ctx echo.Context, request interface{}) error {
	if err := ctx.Bind(request); err != nil {
		return err
	}
	return validate.StructCtx(ctx.Request().Context(), request)
}

// LogResponseError returns error response with logging error for echo context
func LogResponseError(ctx echo.Context, logger logger.Logger, err error) {
	logger.Errorf(
		"ErrResponseWithLog, RequestID: %s, IPAddress: %s, Error: %s",
		GetRequestID(ctx),
		GetIPAddress(ctx),
		err,
	)
}
