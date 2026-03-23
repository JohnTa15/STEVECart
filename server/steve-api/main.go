package main

import (
	"net/http"
	"time"
	"steve-api/controllers"
	"steve-api/models"
	"github.com/gin-gonic/gin"
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
		"cart_id": cart.Cart_ID})
}

func main() {
	r := gin.Default()
	r.POST("/registerCartID", CartRegister)
}

//https://www.emqx.com/en/blog/how-to-use-mqtt-in-golang
//https://github.com/eclipse-paho/paho.mqtt.golang?tab=readme-ov-file
//https://medium.com/widle-studio/integrating-mariaDB-with-golang-microservice-restful-apis-building-efficient-data-storage-4054e1588ce
//https://go.dev/doc/modules/gomod-ref
//https://docs.influxdata.com/influxDB/v2/api-guide/client-libraries/go/
//https://gorm.io/docs/associations.html#tags
//https://gemini.google.com/share/250c2d25448c
