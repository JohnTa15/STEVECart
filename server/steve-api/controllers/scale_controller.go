package controllers

import (
	"fmt"
	"steve-api/controllers"
	"steve-api/models"
)


func MeasureWeight(weight float64){
	var product models.Product
	var cart models.Cart

	controllers.CartChecking(cart.Cart_ID)
	controllers.ProductChecking(product.NFCTag)
		
}