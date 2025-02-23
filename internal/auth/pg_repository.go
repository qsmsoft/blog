package auth

import (
	"context"
	"github.com/google/uuid"
	"github.com/qsmsoft/blog/internal/models"
	"github.com/qsmsoft/blog/pkg/utils"
)

type PostgresRepository interface {
	Register(ctx context.Context, user *models.User) (*models.User, error)
	Update(ctx context.Context, user *models.User) (*models.User, error)
	Delete(ctx context.Context, userID uuid.UUID) error
	GetByID(ctx context.Context, userID uuid.UUID) (*models.User, error)
	FindByName(ctx context.Context, name string, query *utils.PaginationQuery) (*models.UsersList, error)
	GetUsers(ctx context.Context, query *utils.PaginationQuery) (*models.UsersList, error)
	FindByEmail(ctx context.Context, user *models.User) (*models.User, error)
}
