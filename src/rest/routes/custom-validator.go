package routes

import (
	"guild-be/src/rest/utils"

	"github.com/go-playground/validator"
)

type ArrToValid struct {
	Rank  []string
	Class []string
	Name  []string
}

func (arr *ArrToValid) CustomArrayNamesValidator(model interface{}, val *validator.Validate) error {
	var n = utils.ArrValidation(arr.Name)
	if err := RankCustomValidator(val, n); err != nil {
		return err
	}
	if err := val.Struct(model); err != nil {
		return err
	}
	return nil
}

func (arr *ArrToValid) CustomArrayValidatorGin(model interface{}, val *validator.Validate) error {
	var r = utils.ArrValidation(arr.Rank)
	var c = utils.ArrValidation(arr.Class)
	if err := RankCustomValidator(val, r); err != nil {
		return err
	}
	if err := ClassCustomValidator(val, c); err != nil {
		return err
	}
	if err := val.Struct(model); err != nil {
		return err
	}
	return nil
}

func RankCustomValidator(val *validator.Validate, r utils.ArrValidation) error {
	if err := val.RegisterValidation("ranking", r.StringArrayValidator); err != nil {
		return err
	}
	return nil
}

func ClassCustomValidator(val *validator.Validate, r utils.ArrValidation) error {
	if err := val.RegisterValidation("class", r.StringArrayValidator); err != nil {
		return err
	}
	return nil
}

func NameCustomValidator(val *validator.Validate, r utils.ArrValidation) error {
	if err := val.RegisterValidation("name", r.StringArrayValidator); err != nil {
		return err
	}
	return nil
}
