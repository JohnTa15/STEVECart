package controllers

import (
	"fmt"
	"steve-api/initializers"
	"steve-api/models"
)

func register(email string, password string) {
	initializers.ConnectDB()
	var user models.User
	user.Email = email
	user.PasswordHash = password
	initializers.DB.Save(&user)
}

func login(email string, password string) {
	initializers.ConnectDB()
	var user models.User
	user.Email = email
	user.PasswordHash = password
	initializers.DB.Save(&user)
}
