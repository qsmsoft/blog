package http

import (
	"github.com/labstack/echo/v4"
	"github.com/qsmsoft/blog/internal/auth"
	"github.com/qsmsoft/blog/internal/middleware"
)

func MapAuthRoutes(authGroup *echo.Group, h auth.Handlers, mw *middleware.MiddlewareManager) {
	authGroup.POST("/register", h.Register())
	authGroup.POST("/login", h.Login())
	authGroup.POST("/logout", h.Logout())
	authGroup.GET("/find", h.FindByName())
	authGroup.GET("/all", h.GetUsers())
	authGroup.GET("/:id", h.GetUserByID())
	// authGroup.Use(middleware.AuthJWTMiddleware(authUC, cfg))
	authGroup.Use(mw.AuthSessionMiddleware)
	authGroup.GET("/me", h.GetMe())
	authGroup.GET("/token", h.GetCSRFToken())
	authGroup.POST("/:id/avatar", h.UploadAvatar(), mw.CSRF)
	authGroup.PUT("/:id", h.Update(), mw.OwnerOrAdminMiddleware(), mw.CSRF)
	authGroup.DELETE("/:id", h.Delete(), mw.CSRF, mw.RoleBasedAuthMiddleware([]string{"admin"}))
}
