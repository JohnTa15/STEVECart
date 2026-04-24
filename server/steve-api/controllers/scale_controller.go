package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

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

	// 3. Compare the weights and send JSON response
	if weight == float64(product.Weight) {
		fmt.Println("Weight is correct! Adding it to the display")
		c.JSON(200, gin.H{
			"cart_id":         cart.Cart_ID,
			"product_name":    product.ProductName,
			"expected_weight": product.Weight,
			"actual_weight":   weight,
			"status":          "correct",
		})
	} else {
		fmt.Println("Weight is incorrect, check again if this is the right product..")
		c.JSON(200, gin.H{
			"cart_id":         cart.Cart_ID,
			"product_name":    product.ProductName,
			"expected_weight": product.Weight,
			"actual_weight":   weight,
			"status":          "incorrect",
		})
	}
}