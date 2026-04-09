package controllers

import (
	"fmt"
	"steve-api/models"
	"steve-api/controllers"
)

func NFCScan(tagID string) {
	var product models.Product
	var cart models.Cart

	controllers.CartChecking(cart.Cart_ID)
	controllers.NFCChecking(product.NFCTag)
	SumPriceProducts(product.Price, cart.TotalPrice)
}

func SumPriceProducts(price float32, totalPrice float32) {
	totalPrice += price
	fmt.Println("Total Price: %f\n", totalPrice)
}