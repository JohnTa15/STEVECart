package controllers

import (
	"github.com/gin-gonic/gin"
)

func GetLatestNFC(c *gin.Context) {
	cartID := c.Query("cartID")

	_, err := CartChecking(cartID)
	if err != nil {
		c.JSON(500, gin.H{"status": "error", "message": "Cart not found"})
		return
	}

	nfcData, err := GetNFC(cartID)
	if err != nil {
		c.JSON(500, gin.H{"status": "error", "message": "Failed to read NFC data"})
		return
	}

	if nfcData == "" {
		c.JSON(200, gin.H{
			"status": "waiting",
			"message": "No NFC tag scanned recently",
			"nfc_tag": "",
		})
	} else {
		c.JSON(200, gin.H{
			"status": "success",
			"message": "NFC tag retrieved successfully",
			"nfc_tag": nfcData,
		})
	}
}
