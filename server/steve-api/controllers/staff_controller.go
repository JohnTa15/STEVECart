package controllers

import (
	"errors"
	"net/http"
	"steve-api/initializers"
	"steve-api/models"
	"strconv"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func AdminLogin(c *gin.Context) {
	var body struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	if err := initializers.DB.Where("email = ?", body.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email or password"})
		return
	}

	if user.Role != "admin" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Access denied: user is not an administrator"})
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(body.Password))
	if err != nil {
		// Plain text fallback
		if user.PasswordHash != body.Password {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
			return
		}
	}
	
	// On success, return details
	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"role":    user.Role,
		"email":   user.Email,
	})
}

// to register products update prices etc etc
func addProduct(productName string, productCategory string, nfcTag string, productDescription string, weight float64, pcs int, price float64, shelveID uint64, role string) error {
	if role != "admin" {
		return errors.New("unauthorized: admin role required")
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
	if err := initializers.DB.Save(&product).Error; err != nil {
		return err
	}
	return nil
}

func deleteProduct(nfcTag string, role string) error {
	if role != "admin" {
		return errors.New("unauthorized: admin role required")
	}
	initializers.ConnectDB()
	var product models.Product
	if err := initializers.DB.Where("nfc_tag = ?", nfcTag).Delete(&product).Error; err != nil {
		return err
	}
	return nil
}

func updateProduct(nfcTag string, productName string, productCategory string, productDescription string, weight float64, pcs int, price float64, role string, shelveID uint64) error {
	if role != "admin" {
		return errors.New("unauthorized: admin role required")
	}
	initializers.ConnectDB()
	var product models.Product
	if err := initializers.DB.Where("nfc_tag = ?", nfcTag).First(&product).Error; err != nil {
		return err
	}
	product.ProductName = productName
	product.ProductCategory = productCategory
	product.ProductDescription = productDescription
	product.Weight = weight
	product.Pcs = pcs
	product.Price = price
	product.ShelveID = shelveID
	if err := initializers.DB.Save(&product).Error; err != nil {
		return err
	}
	return nil
}

// managing carts
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

// //managing users/admins
func setAdmin(userID string, role string) error {
	if role != "admin" {
		return errors.New("unauthorized: only administrators can promote users")
	}
	initializers.ConnectDB()

	var user models.User
	if err := initializers.DB.Where("id = ?", userID).First(&user).Error; err != nil {
		return errors.New("user not found: " + err.Error())
	}

	user.Role = "admin"
	if err := initializers.DB.Save(&user).Error; err != nil {
		return errors.New("failed to promote user to admin: " + err.Error())
	}

	return nil
}

// managing shelve positions
func addShelvePosition(shelveID uint64, xCoord float64, yCoord float64, description string, role string) error {
	if role != "admin" {
		return errors.New("unauthorized: admin role required")
	}
	initializers.ConnectDB()
	var shelvePosition models.Shelves
	shelvePosition.ShelveID = shelveID
	shelvePosition.X_Coordinate = xCoord
	shelvePosition.Y_Coordinate = yCoord
	shelvePosition.Description = description
	if err := initializers.DB.Save(&shelvePosition).Error; err != nil {
		return err
	}
	return nil
}

func deleteShelvePosition(shelveID uint64, role string) error {
	if role != "admin" {
		return errors.New("unauthorized: admin role required")
	}
	initializers.ConnectDB()
	var shelvePosition models.Shelves
	if err := initializers.DB.Where("shelve_id = ?", shelveID).Delete(&shelvePosition).Error; err != nil {
		return err
	}
	return nil
}

func updateShelvePosition(shelveID uint64, xCoord float64, yCoord float64, description string, role string) error {
	if role != "admin" {
		return errors.New("unauthorized: admin role required")
	}
	initializers.ConnectDB()
	var shelvePosition models.Shelves
	if err := initializers.DB.Where("shelve_id = ?", shelveID).First(&shelvePosition).Error; err != nil {
		return err
	}
	shelvePosition.X_Coordinate = xCoord
	shelvePosition.Y_Coordinate = yCoord
	shelvePosition.Description = description
	if err := initializers.DB.Save(&shelvePosition).Error; err != nil {
		return err
	}
	return nil
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
		Weight             float64 `json:"weight"`
		Pcs                int     `json:"pcs"`
		Price              float64 `json:"price"`
		ShelveID           uint64  `json:"shelve_id"`
		Role               string  `json:"role"`
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := addProduct(body.ProductName, body.ProductCategory, body.NFCTag, body.ProductDescription, body.Weight, body.Pcs, body.Price, body.ShelveID, body.Role); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Product added"})
}

func UpdateProduct(c *gin.Context) {
	var body struct {
		NFCTag             string  `json:"nfc_tag"`
		ProductName        string  `json:"product_name"`
		ProductCategory    string  `json:"product_category"`
		ProductDescription string  `json:"product_description"`
		Weight             float64 `json:"weight"`
		Pcs                int     `json:"pcs"`
		Price              float64 `json:"price"`
		ShelveID           uint64  `json:"shelve_id"`
		Role               string  `json:"role"`
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := updateProduct(body.NFCTag, body.ProductName, body.ProductCategory, body.ProductDescription, body.Weight, body.Pcs, body.Price, body.Role, body.ShelveID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Product updated"})
}

func DeleteProduct(c *gin.Context) {
	nfcTag := c.Query("nfcTag")
	role := c.Query("role")
	if nfcTag == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "nfcTag is required"})
		return
	}
	if err := deleteProduct(nfcTag, role); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Product deleted"})
}

func GetAllShelves(c *gin.Context) {
	var shelves []models.Shelves
	if err := initializers.DB.Find(&shelves).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": shelves})
}

func AddShelvePosition(c *gin.Context) {
	var body struct {
		ShelveID    uint64  `json:"shelve_id"`
		XCoord      float64 `json:"x_coord"`
		YCoord      float64 `json:"y_coord"`
		Description string  `json:"description"`
		Role        string  `json:"role"`
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := addShelvePosition(body.ShelveID, body.XCoord, body.YCoord, body.Description, body.Role); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Shelve added"})
}

func DeleteShelvePosition(c *gin.Context) {
	shelveIDStr := c.Query("shelveID")
	role := c.Query("role")
	if shelveIDStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "shelveID is required"})
		return
	}
	shelveID, err := strconv.ParseUint(shelveIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid shelveID format"})
		return
	}
	if err := deleteShelvePosition(shelveID, role); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Shelve deleted"})
}

func UpdateShelvePosition(c *gin.Context) {
	var body struct {
		ShelveID    uint64  `json:"shelve_id"`
		XCoord      float64 `json:"x_coord"`
		YCoord      float64 `json:"y_coord"`
		Description string  `json:"description"`
		Role        string  `json:"role"`
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := updateShelvePosition(body.ShelveID, body.XCoord, body.YCoord, body.Description, body.Role); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Shelve updated"})
}

func SetAdmin(c *gin.Context) {
	var body struct {
		UserID string `json:"user_id"`
		Role   string `json:"role"`
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := setAdmin(body.UserID, body.Role); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "User successfully promoted to Admin"})
}
