package rest

import (
	"apocalypse/models"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

func CustomValidatorGin(model interface{}, c *gin.Context) error {
	var val = validator.New()
	if err := val.RegisterValidation("class", models.WowClassValidator); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return err
	}
	if err := val.RegisterValidation("ranking", models.RankingValidator); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return err
	}
	if err := val.Struct(model); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return err
	}
	return nil
}
