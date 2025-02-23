package auth

import (
	"context"
	"github.com/google/uuid"
	"github.com/qsmsoft/blog/internal/models"
	"github.com/qsmsoft/blog/pkg/utils"
)

// UseCase interface fot Auth repository
type UseCase interface {
	Register(ctx context.Context, user *models.User) (*models.UserWithToken, error)
	Update(ctx context.Context, user *models.User) (*models.User, error)
	Delete(ctx context.Context, ID uuid.UUID) error
	GetByID(ctx context.Context, ID uuid.UUID) (*models.User, error)
	FindByName(ctx context.Context, name string, query *utils.PaginationQuery) (*models.UsersList, error)
	GetUsers(ctx context.Context, pq *utils.PaginationQuery) (*models.UsersList, error)
	Login(ctx context.Context, user *models.User) (*models.UserWithToken, error)
	UploadAvatar(ctx context.Context, userID uuid.UUID, file models.UploadInput) (*models.User, error)
}
