package models

import (
	"strings"
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        string `gorm:"primary_key"`
	Name      string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeleteAt  time.Time
}

func (u *User) PrepareCreate() error {
	u.Email = strings.ToLower(strings.TrimSpace(u.Email))
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
	return nil
}

func (user *User) GenerateID() error {
	user.ID = uuid.NewString()
	return nil
}

func (user *User) Validate() error {
	return nil
}
