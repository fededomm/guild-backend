package models

import "github.com/go-playground/validator"

type Rank []string

func (r *Rank)RankingValidator(fl validator.FieldLevel) bool {
    rank := fl.Field().String()
    for _, validRank := range *r {
        if rank == validRank {
            return true
        }
    }
    return false
}