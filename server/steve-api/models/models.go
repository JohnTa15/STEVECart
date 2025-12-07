package models

import (
	"time"
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
