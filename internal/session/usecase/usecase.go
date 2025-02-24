package usecase

import (
	"context"
	"github.com/qsmsoft/blog/config"
	"github.com/qsmsoft/blog/internal/models"
	"github.com/qsmsoft/blog/internal/session"
)

type sessionUC struct {
	sessionRepo session.SessRepository
	cfg         *config.Config
}

func NewSessionUseCase(sessionRepo session.SessRepository, cfg *config.Config) session.SessionUC {
	return &sessionUC{sessionRepo: sessionRepo, cfg: cfg}
}

func (u *sessionUC) CreateSession(ctx context.Context, session *models.Session, expire int) (string, error) {
	return u.sessionRepo.CreateSession(ctx, session, expire)
}

func (u *sessionUC) DeleteByID(ctx context.Context, sessionID string) error {
	return u.sessionRepo.DeleteByID(ctx, sessionID)
}

func (u *sessionUC) GetSessionByID(ctx context.Context, sessionID string) (*models.Session, error) {
	return u.sessionRepo.GetSessionByID(ctx, sessionID)
}
