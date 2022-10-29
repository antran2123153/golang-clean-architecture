package repository

import (
	"clean-architecture/internal/user"
	"clean-architecture/internal/user/models"
	"context"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) user.Repository {
	return &userRepository{db: db}
}

func (repo *userRepository) CreateUser(ctx context.Context, u *models.User) error {
	if err := repo.db.Create(u).Error; err != nil {
		return err
	}
	return nil
}

func (repo *userRepository) UpdateUser(ctx context.Context, userID string, userUpdate map[string]interface{}) error {
	if err := repo.db.Model(models.User{}).Where("id = ?", userID).Updates(userUpdate).Error; err != nil {
		return err
	}
	return nil
}

func (repo *userRepository) GetUser(ctx context.Context, userID string) (*models.User, error) {
	var u models.User
	if err := repo.db.Model(models.User{}).Where("id = ?", userID).First(&u).Error; err != nil {
		return nil, err
	}
	return &u, nil
}

func (repo *userRepository) GetUsers(ctx context.Context) ([]models.User, error) {
	var users []models.User
	if err := repo.db.Model(models.User{}).Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}
