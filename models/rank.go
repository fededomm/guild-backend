package models

import "github.com/go-playground/validator"

var ranking = [...]string{
    "Morte",
    "Carestia",
    "Pestilenza",
    "Guerra",
    "Drago",
    "Immortale",
    "Bestia del mare",
    "Profeta",
    "Scudiero",
    "Limbo",
}

func RankingValidator(fl validator.FieldLevel) bool {
    rank := fl.Field().String()
    for _, validRank := range ranking {
        if rank == validRank {
            return true
        }
    }
    return false
}