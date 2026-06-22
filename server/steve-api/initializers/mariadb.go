package initializers

import (
	"fmt"
	"log"
	"os"
	"steve-api/models"
	"strings"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() *gorm.DB {
	dsn := os.Getenv("DB_URL")
	if !strings.Contains(dsn, "parseTime=") {
		if strings.Contains(dsn, "?") {
			dsn += "&parseTime=true"
		} else {
			dsn += "?parseTime=true"
		}
	}
	fmt.Println("Connecting with DSN:", dsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to MariaDB: ", err)
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
