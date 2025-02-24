package session

import (
	"context"
	"github.com/qsmsoft/blog/internal/models"
)

type SessionUC interface {
	CreateSession(ctx context.Context, session *models.Session, expire int) (string, error)
	GetSessionByID(ctx context.Context, sessionID string) (*models.Session, error)
	DeleteByID(ctx context.Context, sessionID string) error
}
