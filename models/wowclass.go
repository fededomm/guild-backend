package models

import "github.com/go-playground/validator"


var wowClassNames = [...]string{
	"Warrior",
	"Paladin",
	"Hunter",
	"Rogue",
	"Priest",
	"Death Knight",
	"Shaman",
	"Mage",
	"Warlock",
	"Monk",
	"Druid",
	"Demon Hunter",
	"Evoker",
}

func WowClassValidator(fl validator.FieldLevel) bool {
	class := fl.Field().String()
    for _, validClass := range wowClassNames {
        if class == validClass {
            return true
        }
    }
    return false
}