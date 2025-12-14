package controllers

import (
	"form/config"
	"form/models"

	"github.com/gin-gonic/gin"
)

func SaveKycData(c *gin.Context) {
	var body models.KycData
	if err := c.BindJSON(&body); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	config.DB.Create(&body)
	c.JSON(200, gin.H{"id": body.ID})
}
