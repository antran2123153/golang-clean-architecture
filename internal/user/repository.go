package user

import (
	"clean-architecture/internal/user/models"
	"context"
)

type Repository interface {
	CreateUser(ctx context.Context, user *models.User) error
	UpdateUser(ctx context.Context, userID string, userUpdate map[string]interface{}) error
	GetUser(ctx context.Context, userID string) (*models.User, error)
	GetUsers(ctx context.Context) ([]models.User, error)
}
