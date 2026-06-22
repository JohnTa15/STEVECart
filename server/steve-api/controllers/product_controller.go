package controllers

import (
	"fmt"
	"steve-api/initializers"
	"steve-api/models"

	"github.com/gin-gonic/gin"
)

func updateRemainingProduct(productID uint, totalPcs int) {
	var totalInCarts int64
	initializers.DB.Model(&models.UserCartItem{}).Where("product_id = ?", productID).Select("COALESCE(SUM(quantity), 0)").Row().Scan(&totalInCarts)

	remainingQty := totalPcs - int(totalInCarts)
	if remainingQty < 0 {
		remainingQty = 0
	}

	var remainingProduct models.RemainingProduct
	err := initializers.DB.Where("product_id = ?", productID).First(&remainingProduct).Error
	if err == nil {
		remainingProduct.Quantity = remainingQty
		initializers.DB.Save(&remainingProduct)
	} else {
		remainingProduct.ProductID = productID
		remainingProduct.Quantity = remainingQty
		initializers.DB.Create(&remainingProduct)
	}
}

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
	
	// Add to user's cart items
	var userCartItem models.UserCartItem
	userCartItem.UserCartID = cart.ID
	userCartItem.ProductID = product.ID
	
	err = initializers.DB.Where("user_cart_id = ? AND product_id = ?", cart.ID, product.ID).First(&userCartItem).Error
	if err == nil {
		// increment quantity if already in the cart
		userCartItem.Quantity++
		initializers.DB.Save(&userCartItem)
	} else {
		// if the product is not in the cart, add it
		userCartItem.UserCartID = cart.ID
		userCartItem.ProductID = product.ID
		userCartItem.Quantity = 1
		initializers.DB.Save(&userCartItem)		
	}

	// Update remaining product statistics
	updateRemainingProduct(product.ID, product.Pcs)

	// Update cart total price and weight
	cart.TotalPrice += product.Price
	cart.NetWeight += product.Weight

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

	var userCartItem models.UserCartItem
	if err := initializers.DB.Where("user_cart_id = ? AND product_id = ?", cart.ID, product.ID).First(&userCartItem).Error; err == nil {
		// Update user cart item quantity
		if userCartItem.Quantity > 1 {
			userCartItem.Quantity--
			initializers.DB.Save(&userCartItem)
		} else {
			initializers.DB.Delete(&userCartItem)
		}

		// Update remaining product statistics
		updateRemainingProduct(product.ID, product.Pcs)

		// Update cart total price and weight
		cart.TotalPrice -= product.Price
		cart.NetWeight -= product.Weight
		if cart.TotalPrice < 0 {
			cart.TotalPrice = 0
		}
		if cart.NetWeight < 0 {
			cart.NetWeight = 0
		}
		initializers.DB.Save(&cart)

		fmt.Printf("Success! Removed product %s from Cart %s. New total price: %.2f\n", product.ProductName, cart.Cart_ID, cart.TotalPrice)
	} else {
		fmt.Println("Error: Product not in cart.")
	}
}

func GetCartItems(c *gin.Context) {
	cartID := c.Query("cartID")
	if cartID == "" {
		c.JSON(400, gin.H{"error": "cartID is required"})
		return
	}

	cart, err := CartChecking(cartID)
	if err != nil {
		c.JSON(404, gin.H{"error": "Cart not found"})
		return
	}

	type CartItemResponse struct {
		Name          string  `json:"name"`
		CurrentNFCTag string  `json:"current_nfc_tag"`
		Weight        string  `json:"weight"`
		Price         string  `json:"price"`
		Capacity      int     `json:"capacity"`
		Status        string  `json:"status"`
	}

	var results []CartItemResponse
	query := `
		SELECT p.product_name as name, p.nfc_tag as current_nfc_tag, 
		       CONCAT(p.weight, ' kg') as weight, CONCAT(p.price, ' €') as price, 
		       uci.quantity as capacity, 'Success' as status
		FROM user_cart_items uci
		JOIN products p ON uci.product_id = p.id
		WHERE uci.user_cart_id = ?
	`
	if err := initializers.DB.Raw(query, cart.ID).Scan(&results).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to fetch cart items: " + err.Error()})
		return
	}

	c.JSON(200, gin.H{"status": 200, "data": results})
}
