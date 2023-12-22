package utils

import (
	"github.com/go-playground/validator"
)

type ArrToValid struct {
	Rank  []string
	Class []string
}


func (arr *ArrToValid) CustomArrayRankClassValidatorGin(model interface{}, val *validator.Validate) error {
	//cast type to ArrayValidation
	var r = ArrValidation(arr.Rank)
	var c = ArrValidation(arr.Class)
	if err := customValidator(val, r, "rank"); err != nil {
		return err
	}
	if err := customValidator(val, c, "class"); err != nil {
		return err
	}
	if err := val.Struct(model); err != nil {
		return err
	}
	return nil
}


func customValidator(val *validator.Validate, arr ArrValidation, field string) error {
	if err := val.RegisterValidation(field, arr.StringArrayValidator); err != nil {
		return err
	}
	return nil
}
