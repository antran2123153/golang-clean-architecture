package service

import (
	"clean-architecture/internal/user/delivery/http/handler"
	"clean-architecture/internal/user/repository"
	"clean-architecture/internal/user/usecase"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func StartUserService(r *gin.Engine, db *gorm.DB) {
	userRepo := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepo)
	userHandler := handler.NewUserHandler(userUsecase)

	userRouter := r.Group("users")
	{
		userRouter.POST("", userHandler.CreateUser())
		userRouter.GET("", userHandler.GetUsers())
		userRouter.GET("/:user_id", userHandler.GetUser())
		userRouter.PUT("/:user_id", userHandler.UpdateUser())
	}
}