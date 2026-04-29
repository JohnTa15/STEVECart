package controllers

import(
	"github.com/gin-gonic/gin"
)

func MeasureLight(c *gin.Context){
	cartID := c.Query("cartID")
	_ = c.Query("tag")

	_, err := CartChecking(cartID)
	if err != nil {
		c.JSON(500, gin.H{"status": "error", "message": "Cart not found"})
		return
	}

	var minLux float64 = 80
	lux, err := GetLux(cartID)
	if err != nil {
		c.JSON(500, gin.H{"status": "error", "message": "Failed to read lux"})
		return
	}

	// Darkness is a normal physical state, not a server crash! Return 200 OK.
	if lux < minLux { 
		c.JSON(200, gin.H{
			"status": "dark", 
			"message": "Light is not enough, opening flashlight",
			"lux": lux,
		})
	} else {
		c.JSON(200, gin.H{
			"status": "bright", 
			"message": "Light is okay",
			"lux": lux,
		})
	}
}