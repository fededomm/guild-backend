package rest

import (
	"guild-be/models"
	"github.com/go-playground/validator"
)

func CustomValidatorGin(model interface{}, rank []string) error {
	var val = validator.New()
	var r models.Rank = rank

	if err := RankCustomValidator(val, r); err != nil {
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

func RankCustomValidator(val *validator.Validate, r models.Rank) error{
	if err := val.RegisterValidation("ranking", r.RankingValidator); err != nil {
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
