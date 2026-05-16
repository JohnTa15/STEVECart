package controllers

import (
	"github.com/gin-gonic/gin"
)

func MeasureBattery(c *gin.Context){
	cartID := c.Query("cartID")
	_ = c.Query("tag")

	_, err := CartChecking(cartID)
	if err != nil {
		c.JSON(500, gin.H{"status": "error", "message": "Cart not found"})
		return
	}

	var minBatteryLevel float64 = 20.0
	battery, err := GetBattery(cartID)
	if err != nil {
		c.JSON(500, gin.H{"status": "error", "message": "Failed to read battery"})
		return
	}

	// checking battery level
	if battery < minBatteryLevel {
		c.JSON(200, gin.H{
			"status": "warning",
			"message": "Battery is low needs charging!",
			"battery" : battery,
		})
	} else {
		c.JSON(200, gin.H{
			"status": "success",
			"message": "Battery is fine!",
			"battery" : battery,
		})
	}

}
