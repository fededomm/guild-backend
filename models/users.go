package models

type User struct {
	ID        int         `json:"id" yaml:"id" mapstructure:"id"`
	Name      string      `json:"name" yaml:"name" mapstructure:"name"`
	Surname   string      `json:"surname" yaml:"surname" mapstructure:"surname"`
	Username  string      `json:"username" yaml:"username" mapstructure:"username"`
	BattleTag string      `json:"battle_tag" yaml:"battle_tag" mapstructure:"battle_tag"`
	Pg        Personaggio `json:"pg" yaml:"pg" mapstructure:"pg"`
}

type Personaggio struct {
	ID            int    `json:"id" yaml:"id" mapstructure:"id"`
	UserID        int    `json:"user_id" yaml:"user_id" mapstructure:"user_id"`
	Name          string `json:"name" yaml:"name" mapstructure:"name"`
	Class         string `json:"class" yaml:"class" mapstructure:"class" validate:"class"`
	TierSetPieces int    `json:"tier_set_pieces" yaml:"tier_set_pieces" mapstructure:"tier_set_pieces" validate:"lte=4"`
	Rank          string `json:"rank" yaml:"rank" mapstructure:"rank" validate:"ranking"`
}
