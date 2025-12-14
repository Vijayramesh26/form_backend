package controllers

import (
	"form/config"
	"form/models"

	"github.com/gin-gonic/gin"
)

func GetFields(c *gin.Context) {
	typeKey := c.Query("type_key")

	var fields []models.KycField
	config.DB.Order("sequence asc").Find(&fields)

	var rules []models.KycTypeFieldRule
	config.DB.Where("type_key = ?", typeKey).Find(&rules)

	override := make(map[string]int)
	for _, r := range rules {
		override[r.FieldKey] = r.Required
	}

	output := []gin.H{}

	for _, f := range fields {
		req := f.BaseRequired
		if v, ok := override[f.FieldKey]; ok {
			req = v
		}

		output = append(output, gin.H{
			"step_key":  f.StepKey,
			"field_key": f.FieldKey,
			"label":     f.Label,
			"type":      f.Type,
			"items":     f.Items,
			"required":  req,
			"sequence":  f.Sequence,
		})
	}

	c.JSON(200, output)
}
