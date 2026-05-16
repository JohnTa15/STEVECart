package controllers

import (
	"fmt"
	"steve-api/initializers"
	"steve-api/models"
)

//user management
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

func deleteUser(email string){
	initializers.ConnectDB()
	var user models.User
	user.Email = email
	if(user.Role == "admin"){
		fmt.Println("Are you sure about that?")
		fmt.Println("I mean you're deleting an admin")
	} else {
		fmt.Println("Are you sure? ")
	}
	fmt.Println("Press 'y' to delete")
	var input string
	fmt.Scanln(&input)
	if input == "y" {
		initializers.DB.Delete(&user)
	}
}
