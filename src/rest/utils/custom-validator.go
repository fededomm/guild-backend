package utils

import (
	"github.com/go-playground/validator"
	"github.com/rs/zerolog/log"
)

type ArrToValid struct {
	Rank  []string
	Class []string
}


func (arr *ArrToValid) CustomArrayRankClassValidatorGin(model interface{}, val *validator.Validate) error {
	var r = ArrValidation(arr.Rank)
	var c = ArrValidation(arr.Class)
	if err := customValidator(val, r, "rank"); err != nil {
		log.Err(err).Msg(err.Error())
		return err
	}
	if err := customValidator(val, c, "class"); err != nil {
		log.Err(err).Msg(err.Error())
		return err
	}
	if err := val.Struct(model); err != nil {
		log.Err(err).Msg(err.Error())
		return err
	}
	return nil
}


func customValidator(val *validator.Validate, r ArrValidation, field string) error {
	if err := val.RegisterValidation(field, r.StringArrayValidator); err != nil {
		log.Err(err).Msg(err.Error())
		return err
	}
	return nil
}
