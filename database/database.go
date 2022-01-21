package database

import (
	"fmt"
	"log"
	"strconv"

	"github.com/pace-noge/fiber-notes-api/config"
	"github.com/pace-noge/fiber-notes-api/internals/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	var err error
	p := config.Config("DB_PORT")
	port, err := strconv.ParseUint(p, 10, 32)

	if err != nil {
		log.Println("Cannot get port")
	}

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", config.Config("DB_HOST"), port, config.Config("DB_USER"), config.Config("DB_PASSWORD"), config.Config("DB_NAME"))

	DB, err = gorm.Open(postgres.Open(dsn))

	if err != nil {
		panic("failed to connect to database")
	}

	log.Println("Connection opened to database")

	DB.AutoMigrate(&model.Note{})
	log.Println("Database Migrated")
}
