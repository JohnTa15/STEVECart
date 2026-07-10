package initializers

import (
	"fmt"
	"log"
	"os"
	"steve-api/models"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() *gorm.DB {
	dsn := os.Getenv("DB_URL")

	var db *gorm.DB
	var err error
	for attempt := 1; attempt <= 10; attempt++ {
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err == nil {
			break
		}
		log.Printf("MariaDB not ready (attempt %d/10): %v", attempt, err)
		time.Sleep(5 * time.Second)
	}
	if err != nil {
		log.Fatal("Failed to connect to MariaDB after 10 attempts: ", err)
	}
	fmt.Println("Successfully connected to MariaDB")

	// store connection in package variable for other packages to use
	DB = db

	if err := db.AutoMigrate(
		&models.Product{},
		&models.Cart{},
		&models.User{},
		&models.UserCartSession{},
		&models.UserCartItem{},
		&models.RemainingProduct{},
		&models.Shelves{},
		&models.UWBData{}); err != nil {
		log.Fatal("AutoMigrate failed:", err)
	}
	return db
}
