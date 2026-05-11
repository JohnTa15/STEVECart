package models

import (
	"time"
)

type ShelvePosition struct{
	ID uint `gorm:"primaryKey"`
	Shelve_id string `gorm:"unique" json:"shelve_id"`
	X_Coordinate float64 `json:"x_coordinate"`
	Y_Coordinate float64 `json:"y_coordinate"`
	Description string `json:"description"`
}

type Product struct {
	ID                 uint `gorm:"primaryKey"`
	ProductName        string
	ProductCategory    string
	NFCTag             string
	ProductAddedDate   time.Time
	ProductDescription string
	Weight             float32
	Pcs                int
	Price              float32
	ShelveID string `gorm:"index" json:"shelve_id"` //shelve where the product is located
}

type Cart struct {
	ID               uint      `gorm:"primaryKey" json:"-"`
	Cart_ID          string    `gorm:"unique" json:"cart_id"`
	MacAddress       string    `gorm:"unique" json:"mac_address"`
	IsActive         bool      `json:"is_active"`
	LastSeen         time.Time `json:"last_seen"`
	FwVersion        string    `json:"fw_version"`
	TotalPrice       float32   `json:"total_price"`
	TotalWeight      float32   `json:"total_weight"`
	NeedsAssistance  bool      `json:"needs_assistance"`
}

type User struct {
	ID           uint      `gorm:"primaryKey" json:"-"`
	Username     string    `gorm:"unique" json:"username"`
	Email        string    `gorm:"unique" json:"email"`
	PasswordHash string    `json:"password_hash"`
	UserCreation time.Time `json:"user_creation"`
	Role		 string		`json:"role"`
}

//the one who uses a cart with id cart_id
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

// for mqtt.go
type WeightData struct {
	ID           uint      `gorm:"primaryKey"`
	Value        float64   `json:"weight"`
	StableWeight bool      `json:"isStable"`
	Timestamp    time.Time `json:"timestamp"`
}

type NFCData struct {
	TagID     string `json:"tag_id"` // foreign key for product ??
	ScannerID string `json:"scanner_id"`
}

type UltraSonicData struct {
	ID                 uint      `gorm:"primaryKey"`
	Distance           float64   `json:"distance"`
	Timestamp_Distance time.Time `json:"timestamp"`
}

type BatteryData struct {
	ID       uint    `gorm:"primaryKey"`
	BatLevel float64 `json:"battery_level"`
	Charging bool    `json:"isCharging"`
}

type LightSensorData struct {
	ID            uint      `gorm:"primaryKey"`
	LuxLevel      float64   `json:"lux_level"`
	Timestamp_Lux time.Time `json:"timestamp"`
}

type UWBData struct { //for UWB location data
	ID uint `gorm:"primaryKey"`
	UWB_NODEID    string    `json:"uwb_node_id"`
	X_Coordinate  float64   `json:"x_coordinate"`
	Y_Coordinate  float64   `json:"y_coordinate"`
	LastSeen_UWB time.Time `json:"last_seen_uwb"`
	Description string `json:"description"`
}

