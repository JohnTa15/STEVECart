package controllers

import (
	"steve-api/models"
	"steve-api/initializers"
	"steve-api/controllers"
)

func register(username string, password string, email string) {
	initializers.ConnectDB()
	var user models.User
	user.Username = username
	user.PasswordHash = password
	user.Email = email
	initializers.DB.Save(&user)
}

func login(username string, password string) {
	initializers.ConnectDB()
	var user models.User
	user.Username = username
	user.PasswordHash = password
	initializers.DB.Save(&user)
}

func addProdtoCart(nfcTag string, cartID string) {
	initializers.ConnectDB()
	var product models.Product
	var cart models.Cart
	controllers.CartChecking(cart.Cart_ID)
	controllers.ProductChecking(product.NFCTag)
	var cartOperatorItem models.CartOperatorItem
	cartOperatorItem.UserCartID = cart.ID
	cartOperatorItem.ProductID = product.ID
	initializers.DB.Save(&cartOperatorItem)
}

func removeProdfromCart(nfcTag string, cartID string) {
	initializers.ConnectDB()
	var product models.Product
	var cart models.Cart
	controllers.CartChecking(cart.Cart_ID)
	controllers.ProductChecking(product.NFCTag)
	var cartOperatorItem models.CartOperatorItem
	cartOperatorItem.UserCartID = cart.ID
	cartOperatorItem.ProductID = product.ID
	initializers.DB.Where("user_cart_id = ? AND product_id = ?", cartOperatorItem.UserCartID, cartOperatorItem.ProductID).Delete(&cartOperatorItem)
}