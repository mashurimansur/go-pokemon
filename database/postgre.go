package database

import (
	"fmt"
	"go-pokemon/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectionDatabase(config config.Config) *gorm.DB {
	username := config.Get("PG_USERNAME")
	password := config.Get("PG_PASSWORD")
	database := config.Get("PG_DATABASE")
	host := config.Get("PG_HOST")
	port := config.Get("PG_PORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Makassar", host, username, password, database, port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return db
}
