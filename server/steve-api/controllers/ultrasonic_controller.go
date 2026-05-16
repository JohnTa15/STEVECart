package controllers

import (
	"github.com/gin-gonic/gin"
)

func MeasureDistance(c *gin.Context){
	cartID := c.Query("cartID")
	_ = c.Query("tag")

	_, err := CartChecking(cartID)
	if err != nil {
		c.JSON(500, gin.H{"status": "error", "message": "Cart not found"})
		return
	}

	distance, err := GetDistance(cartID)
	if err != nil {
		c.JSON(500, gin.H{"status": "error", "message": "Failed to read distance"})
		return
	}

	var maxDistance float64 = 20.0
	if distance > maxDistance {
		c.JSON(200, gin.H{
			"status": "success",
			"message": "All clear!",
			"distance": distance,
		})
	} else {
		c.JSON(200, gin.H{
			"status": "warning",
			"message": "The cart is getting close to an object!",
			"distance": distance,
		})
	}
}