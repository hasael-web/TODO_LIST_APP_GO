package config

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// var Db *gorm.DB

type Databse struct {
	DB *gorm.DB
}

func InitDB() *Databse {
	return &Databse{}
}

func (pg *Databse) ConnectDb() (*gorm.DB, error) {

	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	port := os.Getenv("DB_PORT")
	name := os.Getenv("DATABASE")
	sslMode := os.Getenv("DB_SSLMODE")
	tz := os.Getenv("DB_TZ")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s", host, user, pass, name, port, sslMode, tz)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		return nil, err
	}

	pg.DB = db
	return db, err

}
