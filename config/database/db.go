package database

import (
	"fmt"
	"log"
	"mygram-api/domain"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func StartDB() *gorm.DB {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file: ", err)
	}

	env := os.Getenv("ENV")
	host := os.Getenv("PGHOST")
	user := os.Getenv("PGUSER")
	password := os.Getenv("PGPASSWORD")
	dbname := os.Getenv("PGDBNAME")
	port := os.Getenv("PGPORT")
	dsn := ""

	if env == "production" {
		dsn = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=require", host, user, password, dbname, port)
	} else {
		dsn = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbname, port)
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Error connecting to database: ", err)
	}

	if env == "production" {
		if err := db.AutoMigrate(&domain.User{}, &domain.Photo{}, &domain.Comment{}, &domain.SocialMedia{}); err != nil {
			log.Fatal("Error migrating database: ", err.Error())
		}
	} else {
		if err := db.Debug().AutoMigrate(&domain.User{}, &domain.Photo{}, &domain.Comment{}, &domain.SocialMedia{}); err != nil {
			log.Fatal("Error migrating database: ", err.Error())
		}
	}

	return db
}

func GetDB() (db *gorm.DB) {
	return db
}
