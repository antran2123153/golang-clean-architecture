package main

import (
	"clean-architecture/config"
	"clean-architecture/db"
	userModel "clean-architecture/internal/user/models"
	"log"

	"gorm.io/gorm"
)

func main() {
	loadConfig()
	mysqlDB := loadMySQLDB()
	// postgresDB := loadPosgresDB()

	if err := mysqlDB.AutoMigrate(&userModel.User{}); err != nil {
		log.Fatal(err)
	}

	log.Println("Migrate successful")
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
