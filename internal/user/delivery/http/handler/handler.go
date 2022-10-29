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

// CreateUser godoc
// @Summary      Create new user
// @Description  Create new user with user name and email
// @Tags         users
// @Accept       json
// @Produce      json
// @Router       /users [post]
func (handler *userHandler) CreateUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var request presenter.CreateUser
		if err := ctx.Bind(&request); err != nil {
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

// UpdateUser godoc
// @Summary      Update  user
// @Description  Update user by user_id
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        user_id  path  string  true  "User ID"
// @Router       /users/:user_id [put]
func (handler *userHandler) UpdateUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var request presenter.UpdateUser
		if err := ctx.Bind(&request); err != nil {
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

// GetUser godoc
// @Summary      Get user info
// @Description  Get user info by user_id
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        user_id  path      string  true  "User ID"
// @Success      200      {object}  presenter.User
// @Router       /users/:user_id [get]
func (handler *userHandler) GetUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userID := ctx.Param("user_id")
		user, err := handler.usecase.GetUser(ctx.Request.Context(), userID)
		if err != nil {
			ctx.JSON(http_errors.ErrorResponse(err))
			return
		}

		var response presenter.User
		err = response.MakeUserPresenter(user)
		if err != nil {
			ctx.JSON(http_errors.ErrorResponse(err))
			return
		}

		ctx.JSON(http.StatusOK, response)
	}
}

// GetUsers godoc
// @Summary      Get list user
// @Description  Get list user
// @Tags         users
// @Accept       json
// @Produce      json
// @Success      200  {object}  presenter.ListUser
// @Router       /users [get]
func (handler *userHandler) GetUsers() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		users, err := handler.usecase.GetUsers(ctx.Request.Context())
		if err != nil {
			ctx.JSON(http_errors.ErrorResponse(err))
			return
		}

		var response presenter.ListUser
		err = response.MakeListUserPresenter(users)
		if err != nil {
			ctx.JSON(http_errors.ErrorResponse(err))
			return
		}

		ctx.JSON(http.StatusOK, response)
	}
}
