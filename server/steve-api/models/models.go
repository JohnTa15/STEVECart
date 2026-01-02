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
