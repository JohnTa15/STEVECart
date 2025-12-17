package controllers

import (
	"github.com/gin-gonic/gin",
	"net/http",
)

func PostsCreate(c *gin.Context) {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "maou",
		})
	})
	r.Run(":8089")
}