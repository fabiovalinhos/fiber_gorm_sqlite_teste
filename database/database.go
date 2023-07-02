package database

import (
	"log"
	"os"

	"github.com/fabiovalinhos/fiber-api/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DbInstance struct {
	Db *gorm.DB
}

var Database DbInstance

func ConnectDb() {
	db, err := gorm.Open(sqlite.Open("api.db"))
	if err != nil {
		log.Fatal("Failed to connect to database! ", err.Error())
		os.Exit(2)
	}

	log.Println("Connected to the database successfully")

	db.Logger = logger.Default.LogMode(logger.Info)
	log.Println("Running Migrations")
	db.AutoMigrate(&models.User{}, &models.Product{}, &models.Order{})

	Database = DbInstance{Db: db}
}
