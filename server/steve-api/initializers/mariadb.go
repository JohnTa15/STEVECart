package initializers

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	ID                 uint `gorm:"primaryKey"`
	ProductName        string
	ProductCategory    string
	ProductAddedDate   time.Time
	ProductDescription string
	Weight             float32
	Pcs                int
	Price              float32
}

type Cart struct {
	ID        uint `gorm:"primaryKey"`
	IsActive  bool
	FwVersion string
}

type User struct {
	ID           uint   `gorm:"primaryKey"`
	Email        string `gorm:"unique"`
	PasswordHash string
	UserCreation time.Time
}

type CartOperator struct {
	ID        uint `gorm:"primaryKey"`
	UserID    uint
	CartID    uint
	EventTime time.Time
}

type CartOperatorItem struct {
	ID         uint `gorm:"primaryKey"`
	UserCartID uint
	ProductID  uint
	Quantity   int
}

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
