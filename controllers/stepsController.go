package controllers

import (
	"form/config"
	"form/models"

	"github.com/gin-gonic/gin"
)

func GetSteps(c *gin.Context) {
	var steps []models.KycStep
	config.DB.Order("sequence asc").Find(&steps)
	c.JSON(200, steps)
}
