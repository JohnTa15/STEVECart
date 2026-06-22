package migrate

import (
	"steve-api/initializers"
	"steve-api/models"
)

func init() {
	initializers.LoadENV()
	initializers.ConnectDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Shelves{})
	initializers.DB.AutoMigrate(&models.Product{})
	initializers.DB.AutoMigrate(&models.Cart{})
	initializers.DB.AutoMigrate(&models.User{})
	initializers.DB.AutoMigrate(&models.UserCartSession{})
	initializers.DB.AutoMigrate(&models.UserCartItem{})
	initializers.DB.AutoMigrate(&models.RemainingProduct{})
}
