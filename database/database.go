package database

import (
	"log"
	"ramen-go/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConnectWithDB() {
	dsn := "host=localhost user=root password=root dbname=ramendb port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(dsn))
	if err != nil {
		log.Fatal("Erro ao conectar com DB: ", err)
	}

	var model models.Broth
	DB.AutoMigrate(&model)

}
