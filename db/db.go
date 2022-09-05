package db

import (
	"log"

	"github.com/symon-nascimento/api/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() *gorm.DB {
	dbURL := "postgres://engine:engine@localhost:5442/engine"
	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}
	
	db.AutoMigrate(model.Student{})

	return db
}