package rest

import (
	"apocalypse/models"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"github.com/rs/zerolog/log"
)

func CustomValidatorGin(model interface{}, c *gin.Context) error {
	var val = validator.New()
	if err := val.RegisterValidation("class", models.WowClassValidator); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		log.Err(err).Msgf("Error registering class validator: %s", err.Error())
		return err
	}
	if err := val.RegisterValidation("ranking", models.RankingValidator); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		log.Err(err).Msgf("Error registering ranking validator: %s", err.Error())
		return err
	}
	if err := val.Struct(model); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		log.Err(err).Msgf("Error validating model: %s", err.Error())
		return err
	}
	return nil
}
