package db

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func OpenDatabase() {
	err := godotenv.Load()
	if err != nil{
		log.Fatal("Error Loading .env")
	}
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",os.Getenv("DB_HOST"),os.Getenv("DB_USER"),os.Getenv("DB_PASSWORD"),os.Getenv("DB_DBNAME"),os.Getenv("DB_PORT"),os.Getenv("DB_SSLMODE"),os.Getenv("DB_TIMEZONE"))
	db,err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil{
		panic(err)
	}

	fmt.Println("BERHASIL",os.Getenv("DB_HOST"))

	DB = db
}