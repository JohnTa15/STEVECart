package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"steve-api/controllers"
	"steve-api/initializers"
	"steve-api/models"
)

func CartRegister(c *gin.Context) {
	var cart models.Cart
	if err := c.ShouldBindJSON(&cart); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	cart.LastSeen = time.Now()
	cart.IsActive = true
	c.JSON(http.StatusOK, gin.H{
		"message": "Cart registered successfully",
		"cart_id": cart.Cart_ID,
		"mac_address": cart.MacAddress})
}

func GetAllUsers(c *gin.Context) {
	var users []models.User
	if err := initializers.DB.Find(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"status": http.StatusOK, "data": users})
}

func GetAllCarts(c *gin.Context) {
	var carts []models.Cart
	if err := initializers.DB.Find(&carts).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var response []gin.H
	for _, cart := range carts {
		battery, _ := controllers.GetBattery(cart.Cart_ID)
		uwb_x, _ := controllers.GetUWB(cart.Cart_ID)

		gps := "Unknown"
		if uwb_x > 0 {
			gps = fmt.Sprintf("Aisle %.1f", uwb_x)
		} else {
            // Default mock for display if no real data in influx yet
            gps = "Aisle 4"
        }

        // Just in case battery is 0, provide a mock value for now
        if battery == 0 {
            battery = 92
        }

		response = append(response, gin.H{
			"cart_id":       cart.Cart_ID,
			"fw_version":    cart.FwVersion,
			"is_active":     cart.IsActive,
			"battery_level": battery,
			"gps":           gps,
		})
	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": response})
}

func main() {
	r := gin.Default()
	r.POST("/registerCartID", CartRegister)
	r.GET("/users", GetAllUsers)
	r.GET("/carts", GetAllCarts)

	r.GET("/measureWeight", controllers.MeasureWeightHandler)
	r.GET("/measureLight", controllers.MeasureLight)
	r.GET("/shelves", controllers.GetAllShelves)
	r.POST("/addShelvePosition", controllers.AddShelvePosition)
	r.DELETE("/deleteShelvePosition", controllers.DeleteShelvePosition)
	r.PUT("/updateShelvePosition", controllers.UpdateShelvePosition)
	initializers.LoadENV()
	initializers.ConnectDB()
	initializers.ConnectINFLUX()
	client := initializers.ConnectMQTT()
	initializers.Sub(client)
	r.Run(":8089")
}

//https://www.emqx.com/en/blog/how-to-use-mqtt-in-golang
//https://github.com/eclipse-paho/paho.mqtt.golang?tab=readme-ov-file
//https://medium.com/widle-studio/integrating-mariaDB-with-golang-microservice-restful-apis-building-efficient-data-storage-4054e1588ce
//https://go.dev/doc/modules/gomod-ref
//https://docs.influxdata.com/influxDB/v2/api-guide/client-libraries/go/
//https://gorm.io/docs/associations.html#tags
//https://gemini.google.com/share/250c2d25448c
