package usecase

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/qsmsoft/blog/config"
	"github.com/qsmsoft/blog/internal/auth"
	"github.com/qsmsoft/blog/internal/models"
	"github.com/qsmsoft/blog/pkg/httpErrors"
	"github.com/qsmsoft/blog/pkg/logger"
	"github.com/qsmsoft/blog/pkg/utils"
	"net/http"
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
	existsUser, err := u.authRepo.FindByEmail(ctx, user)
	if existsUser != nil || err == nil {
		return nil, httpErrors.NewRestErrorWithMessage(http.StatusBadRequest, httpErrors.ErrEmailAlreadyExists, nil)
	}

	if err = user.PrepareCreate(); err != nil {
		return nil, httpErrors.NewBadRequestError(errors.Wrap(err, "authUC.Register.PrepareCreate"))
	}

	createdUser, err := u.authRepo.Register(ctx, user)
	if err != nil {
		return nil, err
	}

	createdUser.SanitizePassword()

	token, err := utils.GenerateJWTToken(createdUser, u.cfg)
	if err != nil {
		return nil, httpErrors.NewInternalServerError(errors.Wrap(err, "authUC.Register.GenerateJWTToken"))
	}

	return &models.UserWithToken{
		User:  createdUser,
		Token: token,
	}, nil
}

func (u *authUC) Update(ctx context.Context, user *models.User) (*models.User, error) {
	if err := user.PrepareUpdate(); err != nil {
		return nil, httpErrors.NewBadRequestError(errors.Wrap(err, "authUC.Update.PrepareUpdate"))
	}

	updatedUser, err := u.authRepo.Update(ctx, user)
	if err != nil {
		return nil, err
	}

	updatedUser.SanitizePassword()

	if err = u.redisRepo.DeleteUserCtx(ctx, u.GenerateUserKey(user.ID.String())); err != nil {
		u.logger.Errorf("AuthUC.Update.DeleteUserCtx: %v", err)
	}

	updatedUser.SanitizePassword()

	return updatedUser, nil
}

func (u *authUC) Delete(ctx context.Context, ID uuid.UUID) error {
	if err := u.authRepo.Delete(ctx, ID); err != nil {
		return err
	}

	if err := u.redisRepo.DeleteUserCtx(ctx, u.GenerateUserKey(ID.String())); err != nil {
		u.logger.Errorf("AuthUC.Delete.DeleteUserCtx: %s", err)
	}

	return nil
}

func (u *authUC) GetByID(ctx context.Context, ID uuid.UUID) (*models.User, error) {
	cachedUser, err := u.redisRepo.GetByIDCtx(ctx, u.GenerateUserKey(ID.String()))
	if err != nil {
		u.logger.Errorf("authUC.GetByID.GetByIDCtx: %v", err)
	}
	if cachedUser != nil {
		return cachedUser, nil
	}

	user, err := u.authRepo.GetByID(ctx, ID)
	if err != nil {
		return nil, err
	}

	if err = u.redisRepo.SetUserCtx(ctx, u.GenerateUserKey(ID.String()), cacheDuration, user); err != nil {
		u.logger.Errorf("authUC.GetByID.SetUserCtx: %v", err)
	}

	user.SanitizePassword()

	return user, nil
}

func (u *authUC) FindByName(ctx context.Context, name string, query *utils.PaginationQuery) (*models.UsersList, error) {
	return u.authRepo.FindByName(ctx, name, query)
}

func (u *authUC) GetUsers(ctx context.Context, query *utils.PaginationQuery) (*models.UsersList, error) {
	return u.authRepo.GetUsers(ctx, query)
}

func (u *authUC) Login(ctx context.Context, user *models.User) (*models.UserWithToken, error) {
	foundUser, err := u.authRepo.FindByEmail(ctx, user)
	if err != nil {
		return nil, err
	}

	if err = foundUser.ValidatePassword(user.Password); err != nil {
		return nil, httpErrors.NewUnauthorizedError(errors.Wrap(err, "authUC.GetUsers.ValidatePassword"))
	}

	foundUser.SanitizePassword()

	token, err := utils.GenerateJWTToken(foundUser, u.cfg)
	if err != nil {
		return nil, httpErrors.NewInternalServerError(errors.Wrap(err, "authUC.GetUsers.GenerateJWTToken"))
	}

	return &models.UserWithToken{
		User:  foundUser,
		Token: token,
	}, nil
}

func (u *authUC) UploadAvatar(ctx context.Context, ID uuid.UUID, file models.UploadInput) (*models.User, error) {
	uploadInfo, err := u.awsRepo.PutObject(ctx, file)
	if err != nil {
		return nil, httpErrors.NewInternalServerError(errors.Wrap(err, "authUC.UploadAvatar.PutObject"))
	}

	avatarURL := u.generateAWSMinioURL(file.BucketName, uploadInfo.Key)

	updatedUser, err := u.authRepo.Update(ctx, &models.User{
		ID:     ID,
		Avatar: &avatarURL,
	})
	if err != nil {
		return nil, err
	}

	updatedUser.SanitizePassword()

	return updatedUser, nil
}

func (u *authUC) GenerateUserKey(ID string) string {
	return fmt.Sprintf("%s: %s", basePrefix, ID)
}

func (u *authUC) generateAWSMinioURL(bucket string, key string) string {
	return fmt.Sprintf("%s/minio/%s/%s", u.cfg.AWS.MinioEndpoint, bucket, key)
}
