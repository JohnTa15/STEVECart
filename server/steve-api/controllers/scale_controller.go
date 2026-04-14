package controllers

import (
	"steve-api/models"
)

//scaling the products and then compare them with the database(influxdb <-> mariadb)
func MeasureWeight(weight float64){
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
		
}