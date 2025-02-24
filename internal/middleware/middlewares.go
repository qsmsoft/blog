package middleware

import (
	"github.com/qsmsoft/blog/config"
	"github.com/qsmsoft/blog/internal/auth"
	"github.com/qsmsoft/blog/internal/session"
	"github.com/qsmsoft/blog/pkg/logger"
)

type MiddlewareManager struct {
	sessUC  session.SessionUC
	authUC  auth.UseCase
	cfg     *config.Config
	origins []string
	logger  logger.Logger
}

func NewMiddlewareManager(sessUC session.SessionUC, authUC auth.UseCase, cfg *config.Config, origins []string, logger logger.Logger) *MiddlewareManager {
	return &MiddlewareManager{
		sessUC:  sessUC,
		authUC:  authUC,
		cfg:     cfg,
		origins: origins,
		logger:  logger,
	}
}
