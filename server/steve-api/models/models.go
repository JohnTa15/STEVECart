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
	ID         uint      `gorm:"primaryKey" json:"-"`
	Cart_ID    string    `gorm:"unique" json:"cart_id"`
	MacAddress string    `gorm:"unique" json:"mac_address"`
	IsActive   bool      `json:"is_active"`
	LastSeen   time.Time `json:"last_seen"`
	FwVersion  string    `json:"fw_version"`
}

type User struct {
	ID           uint   `gorm:"primaryKey" json:"-"`
	Username     string `gorm:"unique" json:"username"`
	Email        string `gorm:"unique" json:"email"`
	PasswordHash string `json:"password_hash"`
	UserCreation time.Time `json:"user_creation"`
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

// for mqtt.go
type WeightData struct {
	Value        float64   `json:"weight"`
	StableWeight bool      `json:"isStable"`
	Timestamp    time.Time `json:"timestamp"`
}

type NFCData struct {
	TagID     string `json:"tag_id"`
	ScannerID string `json:"scanner_id"`
}

type UltraSonicData struct {
	Distance           float64   `json:"distance"`
	Timestamp_Distance time.Time `json:"timestamp"`
}

type BatteryData struct {
	BatLevel float64 `json:"battery_level"`
	Charging bool    `json:"isCharging"`
}

type LightSensorData struct {
	LuxLevel      float64   `json:"lux_level"`
	Timestamp_Lux time.Time `json:"timestamp"`
}

type UWBData struct { //for UWB location data
	UWB_NODEID    string    `json:"uwb_node_id"`
	X_Coordinate  float64   `json:"x_coordinate"`
	Y_Coordinate  float64   `json:"y_coordinate"`
	Z_Coordinate  float64   `json:"z_coordinate"`
	Timestamp_UWB time.Time `json:"timestamp"`
}
