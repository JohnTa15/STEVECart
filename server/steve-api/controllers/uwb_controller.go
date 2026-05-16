package controllers

import (
	"fmt"
	"net/http"
	"steve-api/initializers"
	"steve-api/models"
	"time"

	"github.com/gin-gonic/gin"
)

func SaveUWBPosition(nodeID string, x float64, y float64) error {
	if nodeID == "" {
		return fmt.Errorf("uwb node id is empty")
	}

	var uwbData models.UWBData
	result := initializers.DB.Where("uwb_node_id = ?", nodeID).First(&uwbData)

	uwbData.UWB_NODEID = nodeID
	uwbData.X_Coordinate = x
	uwbData.Y_Coordinate = y
	uwbData.LastSeen_UWB = time.Now()
	uwbData.Description = "Cart position updated via MQTT"

	if result.Error != nil {
		if err := initializers.DB.Create(&uwbData).Error; err != nil {
			return fmt.Errorf("failed to create UWB record: %w", err)
		}
	} else {
		if err := initializers.DB.Save(&uwbData).Error; err != nil {
			return fmt.Errorf("failed to update UWB record: %w", err)
		}
	}
	return nil
}

func GetAllUWBPositions(c *gin.Context) {
	var positions []models.UWBData
	if err := initializers.DB.Find(&positions).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": positions})
}

func GetCartLocation(c *gin.Context) {
	cartID := c.Query("cartID")

	_, err := CartChecking(cartID)
	if err != nil {
		c.JSON(500, gin.H{"status": "error", "message": "Cart not found"})
		return
	}

	rangeData, err := GetUWB(cartID)
	if err != nil {
		c.JSON(500, gin.H{"status": "error", "message": "Failed to get UWB data"})
		return
	}

	c.JSON(200, gin.H{
		"status": "success",
		"cart_id": cartID,
		"range_data": rangeData,
	})
}
