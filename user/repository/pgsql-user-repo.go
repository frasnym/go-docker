package repository

import (
	"context"
	"errors"

	"github.com/frasnym/go-docker/models"
	"github.com/frasnym/go-docker/user"
)

type pgsqlUserRepository struct{}

// NewPgSqlUserRepository will create new an userUsecase object representation of user.Usecase interface
func NewPgSqlUserRepository() user.Repository {
	return &pgsqlUserRepository{}
}

func (*pgsqlUserRepository) Validate(user *models.User) error {
	if user == nil {
		return errors.New("User is empty")
	}
	if user.Name == "" {
		return errors.New("Name cannot be empty")
	}
	if len(user.Password) < 5 {
		return errors.New("Please provide a stronger password")
	}

	return nil
}

func (*pgsqlUserRepository) Create(ctx context.Context, user *models.User) (*models.User, error) {
	return user, nil
}
