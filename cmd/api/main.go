package main

import (
	"clean-architecture/config"
	"clean-architecture/db"
	userHandler "clean-architecture/internal/user/delivery/http/handler"
	userRepository "clean-architecture/internal/user/repository"
	userUsecase "clean-architecture/internal/user/usecase"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func main() {
	loadConfig()
	mysqlDB := loadMySQLDB()
	// postgresDB := loadPosgresDB()

	userRepo := userRepository.NewUserRepository(mysqlDB)
	userUsecase := userUsecase.NewUserUsecase(userRepo)
	userHandler := userHandler.NewUserHandler(userUsecase)

	r := gin.Default()

	r.POST("users", userHandler.CreateUser())
	r.GET("users/:user_id", userHandler.GetUser())
	r.GET("users", userHandler.GetUsers())
	r.PUT("users/user_id", userHandler.UpdateUser())

	if err := r.Run(); err != nil {
		log.Fatalf("Server run error: %v", err)
	}
}

func loadConfig() {
	cfgFile, err := config.LoadConfig("dev.yml")
	if err != nil {
		log.Fatalf("LoadConfig: %v", err)
	}

	if err := config.ParseConfig(cfgFile); err != nil {
		log.Fatalf("ParseConfig: %v", err)
	}
}

func loadPosgresDB() *gorm.DB {
	db, err := db.PgConnection()
	if err != nil {
		log.Fatalf("Postgresql init error: %s", err)
	} else {
		log.Println("Postgres connected")
	}
	return db
}

func loadMySQLDB() *gorm.DB {
	db, err := db.MysqlConnection()
	if err != nil {
		log.Fatalf("MySQL init error: %s", err)
	} else {
		log.Println("MySQL connected")
	}
	return db
}
