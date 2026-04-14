package controllers

import(
	"steve-api/models"
	"steve-api/initializers"
)
//to register products update prices etc etc
func addProduct(productName string, productCategory string, nfcTag string, productDescription string, weight float32, pcs int, price float32) {
	initializers.ConnectDB()
	var product models.Product
	product.ProductName = productName
	product.ProductCategory = productCategory
	product.NFCTag = nfcTag
	product.ProductDescription = productDescription
	product.Weight = weight
	product.Pcs = pcs
	product.Price = price
	initializers.DB.Save(&product)
}

func deleteProduct(nfcTag string, role string) {
	if role != "admin" {
		return
	}
	initializers.ConnectDB()
	var product models.Product
	initializers.DB.Where("nfc_tag = ?", nfcTag).Delete(&product)
}

func updateProduct(nfcTag string, productName string, productCategory string, productDescription string, weight float32, pcs int, price float32, role string) {
	if role != "admin" {
		return
	}
	initializers.ConnectDB()
	var product models.Product
	initializers.DB.Where("nfc_tag = ?", nfcTag).First(&product)
	product.ProductName = productName
	product.ProductCategory = productCategory
	product.ProductDescription = productDescription
	product.Weight = weight
	product.Pcs = pcs
	product.Price = price
	initializers.DB.Save(&product)
}

//managing carts
func addCart(cartID string, role string) {
	if role != "admin" {
		return
	}
	initializers.ConnectDB()
	var cart models.Cart
	cart.Cart_ID = cartID
	initializers.DB.Save(&cart)
}

func deleteCart(cartID string, role string) {
	if role != "admin" {
		return
	}
	initializers.ConnectDB()
	var cart models.Cart
	initializers.DB.Where("cart_id = ?", cartID).Delete(&cart)
}

//managing users/admins
func setAdmin(userID string, role string) {
	if role != "admin" {
		return
	}
	initializers.ConnectDB()
	var user models.User
	initializers.DB.Where("id = ?", userID).First(&user)
	user.Role = "admin"
	initializers.DB.Save(&user)
}

func deleteUser(userID string, role string) {
	if role != "admin" {
		return
	}
	initializers.ConnectDB()
	var user models.User
	initializers.DB.Where("id = ?", userID).Delete(&user)
}