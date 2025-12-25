package initializers

import (
	"fmt"
	"log"
	"os"
	"time"
  "steve-api/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() *gorm.DB {
	dsn := os.Getenv("DB_URL")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to MariaDB: ", err)
	}
	fmt.Println("Successfully connected to MariaDB")

	// store connection in package variable for other packages to use
	DB = db

	if err := db.AutoMigrate(&Product{}, &Cart{}, &User{}, &CartOperator{}, &CartOperatorItem{}); err != nil {
		log.Fatal("AutoMigrate failed:", err)
	}
	return db
}
