package utils

import "github.com/go-playground/validator"

type ArrValidation []string

func (r *ArrValidation) StringArrayValidator(fl validator.FieldLevel) bool {
	rank := fl.Field().String()
	for _, validRank := range *r {
		if rank == validRank {
			return true
		}
	}
	return false
}
