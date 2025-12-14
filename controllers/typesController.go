package controllers

import (
	"form/config"
	"form/models"

	"github.com/gin-gonic/gin"
)

func GetTypes(c *gin.Context) {
	var t []models.KycType
	config.DB.Find(&t)
	c.JSON(200, t)
}
