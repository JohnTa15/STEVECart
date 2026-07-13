package controllers

import (
	"fmt"
	"math"
	"steve-api/initializers"

	"github.com/gin-gonic/gin"
)

// RawWeight returns the current scale reading and the last verified cart
// weight, so the display can detect items placed in the cart without an NFC scan.
func RawWeight(c *gin.Context) {
	cartID := c.Query("cartID")
	cart, err := CartChecking(cartID)
	if err != nil {
		c.JSON(404, gin.H{"status": "error", "message": "Cart not found"})
		return
	}
	weight, err := GetWeight(cart.Cart_ID)
	if err != nil {
		c.JSON(500, gin.H{"status": "error", "message": "Failed to read scale"})
		return
	}
	c.JSON(200, gin.H{
		"status":          "ok",
		"scale_weight":    weight,
		"verified_weight": cart.NetWeight,
	})
}

func MeasureWeightHandler(c *gin.Context) {
	// 1. Get arguments from the Vue request URL parameters
	cartID := c.Query("cartID")
	tag := c.Query("tag") // Assumes Vue sends the tag too: ?cartID=123&tag=NFC_CODE

	// 2. Fetch the data
	cart, err := CartChecking(cartID)
	if err != nil {
		c.JSON(500, gin.H{"status": "error", "message": "Cart not found"})
		return
	}

	product, err := ProductChecking(tag) // Now passing the actual tag from the request
	if err != nil {
		c.JSON(500, gin.H{"status": "error", "message": "Product not found"})
		return
	}

	weight, err := GetWeight(cart.Cart_ID)
	if err != nil {
		c.JSON(500, gin.H{"status": "error", "message": "Failed to read scale"})
		return
	}

	//the scale is not taring(reseting to 0kg each time an item is removed a product)
	weightDif := weight - cart.NetWeight
	expectedWeight := product.Weight
	tolerance := 0.05 // tolerance of scale errors.. 

	if math.Abs(math.Abs(weightDif) - expectedWeight) < tolerance {
		fmt.Println("Weight is correct! Adding/Removing it to the display")
		cart.NetWeight = weight // Update to the new absolute weight from the scale
		initializers.DB.Save(&cart)
		c.JSON(200, gin.H{
			"cart_id":         cart.Cart_ID,
			"product_name":    product.ProductName,
			"expected_weight": product.Weight,
			"actual_weight":   weight,
			"net_weight":      cart.NetWeight,
			"price":           product.Price,
			"status":          "correct",
		})
	} else {
		fmt.Println("Weight is incorrect, check again if this is the right product..")
		c.JSON(200, gin.H{
			"cart_id":         cart.Cart_ID,
			"product_name":    product.ProductName,
			"expected_weight": product.Weight,
			"net_weight":      cart.NetWeight,
			"actual_weight":   weight,
			"price":           product.Price,
			"status":          "incorrect",
		})
	}
}
