package user

import (
	"clean-architecture/internal/user/models"
	"context"
)

type UseCase interface {
	CreateUser(ctx context.Context, name string, email string) error
	UpdateUser(ctx context.Context, id string, name string, email string) error
	GetUser(ctx context.Context, id string) (*models.User, error)
	GetUsers(ctx context.Context) ([]models.User, error)
}
