package controllers

import (
	"form/config"
	"form/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SaveIdentity(c *gin.Context) {
	var body models.Identity
	c.BindJSON(&body)
	config.DB.Create(&body)
	c.JSON(200, gin.H{"success": true})
}

func SaveBank(c *gin.Context) {
	var bank models.Bank
	if err := c.BindJSON(&bank); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false})
		return
	}
	config.DB.Create(&bank)
	c.JSON(http.StatusOK, gin.H{"success": true})
}
func SaveAddress(c *gin.Context) {
	var address models.Address
	if err := c.BindJSON(&address); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false})
		return
	}
	config.DB.Create(&address)
	c.JSON(http.StatusOK, gin.H{"success": true})
}
func SaveAdditional(c *gin.Context) {
	var additional models.Additional
	if err := c.BindJSON(&additional); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false})
		return
	}

	config.DB.Create(&additional)
	c.JSON(http.StatusOK, gin.H{"success": true})
}

func SavePromoters(c *gin.Context) {
	var promoters []models.Promoter

	if err := c.BindJSON(&promoters); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false})
		return
	}

	for _, p := range promoters {
		config.DB.Create(&p)
	}

	c.JSON(http.StatusOK, gin.H{"success": true})
}

func SaveTrading(c *gin.Context) {
	var trading models.Trading
	if err := c.BindJSON(&trading); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false})
		return
	}

	config.DB.Create(&trading)
	c.JSON(http.StatusOK, gin.H{"success": true})
}
