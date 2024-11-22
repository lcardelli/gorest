package database

import (
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"log"
)

var (
	DB  *gorm.DB
	err error
)

func Connect() {

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Erro ao carregar o arquivo .env: %v", err)
	}

	dbUser := os.Getenv("user")
	dbPass := os.Getenv("password")

	dsn := "host=localhost " +
       "user=" + dbUser + " " +
       "password=" + dbPass + " " +
       "dbname=root " +
       "port=5432 " +
       "sslmode=disable"

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panic("Failed to connect to database")
	}

}
