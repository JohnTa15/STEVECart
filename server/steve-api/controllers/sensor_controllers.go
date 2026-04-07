package controllers

import (
	"fmt"
	"steve-api/initializers"
	"steve-api/models"
)

func NFCScan(tagID string) {
	var product models.Product

	res := initializers.DB.Where("nfc_tag = ?", tagID).First(&product)
	if res.Error != nil {
		fmt.Println("Product not found")
		return
	}
	fmt.Println("Added: %s to Cart!\n", product.ProductName)
	SumPriceProducts(product.Price)
}

func SumPriceProducts(price float32) {
}
