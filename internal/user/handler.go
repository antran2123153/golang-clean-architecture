package user

import (
	"github.com/gin-gonic/gin"
)

type Handler interface {
	CreateUser() gin.HandlerFunc
	UpdateUser() gin.HandlerFunc
	GetUser() gin.HandlerFunc
	GetUsers() gin.HandlerFunc
}
