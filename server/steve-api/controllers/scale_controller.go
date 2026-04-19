package controllers

import (
	"log"
	"steve-api/models"
)

//scaling the products and then compare them with the database(influxdb <-> mariadb)
func MeasureWeight(cartID string){
	var product models.Product
	var cart models.Cart

	cart, err := CartChecking(cart.Cart_ID)
	if err != nil {
		return 
	}

	product, error := ProductChecking(product.NFCTag)
	if error != nil {
		return
	}

	weight, error := GetWeight(cart.Cart_ID)
		
}