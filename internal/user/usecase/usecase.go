package usecase

import (
	"clean-architecture/internal/user"
	"clean-architecture/internal/user/models"
	"context"
)

type userUsecase struct {
	repository user.Repository
}

func NewUserUsecase(repository user.Repository) user.UseCase {
	return &userUsecase{
		repository: repository,
	}
}

func (usecase *userUsecase) CreateUser(ctx context.Context, name string, email string) error {
	user := models.User{
		Name:  name,
		Email: email,
	}
	if err := user.GenerateID(); err != nil {
		return err
	}
	if err := user.PrepareCreate(); err != nil {
		return err
	}
	if err := user.Validate(); err != nil {
		return err
	}
	if err := usecase.repository.CreateUser(ctx, &user); err != nil {
		return err
	}
	return nil
}

func (usecase *userUsecase) UpdateUser(ctx context.Context, id string, name string, email string) error {
	if err := usecase.repository.UpdateUser(ctx, id, map[string]interface{}{
		"name":  name,
		"email": email,
	}); err != nil {
		return err
	}
	return nil
}

func (usecase *userUsecase) GetUser(ctx context.Context, id string) (*models.User, error) {
	user, err := usecase.repository.GetUser(ctx, id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (usecase *userUsecase) GetUsers(ctx context.Context) ([]models.User, error) {
	users, err := usecase.repository.GetUsers(ctx)
	if err != nil {
		return nil, err
	}
	return users, nil
}
