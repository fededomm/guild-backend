package models

type Users struct {
	ID        int      `json:"id" yaml:"id" mapstructure:"id"`
	Name      string   `json:"name" yaml:"name" mapstructure:"name"`
	Surname   string   `json:"surname" yaml:"surname" mapstructure:"surname"`
	Username  string   `json:"username" yaml:"username" mapstructure:"username"`
	Class     WoWClass `json:"class" yaml:"class" mapstructure:"class"`
	BattleTag string   `json:"battle_tag" yaml:"battle_tag" mapstructure:"battle_tag"`
}

