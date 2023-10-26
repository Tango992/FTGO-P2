package config

import (
	"os"
	"ugc-9/entity"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetupDb() *gorm.DB {
	dsn := os.Getenv("DB")

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&entity.Store{}, &entity.Product{})
	return db
}