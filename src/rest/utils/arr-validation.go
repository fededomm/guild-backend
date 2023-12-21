package utils

import (
	"github.com/go-playground/validator"
	"github.com/rs/zerolog/log"
)

type ArrValidation []string

func (arr *ArrValidation) StringArrayValidator(fl validator.FieldLevel) bool {
	field := fl.Field().String()
	for _, validField := range *arr {
		if field == validField {
			log.Info().Msgf("valid field: %s", field)
			return true
		}
	}
	return false
}
