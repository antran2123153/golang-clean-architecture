package handler

import (
	"clean-architecture/internal/user"
	"clean-architecture/internal/user/delivery/http/presenter"
	"clean-architecture/pkg/http_errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	usecase user.UseCase
}

func NewUserHandler(usecase user.UseCase) *userHandler {
	return &userHandler{usecase: usecase}
}

func (handler *userHandler) CreateUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var request presenter.CreateUser
		if err := ctx.Bind(request); err != nil {
			ctx.JSON(http_errors.ErrorResponse(err))
			return
		}
		if err := handler.usecase.CreateUser(ctx.Request.Context(), request.Name, request.Email); err != nil {
			ctx.JSON(http_errors.ErrorResponse(err))
			return
		}
		ctx.JSON(http.StatusNoContent, nil)
	}
}

func (handler *userHandler) UpdateUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var request presenter.UpdateUser
		if err := ctx.Bind(request); err != nil {
			ctx.JSON(http_errors.ErrorResponse(err))
			return
		}
		userID := ctx.Param("user_id")
		err := handler.usecase.UpdateUser(ctx.Request.Context(), userID, request.Name, request.Email)
		if err != nil {
			ctx.JSON(http_errors.ErrorResponse(err))
			return
		}
		ctx.JSON(http.StatusNoContent, nil)
	}
}

func (handler *userHandler) GetUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userID := ctx.Param("user_id")
		user, err := handler.usecase.GetUser(ctx.Request.Context(), userID)
		if err != nil {
			ctx.JSON(http_errors.ErrorResponse(err))
			return
		}
		ctx.JSON(http.StatusOK, user)
	}
}

func (handler *userHandler) GetUsers() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		users, err := handler.usecase.GetUsers(ctx.Request.Context())
		if err != nil {
			ctx.JSON(http_errors.ErrorResponse(err))
			return
		}
		ctx.JSON(http.StatusOK, users)
	}
}
