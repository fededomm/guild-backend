package models

type WoWClass int

const (
	Warrior WoWClass = iota
	Paladin
	Hunter
	Rogue
	Priest
	DeathKnight
	Shaman
	Mage
	Warlock
	Monk
	Druid
	DemonHunter
	Evoker
)

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

func (c WoWClass) String() string {
	if c < 0 || int(c) >= len(wowClassNames) {
		return "Unknown"
	}
	return wowClassNames[c]
}
