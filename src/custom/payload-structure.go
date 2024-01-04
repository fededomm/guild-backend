package custom

type ExampleBodyUser struct {
	ID        int    `json:"id" yaml:"id" mapstructure:"id"`
	Name      string `json:"name" yaml:"name" mapstructure:"name" validate:"required"`
	Surname   string `json:"surname" yaml:"surname" mapstructure:"surname" validate:"required"`
	Username  string `json:"username" yaml:"username" mapstructure:"username" validate:"required"`
	BattleTag string `json:"battle_tag" yaml:"battle_tag" mapstructure:"battle_tag" validate:"required"`
}

type ExampleBodyPg struct {
	ID            int    `json:"id,omitempty" yaml:"id" mapstructure:"id"`
	Name          string `json:"name,omitempty" yaml:"name" mapstructure:"name" validate:"required"`
	Class         string `json:"class,omitempty" yaml:"class" mapstructure:"class" validate:"required, class"`
	TierSetPieces int    `json:"tiersetpieces,omitempty" yaml:"tier_set_pieces" mapstructure:"tier_set_pieces" validate:"required, lte=4"`
	Rank          string `json:"rank,omitempty" yaml:"rank" mapstructure:"rank" validate:"required, ranking"`
}

type ExampleListOfPGOfAUser struct {
	ID        int             `json:"id,omitempty" yaml:"id" mapstructure:"id"`
	Name      string          `json:"name" yaml:"name" mapstructure:"name"`
	Username  string          `json:"username" yaml:"username" mapstructure:"username"`
	BattleTag string          `json:"battle_tag" yaml:"battle_tag" mapstructure:"battle_tag"`
	PgList    []ExampleBodyPg `json:"pg_list,omitempty" yaml:"pg_list" mapstructure:"pg_list"`
}
