package main

import (
	"fmt"
	"net/http"
	"time"

	"steve-api/controllers"
	"steve-api/initializers"
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
		"message":     "Cart registered successfully",
		"cart_id":     cart.Cart_ID,
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

		var uwbData models.UWBData
		uwb_x := 0.0
		uwb_y := 0.0
		if err := initializers.DB.Where("uwb_node_id = ?", cart.Cart_ID).First(&uwbData).Error; err == nil {
			uwb_x = uwbData.X_Coordinate
			uwb_y = uwbData.Y_Coordinate
		}

		gps := "Unknown"
		if uwb_x > 0 && uwb_y > 0 {
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
			"cart_id":          cart.Cart_ID,
			"fw_version":       cart.FwVersion,
			"is_active":        cart.IsActive,
			"battery_level":    battery,
			"gps":              gps,
			"needs_assistance": cart.NeedsAssistance,
		})
	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": response})
}

func RequestAssistance(c *gin.Context) {
	cartID := c.Query("cartID")
	if cartID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "cartID is required"})
		return
	}
	var cart models.Cart
	if err := initializers.DB.Where("cart_id = ?", cartID).First(&cart).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "cart not found"})
		return
	}
	cart.NeedsAssistance = true
	initializers.DB.Save(&cart)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Assistance requested", "cart_id": cartID})
}

func DismissAssistance(c *gin.Context) {
	cartID := c.Query("cartID")
	if cartID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "cartID is required"})
		return
	}
	var cart models.Cart
	if err := initializers.DB.Where("cart_id = ?", cartID).First(&cart).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "cart not found"})
		return
	}
	cart.NeedsAssistance = false
	initializers.DB.Save(&cart)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Assistance dismissed", "cart_id": cartID})
}

func main() {
	r := gin.Default()
	r.Use(CORSMiddleware())

	r.POST("admin/login", controllers.AdminLogin)

	//routing carts
	r.POST("/registerCartID", CartRegister)
	r.GET("/carts", GetAllCarts)

	//routing users
	r.GET("/users", GetAllUsers)
	r.GET("/userStats", controllers.GetUserStats)
	r.PUT("/setAdmin", controllers.SetAdmin)
	r.POST("/registerUser", controllers.RegisterUser)
	r.POST("/loginUser", controllers.LoginUser)
	r.DELETE("/deleteUserByEmail", controllers.DeleteUserByEmail)

	//routing products
	r.GET("/products", controllers.GetAllProducts)
	r.GET("/GetProducts", controllers.GetAllProducts)
	r.GET("/cartItems", controllers.GetCartItems)
	r.POST("/addProduct", controllers.AddProduct)
	r.PUT("/updateProduct", controllers.UpdateProduct)
	r.DELETE("/deleteProduct", controllers.DeleteProduct)

	//routing sensors
	r.GET("/measureWeight", controllers.MeasureWeightHandler)
	r.GET("/measureLight", controllers.MeasureLight)
	r.GET("/measureDistance", controllers.MeasureDistance)
	r.GET("/measureBattery", controllers.MeasureBattery)
	r.GET("/getLatestNFC", controllers.GetLatestNFC)

	//routing shelves
	r.GET("/shelves", controllers.GetAllShelves)
	r.POST("/addShelvePosition", controllers.AddShelvePosition)
	r.DELETE("/deleteShelvePosition", controllers.DeleteShelvePosition)
	r.PUT("/updateShelvePosition", controllers.UpdateShelvePosition)

	//routing assistance
	r.GET("/assistance", RequestAssistance)
	r.GET("/dismissAssistance", DismissAssistance)

	//routing UWB live tracking
	r.GET("/ws/minimap", controllers.HandleMinimapWS)
	r.GET("/uwb/positions", controllers.GetAllUWBPositions)
	r.GET("/uwb/location", controllers.GetCartLocation)

	//connecting to database, influxdb, and mqtt
	initializers.LoadENV()
	initializers.ConnectDB()
	initializers.ConnectINFLUX()
	client := initializers.ConnectMQTT()
	initializers.Sub(client)
	// Start the goroutine that fans UWB positions out to WebSocket clients
	controllers.StartUWBBroadcaster()
	r.Run(":8089")
}

func CORSMiddleware() gin.HandlerFunc { //middleware which adds headers
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

//https://www.emqx.com/en/blog/how-to-use-mqtt-in-golang
//https://github.com/eclipse-paho/paho.mqtt.golang?tab=readme-ov-file
//https://medium.com/widle-studio/integrating-mariaDB-with-golang-microservice-restful-apis-building-efficient-data-storage-4054e1588ce
//https://go.dev/doc/modules/gomod-ref
//https://docs.influxdata.com/influxDB/v2/api-guide/client-libraries/go/
//https://gorm.io/docs/associations.html#tags
//https://gemini.google.com/share/250c2d25448c
