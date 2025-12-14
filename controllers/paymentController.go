package controllers

import (
	"form/config"
	"form/models"

	"github.com/gin-gonic/gin"
)

func SavePayment(c *gin.Context) {
	var body models.Payment
	if err := c.BindJSON(&body); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	config.DB.Create(&body)
	c.JSON(200, gin.H{"success": true})
}
