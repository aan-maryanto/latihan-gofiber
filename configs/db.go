package configs

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type DBConfig struct {
	HOST     string
	PORT     string
	USER     string
	PASSWORD string
	DATABASE string
}

var Config = DBConfig{
	HOST:     "localhost",
	PORT:     "5432",
	USER:     "postgres",
	PASSWORD: "postgres",
	DATABASE: "postgres",
}

func Connect() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return db
}
