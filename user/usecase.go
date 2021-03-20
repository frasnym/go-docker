package user

import (
	"context"

	"github.com/frasnym/go-docker/models"
)

// Usecase represent the user's usecases
type Usecase interface {
	Validate(*models.User) error
	Create(context.Context, *models.User) (*models.User, error)
}
