package controllers

import (
	"fmt"
	"steve-api/initializers"
	"steve-api/models"
)

// user management
func registerUser(email string, password string) {
	initializers.ConnectDB()
	var user models.User
	user.Email = email
	user.PasswordHash = password
	initializers.DB.Save(&user)
}

func loginUser(email string, password string) {
	initializers.ConnectDB()
	var user models.User
	user.Email = email
	user.PasswordHash = password
	initializers.DB.Save(&user)
}

func deleteUser(email string) {
	initializers.ConnectDB()
	var user models.User
	err := initializers.DB.Where("email = ?", email).First(&user).Error

	var admin models.Admin
	isAdmin := false
	if err != nil {
		// If not found in users table, check the admins table
		if errAdmin := initializers.DB.Where("email = ?", email).First(&admin).Error; errAdmin == nil {
			isAdmin = true
		} else {
			fmt.Println("[ERROR]Account not found")
			return
		}
	}

	if isAdmin {
		fmt.Println("[WARNING] Are you sure about that?")
		fmt.Println("[INFO] I mean you're deleting an admin")
	} else {
		fmt.Println("[INFO] Are you sure? ")
	}
	fmt.Println("[INFO] Press 'y' to delete")
	var input string
	fmt.Scanln(&input)
	if input == "y" {
		if isAdmin {
			fmt.Println("[WARNING] Admin is deleting another admin!")
			initializers.DB.Delete(&admin)
		} else {
			fmt.Println("[INFO] USER deleted successfully!")
			initializers.DB.Delete(&user)
		}
	}
}

// product search
func searchProduct(name string, nfc_tag string) {
	initializers.ConnectDB()
	var product models.Product
	err := initializers.DB.Where("product_name LIKE ? OR nfc_tag LIKE ?", name, nfc_tag).Find(&product)
	if err != nil {
		fmt.Println("[ERROR] Product not found")
		return
	}
	fmt.Println(product)
}
