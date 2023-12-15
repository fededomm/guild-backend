package rest

import (
	"apocalypse/models"
	"github.com/go-playground/validator"
)

func CustomValidatorGin(model interface{}) error {
	var val = validator.New()
	if err := RankCustomValidator(val); err != nil {
		return err
	}
	if err := ClassCustomValidator(val); err != nil {
		return err
	}
	if err := val.Struct(model); err != nil {
		return err
	}
	return nil
}

func RankCustomValidator(val *validator.Validate) error{
	if err := val.RegisterValidation("ranking", models.RankingValidator); err != nil {
		return err
	}
	return nil
}

func ClassCustomValidator(val *validator.Validate) error{
	if err := val.RegisterValidation("class", models.WowClassValidator); err != nil {
		return err
	}
	return nil
}
