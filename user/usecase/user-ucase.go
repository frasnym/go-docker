package usecase

import (
	"context"
	"errors"

	"github.com/frasnym/go-docker/models"
	"github.com/frasnym/go-docker/user"
)

type userUsecase struct {
	userRepo user.Repository
}

// NewUserUsecase will create new an userUsecase object representation of user.Usecase interface
func NewUserUsecase(userRepo user.Repository) user.Usecase {
	return &userUsecase{}
}

func (*userUsecase) Validate(user *models.User) error {
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

func (*userUsecase) Create(ctx context.Context, user *models.User) (*models.User, error) {
	return user, nil
}
