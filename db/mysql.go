package db

import (
	"clean-architecture/config"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func MysqlConnection() (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
		config.ENV.Mysql.MySQLUser,
		config.ENV.Mysql.MySQLPassword,
		config.ENV.Mysql.MySQLHost,
		config.ENV.Mysql.MySQLPort,
		config.ENV.Mysql.MySQLDbname,
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
