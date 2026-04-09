package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
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

func main() {
	r := gin.Default()
	r.POST("/registerCartID", CartRegister)
	r.GET("/users", GetAllUsers)
	initializers.ConnectDB()
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
