package services

import (
	"context"
	"errors"
	"savannah/internal/models"
	"savannah/internal/repository/db"
)

type UserService interface {
	UserCreate(ctx context.Context, user models.User) error
	UserUpdate(ctx context.Context, user models.User) error
	UserDelete(ctx context.Context, user models.User) error
	UserGet(ctx context.Context, userName string) (*models.User, error)
	UserList(ctx context.Context, userName string) ([]*models.User, error)
	Verify(ctx context.Context, token string, username string) error
}

func (s *SQLStore) UserCreate(ctx context.Context, user models.User) error {
	if err := s.store.CreateUser(ctx, db.CreateUserParams{
		UserName: user.Username,
		Token:    user.AuthToken,
	}); err != nil {
		return err
	}
	return nil
}

func (s *SQLStore) UserUpdate(ctx context.Context, user models.User) error {
	if err := s.store.UpdateUser(ctx, db.UpdateUserParams{
		ID:       int32(user.ID),
		UserName: user.Username,
	}); err != nil {
		return err
	}
	return nil
}

func (s *SQLStore) UserDelete(ctx context.Context, user models.User) error {
	if err := s.store.DeleteUser(ctx, int32(user.ID)); err != nil {
		return err
	}
	return nil
}

func (s *SQLStore) UserGet(ctx context.Context, userName string) (*models.User, error) {
	user, err := s.store.GetUser(ctx, userName)
	if err != nil {
		return nil, err
	}
	return &models.User{
		ID:        int64(user.ID),
		CreatedAt: user.CreatedAt.Time,
		UpdatedAt: user.UpdatedAt.Time,
		DeletedAt: user.DeletedAt.Time,
		Username:  user.UserName,
	}, nil
}

func (s *SQLStore) UserList(ctx context.Context, userName string) ([]*models.User, error) {
	users, err := s.store.ListUsers(ctx)
	if err != nil {
		return nil, err
	}
	usersList := make([]*models.User, len(users))
	for i, v := range users {
		usersList[i] = &models.User{
			ID:        int64(v.ID),
			CreatedAt: v.CreatedAt.Time,
			UpdatedAt: v.UpdatedAt.Time,
			DeletedAt: v.DeletedAt.Time,
			Username:  v.UserName,
		}
	}
	return usersList, nil
}

func (s *SQLStore) Verify(ctx context.Context, token string, username string) error {
	user, err := s.store.GetUser(ctx, username)
	if err != nil {
		return err
	}
	if user.Token != token {
		return errors.New("Invalid token")
	}
	return nil
}
