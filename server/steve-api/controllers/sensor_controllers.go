package controllers

import (
	"fmt"
	"steve-api/initializers"
	"steve-api/models"
)

func CartChecking(cartID string) {
	var cart models.Cart
	resC := initializers.DB.Where("cart_id = ?", cartID).First(&cart)
	if resC.Error != nil {
		fmt.Println("Cart not found")
		return
		//pop-up error message to the display
	}
}

func ProductChecking(TagID string) {
	var product models.Product
	resP := initializers.DB.Where("nfc_tag = ?", TagID).First(&product)
	if resP.Error != nil {
		fmt.Println("Product not found")
		return
	}
	//pop-up error message to the display
}
