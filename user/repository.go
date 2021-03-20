package user

import (
	"context"

	"github.com/frasnym/go-docker/models"
)

// Repository represent the user's repository contract
type Repository interface {
	Validate(user *models.User) error
	Create(ctx context.Context, user *models.User) (*models.User, error)
}
