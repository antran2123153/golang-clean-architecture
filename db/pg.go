package db

import (
	"clean-architecture/config"
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func PgConnection() (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%v port=%v user=%v password=%v dbname=%v TimeZone=%v",
		config.ENV.Postgres.PostgresqlHost,
		config.ENV.Postgres.PostgresqlPort,
		config.ENV.Postgres.PostgresqlUser,
		config.ENV.Postgres.PostgresqlPassword,
		config.ENV.Postgres.PostgresqlDbname,
		config.ENV.Postgres.PostgresqlTimezone,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	return db, nil
}
