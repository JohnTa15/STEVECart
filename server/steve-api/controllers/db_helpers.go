package controllers

import (
	"fmt"
	"steve-api/initializers"
	"steve-api/models"
)

func CartChecking(cartID string) (models.Cart, error) {
	var cart models.Cart
	resC := initializers.DB.Where("cart_id = ?", cartID).First(&cart)
	if resC.Error != nil {
		fmt.Println("Cart not found")
		return cart, resC.Error
	}
	return cart, nil
}

func ProductChecking(TagID string) (models.Product, error) {
	var product models.Product
	resP := initializers.DB.Where("nfc_tag = ?", TagID).First(&product)
	if resP.Error != nil {
		fmt.Println("Product not found")
		return product, resP.Error
	}
	return product, nil
}
