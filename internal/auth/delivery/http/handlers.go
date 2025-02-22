package http

import (
	"github.com/labstack/echo/v4"
	"github.com/qsmsoft/blog/config"
	"github.com/qsmsoft/blog/internal/auth"
	"github.com/qsmsoft/blog/internal/models"
	"github.com/qsmsoft/blog/pkg/httpErrors"
	"github.com/qsmsoft/blog/pkg/logger"
	"github.com/qsmsoft/blog/pkg/utils"
	"net/http"
)

// authHandlers is auth handlers
type authHandlers struct {
	cfg    *config.Config
	authUC auth.UseCase
	logger logger.Logger
}

// NewAuthHandlers is auth handlers constructor
func NewAuthHandlers(cfg *config.Config, authUC auth.UseCase, logger logger.Logger) auth.Handlers {
	return &authHandlers{
		cfg:    cfg,
		authUC: authUC,
		logger: logger,
	}
}

func (h *authHandlers) Register() echo.HandlerFunc {
	return func(c echo.Context) error {
		user := &models.User{}
		if err := utils.ReadRequest(c, user); err != nil {
			utils.LogResponseError(c, h.logger, err)
			return c.JSON(httpErrors.ErrorResponse(err))
		}

		createUser, err := h.authUC.Register(c, user)
		if err != nil {
			utils.LogResponseError(c, h.logger, err)
			return c.JSON(httpErrors.ErrorResponse(err))
		}

		sess, err := h.sessUC.CreateSession(c, &models.Session{
			UserID: createUser.User.ID,
		}, h.cfg.Session.Expire)
		if err != nil {
			utils.LogResponseError(c, h.logger, err)
			return c.JSON(httpErrors.ErrorResponse(err))
		}

		c.SetCookie(utils.CreateSessionCookie(h.cfg, sess))

		return c.JSON(http.StatusCreated, createUser)
	}
}

func (h *authHandlers) Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		return nil
	}
}

func (h *authHandlers) Logout() echo.HandlerFunc {
	return func(c echo.Context) error {
		return nil
	}
}

func (h *authHandlers) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		return nil
	}
}

func (h *authHandlers) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		return nil
	}
}

func (h *authHandlers) GetUserByID() echo.HandlerFunc {
	return func(c echo.Context) error {
		return nil
	}
}

func (h *authHandlers) FindByName() echo.HandlerFunc {
	return func(c echo.Context) error {
		return nil
	}
}

func (h *authHandlers) GetUsers() echo.HandlerFunc {
	return func(c echo.Context) error {
		return nil
	}
}

func (h *authHandlers) GetMe() echo.HandlerFunc {
	return func(c echo.Context) error {
		return nil
	}
}

func (h *authHandlers) UploadAvatar() echo.HandlerFunc {
	return func(c echo.Context) error {
		return nil
	}
}

func (h *authHandlers) GetCSRFToken() echo.HandlerFunc {
	return func(c echo.Context) error {
		return nil
	}
}
