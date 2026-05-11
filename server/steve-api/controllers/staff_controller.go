package controllers

import (
	"net/http"
	"steve-api/initializers"
	"steve-api/models"

	"github.com/gin-gonic/gin"
)
//to register products update prices etc etc
func addProduct(productName string, productCategory string, nfcTag string, productDescription string, weight float32, pcs int, price float32, shelveID string, role string) {
	if role != "admin" {
		return
	}
	initializers.ConnectDB()
	var product models.Product
	product.ProductName = productName
	product.ProductCategory = productCategory
	product.NFCTag = nfcTag
	product.ProductDescription = productDescription
	product.Weight = weight
	product.Pcs = pcs
	product.Price = price
	product.ShelveID = shelveID
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

func updateProduct(nfcTag string, productName string, productCategory string, productDescription string, weight float32, pcs int, price float32, role string, shelveID string) {
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
	product.ShelveID = shelveID
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

//managing shelve positions
func addShelvePosition(shelveID string, xCoord float64, yCoord float64, description string, role string) {
	if role != "admin" {
		return
	}
	initializers.ConnectDB()
	var shelvePosition models.ShelvePosition
	shelvePosition.Shelve_id = shelveID
	shelvePosition.X_Coordinate = xCoord
	shelvePosition.Y_Coordinate = yCoord
	shelvePosition.Description = description
	initializers.DB.Save(&shelvePosition)
}

func deleteShelvePosition(shelveID string, role string) {
	if role != "admin" {
		return
	}
	initializers.ConnectDB()
	var shelvePosition models.ShelvePosition
	initializers.DB.Where("shelve_id = ?", shelveID).Delete(&shelvePosition)
}

func updateShelvePosition(shelveID string, xCoord float64, yCoord float64, description string, role string) {
	if role != "admin" {
		return
	}
	initializers.ConnectDB()
	var shelvePosition models.ShelvePosition
	initializers.DB.Where("shelve_id = ?", shelveID).First(&shelvePosition)
	shelvePosition.X_Coordinate = xCoord
	shelvePosition.Y_Coordinate = yCoord
	shelvePosition.Description = description
	initializers.DB.Save(&shelvePosition)
}

// --- Exported Gin HTTP Handlers ---

func GetAllProducts(c *gin.Context) {
	var products []models.Product
	if err := initializers.DB.Find(&products).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": products})
}

func AddProduct(c *gin.Context) {
	var body struct {
		ProductName        string  `json:"product_name"`
		ProductCategory    string  `json:"product_category"`
		NFCTag             string  `json:"nfc_tag"`
		ProductDescription string  `json:"product_description"`
		Weight             float32 `json:"weight"`
		Pcs                int     `json:"pcs"`
		Price              float32 `json:"price"`
		ShelveID           string  `json:"shelve_id"`
		Role               string  `json:"role"`
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	addProduct(body.ProductName, body.ProductCategory, body.NFCTag, body.ProductDescription, body.Weight, body.Pcs, body.Price, body.ShelveID, body.Role)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Product added"})
}

func UpdateProduct(c *gin.Context) {
	var body struct {
		NFCTag             string  `json:"nfc_tag"`
		ProductName        string  `json:"product_name"`
		ProductCategory    string  `json:"product_category"`
		ProductDescription string  `json:"product_description"`
		Weight             float32 `json:"weight"`
		Pcs                int     `json:"pcs"`
		Price              float32 `json:"price"`
		ShelveID           string  `json:"shelve_id"`
		Role               string  `json:"role"`
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	updateProduct(body.NFCTag, body.ProductName, body.ProductCategory, body.ProductDescription, body.Weight, body.Pcs, body.Price, body.Role, body.ShelveID)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Product updated"})
}

func DeleteProduct(c *gin.Context) {
	nfcTag := c.Query("nfcTag")
	role := c.Query("role")
	if nfcTag == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "nfcTag is required"})
		return
	}
	deleteProduct(nfcTag, role)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Product deleted"})
}

func GetAllShelves(c *gin.Context) {
	var shelves []models.ShelvePosition
	if err := initializers.DB.Find(&shelves).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": shelves})
}

func AddShelvePosition(c *gin.Context) {
	var body struct {
		ShelveID    string  `json:"shelve_id"`
		XCoord      float64 `json:"x_coord"`
		YCoord      float64 `json:"y_coord"`
		Description string  `json:"description"`
		Role        string  `json:"role"`
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	addShelvePosition(body.ShelveID, body.XCoord, body.YCoord, body.Description, body.Role)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Shelve added"})
}

func DeleteShelvePosition(c *gin.Context) {
	shelveID := c.Query("shelveID")
	role := c.Query("role")
	if shelveID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "shelveID is required"})
		return
	}
	deleteShelvePosition(shelveID, role)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Shelve deleted"})
}

func UpdateShelvePosition(c *gin.Context) {
	var body struct {
		ShelveID    string  `json:"shelve_id"`
		XCoord      float64 `json:"x_coord"`
		YCoord      float64 `json:"y_coord"`
		Description string  `json:"description"`
		Role        string  `json:"role"`
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	updateShelvePosition(body.ShelveID, body.XCoord, body.YCoord, body.Description, body.Role)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Shelve updated"})
}