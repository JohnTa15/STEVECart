package controllers

import (
	"fmt"
	"steve-api/initializers"
	"steve-api/models"
)

func AddProdtoCart(nfcTag string, cartID string) {
	if initializers.DB == nil {
		initializers.ConnectDB()
	}

	cart, err := CartChecking(cartID)
	if err != nil {
		return
	}

	product, err := ProductChecking(nfcTag)
	if err != nil {
		return
	}

	// Add to cart operator items
	var cartOperatorItem models.CartOperatorItem
	cartOperatorItem.UserCartID = cart.ID
	cartOperatorItem.ProductID = product.ID
	cartOperatorItem.Quantity = 1
	if err := initializers.DB.Save(&cartOperatorItem).Error; err != nil {
		fmt.Println("Error: Could not save cart operator item:", err)
		return
	}

	// Update cart total price and weight
	cart.TotalPrice += product.Price
	cart.TotalWeight += product.Weight

	if err := initializers.DB.Save(&cart).Error; err != nil {
		fmt.Println("Error: Could not update cart totals:", err)
		return
	}

	fmt.Printf("Success! Added product %s to Cart %s. New total price: %.2f\n", product.ProductName, cart.Cart_ID, cart.TotalPrice)
}

func RemoveProdfromCart(nfcTag string, cartID string) {
	if initializers.DB == nil {
		initializers.ConnectDB()
	}

	cart, err := CartChecking(cartID)
	if err != nil {
		return
	}

	product, err := ProductChecking(nfcTag)
	if err != nil {
		return
	}

	var cartOperatorItem models.CartOperatorItem
	if err := initializers.DB.Where("user_cart_id = ? AND product_id = ?", cart.ID, product.ID).First(&cartOperatorItem).Error; err == nil {
		initializers.DB.Delete(&cartOperatorItem)

		// Update cart total price and weight
		cart.TotalPrice -= product.Price
		cart.TotalWeight -= product.Weight
		if cart.TotalPrice < 0 {
			cart.TotalPrice = 0
		}
		if cart.TotalWeight < 0 {
			cart.TotalWeight = 0
		}
		initializers.DB.Save(&cart)

		fmt.Printf("Success! Removed product %s from Cart %s. New total price: %.2f\n", product.ProductName, cart.Cart_ID, cart.TotalPrice)
	} else {
		fmt.Println("Error: Product not in cart.")
	}
}
