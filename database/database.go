package database

import (
	"log"
	"os"
	"ramen-go/models"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConnectWithDB() {
	godotenv.Load()
	user := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	url := os.Getenv("DB_URL")
	dsn := "host=" + url + " user=" + user + " password=" + password + " dbname=ramendb port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(dsn))
	if err != nil {
		log.Fatal("Erro ao conectar com DB: ", err)
	}
	var modelOrder models.Order
	var modelBroth models.Broth
	var modelProtein models.Protein
	DB.AutoMigrate(&modelBroth, &modelProtein, &modelOrder)

}
