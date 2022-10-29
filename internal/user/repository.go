package user

import (
	"clean-architecture/internal/user/models"
	"context"
)

// mockgen -source="internal/user/repository.go" -destination="internal/user/mock/repository.go" -package=mock
type Repository interface {
	CreateUser(ctx context.Context, user *models.User) error
	UpdateUser(ctx context.Context, userID string, userUpdate map[string]interface{}) error
	GetUser(ctx context.Context, userID string) (*models.User, error)
	GetUsers(ctx context.Context) ([]models.User, error)
}
