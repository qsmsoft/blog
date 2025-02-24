package repository

import (
	"context"
	"database/sql"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/qsmsoft/blog/internal/auth"
	"github.com/qsmsoft/blog/internal/models"
	"github.com/qsmsoft/blog/pkg/utils"
	"log"
)

type authRepo struct {
	db *sqlx.DB
}

func NewAuthRepository(db *sqlx.DB) auth.PostgresRepository {
	return &authRepo{db: db}
}

func (r *authRepo) Register(ctx context.Context, user *models.User) (*models.User, error) {
	u := &models.User{}
	if err := r.db.QueryRowxContext(ctx, createUserQuery, &user.FirstName, &user.LastName,
		&user.Email, &user.Password, &user.Role, &user.Avatar).StructScan(u); err != nil {
		return nil, errors.Wrap(err, "authRepo.Register.StructScan")
	}

	return u, nil
}

func (r *authRepo) Update(ctx context.Context, user *models.User) (*models.User, error) {
	u := &models.User{}
	if err := r.db.GetContext(ctx, u, updateUserQuery, &user.FirstName, &user.LastName, &user.Email, &user.Role, &user.Avatar, &user.ID); err != nil {
		return nil, errors.Wrap(err, "authRepo.Update.GetContext")
	}

	return u, nil
}

func (r *authRepo) Delete(ctx context.Context, ID uuid.UUID) error {
	result, err := r.db.ExecContext(ctx, deleteUserQuery, ID)
	if err != nil {
		return errors.Wrap(err, "authRepo.Delete.ExecContext")
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return errors.Wrap(err, "authRepo.Delete.RowsAffected")
	}
	if rowsAffected == 0 {
		return errors.Wrap(sql.ErrNoRows, "authRepo.Delete.rowsAffected")
	}

	return nil
}

func (r *authRepo) GetByID(ctx context.Context, ID uuid.UUID) (*models.User, error) {
	user := &models.User{}
	if err := r.db.QueryRowxContext(ctx, getUserQuery, ID).StructScan(user); err != nil {
		return nil, errors.Wrap(err, "authRepo.GetByID.QueryRowContext")
	}

	return user, nil
}

func (r *authRepo) FindByName(ctx context.Context, name string, query *utils.PaginationQuery) (*models.UsersList, error) {
	var totalCount int
	if err := r.db.GetContext(ctx, &totalCount, getTotalCount, name); err != nil {
		return nil, errors.Wrap(err, "authRepo.FindByName.GetContext.totalCount")
	}

	if totalCount == 0 {
		return &models.UsersList{
			TotalCount: totalCount,
			TotalPages: utils.GetTotalPages(totalCount, query.GetSize()),
			Page:       query.GetPage(),
			Size:       query.GetSize(),
			HasMore:    utils.GetHasMore(query.GetPage(), totalCount, query.GetSize()),
			Users:      make([]*models.User, 0),
		}, nil
	}

	rows, err := r.db.QueryxContext(ctx, findUsers, name, query.GetOffset(), query.GetLimit())
	if err != nil {
		return nil, errors.Wrap(err, "authRepo.FindByName.QueryContext")
	}
	defer func(rows *sqlx.Rows) {
		err := rows.Close()
		if err != nil {
			log.Fatalf("Failed to close rows: %v", err)
		}
	}(rows)

	var users = make([]*models.User, 0, query.GetSize())
	for rows.Next() {
		var user models.User
		if err = rows.StructScan(&user); err != nil {
			return nil, errors.Wrap(err, "authRepo.FindByName.StructScan")
		}
		users = append(users, &user)
	}

	if err = rows.Err(); err != nil {
		return nil, errors.Wrap(err, "authRepo.FindByName.rows.Err")
	}

	return &models.UsersList{
		TotalCount: totalCount,
		TotalPages: utils.GetTotalPages(totalCount, query.GetSize()),
		Page:       query.GetPage(),
		Size:       query.GetSize(),
		HasMore:    utils.GetHasMore(query.GetPage(), totalCount, query.GetSize()),
		Users:      users,
	}, nil
}

func (r *authRepo) GetUsers(ctx context.Context, query *utils.PaginationQuery) (*models.UsersList, error) {
	var totalCount int
	if err := r.db.GetContext(ctx, &totalCount, getTotal); err != nil {
		return nil, errors.Wrap(err, "authRepo.GetUsers.GetContext.totalCount")
	}

	if totalCount == 0 {
		return &models.UsersList{
			TotalCount: totalCount,
			TotalPages: utils.GetTotalPages(totalCount, query.GetSize()),
			Page:       query.GetPage(),
			Size:       query.GetSize(),
			HasMore:    utils.GetHasMore(query.GetPage(), totalCount, query.GetSize()),
			Users:      make([]*models.User, 0),
		}, nil
	}

	var users = make([]*models.User, 0, query.GetSize())
	if err := r.db.SelectContext(
		ctx,
		&users,
		getUsers,
		query.GetOrderBy(),
		query.GetOffset(),
		query.GetLimit(),
	); err != nil {
		return nil, errors.Wrap(err, "authRepo.GetUsers.SelectContext")
	}

	return &models.UsersList{
		TotalCount: totalCount,
		TotalPages: utils.GetTotalPages(totalCount, query.GetSize()),
		Page:       query.GetPage(),
		Size:       query.GetSize(),
		HasMore:    utils.GetHasMore(query.GetPage(), totalCount, query.GetSize()),
		Users:      users,
	}, nil
}

func (r *authRepo) FindByEmail(ctx context.Context, user *models.User) (*models.User, error) {
	foundUser := &models.User{}
	if err := r.db.QueryRowxContext(ctx, findUserByEmail, user.Email).StructScan(foundUser); err != nil {
		return nil, errors.Wrap(err, "authRepo.FindByEmail.QueryRowContext")
	}
	return foundUser, nil
}
