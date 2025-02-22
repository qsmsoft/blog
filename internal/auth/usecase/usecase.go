package usecase

import (
	"context"
	"github.com/google/uuid"
	"github.com/qsmsoft/blog/config"
	"github.com/qsmsoft/blog/internal/auth"
	"github.com/qsmsoft/blog/internal/models"
	"github.com/qsmsoft/blog/pkg/logger"
	"github.com/qsmsoft/blog/pkg/utils"
)

const (
	basePrefix    = "api-auth:"
	cacheDuration = 3600
)

type authUC struct {
	cfg       *config.Config
	authRepo  auth.PostgresRepository
	redisRepo auth.RedisRepository
	awsRepo   auth.AWSRepository
	logger    logger.Logger
}

func NewAuthUseCase(cfg *config.Config, authRepo auth.PostgresRepository, redisRepo auth.RedisRepository, awsRepo auth.AWSRepository, logger logger.Logger) auth.UseCase {
	return &authUC{
		cfg:       cfg,
		authRepo:  authRepo,
		redisRepo: redisRepo,
		awsRepo:   awsRepo,
		logger:    logger,
	}
}

func (u *authUC) Register(ctx context.Context, user *models.User) (*models.UserWithToken, error) {
	return nil, nil
}

func (u *authUC) Login(ctx context.Context, user *models.User) (*models.UserWithToken, error) {
	return nil, nil
}

func (u *authUC) Update(ctx context.Context, user *models.User) (*models.User, error) {
	return nil, nil
}

func (u *authUC) Delete(ctx context.Context, ID uuid.UUID) error {
	return nil
}

func (u *authUC) GetByID(ctx context.Context, ID uuid.UUID) (*models.User, error) {
	return nil, nil
}

func (u *authUC) FindByName(ctx context.Context, name string, query *utils.PaginationQuery) (*models.UsersList, error) {
	return nil, nil
}

func (u *authUC) GetUsers(ctx context.Context, query *utils.PaginationQuery) (*models.UsersList, error) {
	return nil, nil
}

func (u *authUC) UploadAvatar(ctx context.Context, userID uuid.UUID, file models.UploadInput) (*models.User, error) {
	return nil, nil
}
