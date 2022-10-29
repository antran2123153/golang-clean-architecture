package presenter

import (
	"clean-architecture/internal/user/models"
	"github.com/samber/lo"
)

type User struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (user *User) MakeUserPresenter(u *models.User) error {
	user.ID = u.ID
	user.Name = u.Name
	user.Email = u.Email
	return nil
}

type ListUser struct {
	Users []User `json:"users"`
	Count int    `json:"count"`
}

func (users *ListUser) MakeListUserPresenter(us []models.User) error {
	users.Users = lo.Map(us, func(u models.User, index int) User {
		return User{
			ID: u.ID,
			Name: u.Name,
			Email: u.Email,
		}
	})
	users.Count = len(us)
	return nil
}

type CreateUser struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UpdateUser struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}