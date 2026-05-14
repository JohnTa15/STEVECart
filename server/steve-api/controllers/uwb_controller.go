package controllers

import (
	"fmt"
	"net/http"
	"steve-api/initializers"
	"steve-api/models"
	"time"

	"github.com/gin-gonic/gin"
)

// SaveUWBPosition upserts the latest position for a UWB node in MariaDB.
// Called from mqtt.go whenever a UWB message arrives.
func SaveUWBPosition(nodeID string, x float64, y float64) error {
	if nodeID == "" {
		return fmt.Errorf("uwb node id is empty")
	}

	var uwbData models.UWBData
	// Try to find an existing record for this node
	result := initializers.DB.Where("uwb_node_id = ?", nodeID).First(&uwbData)

	// Whether found or not, update all fields
	uwbData.UWB_NODEID = nodeID
	uwbData.X_Coordinate = x
	uwbData.Y_Coordinate = y
	uwbData.LastSeen_UWB = time.Now()
	uwbData.Description = "Cart position updated via MQTT"

	if result.Error != nil {
		// Record doesn't exist yet — create it
		if err := initializers.DB.Create(&uwbData).Error; err != nil {
			return fmt.Errorf("failed to create UWB record: %w", err)
		}
	} else {
		// Record exists — update it
		if err := initializers.DB.Save(&uwbData).Error; err != nil {
			return fmt.Errorf("failed to update UWB record: %w", err)
		}
	}
	return nil
}

// GetAllUWBPositions is the exported Gin handler for GET /uwb/positions
// Returns the latest known position for every UWB node (cart/shelve)
func GetAllUWBPositions(c *gin.Context) {
	var positions []models.UWBData
	if err := initializers.DB.Find(&positions).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": positions})
}

// shelvePosition saves a shelve's fixed UWB anchor position.
// Shelves don't move, so this is called once during setup.
func shelvePosition(shelveID string, x float64, y float64) error {
	if shelveID == "" {
		return fmt.Errorf("shelve id is empty")
	}

	var uwbData models.UWBData
	result := initializers.DB.Where("uwb_node_id = ?", shelveID).First(&uwbData)

	uwbData.UWB_NODEID = shelveID
	uwbData.X_Coordinate = x
	uwbData.Y_Coordinate = y
	uwbData.LastSeen_UWB = time.Now()
	uwbData.Description = "Shelve fixed position"

	if result.Error != nil {
		return initializers.DB.Create(&uwbData).Error
	}
	return initializers.DB.Save(&uwbData).Error
}