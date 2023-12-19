package config

import "guild-be/src/database"

type GlobalConfig struct {
	Log            LogConfig       `json:"log" yaml:"log" mapstructure:"log"`
	DataBaseConfig database.DbInfo `json:"database" yaml:"database" mapstructure:"database"`
	Ranking        []string        `json:"ranking" yaml:"ranking" mapstructure:"ranking"`
	ValidArray     ArrValidation   `json:"arrValidation" yaml:"arrValidation" mapstructure:"arrValidation"`
}

type ArrValidation struct {
	Rank  []string `json:"rank" yaml:"rank" mapstructure:"rank"`
	Class []string `json:"class" yaml:"class" mapstructure:"class"`
}

type LogConfig struct {
	Level      int  `json:"level" yaml:"level" mapstructure:"level"`
	EnableJSON bool `json:"enable_json" yaml:"enable_json" mapstructure:"enable_json"`
}

var DefaultConfig = GlobalConfig{
	Log: LogConfig{
		Level:      -1,
		EnableJSON: false,
	},
	DataBaseConfig: database.DbInfo{
		Host:     "localhost",
		Port:     "5432",
		User:     "postgres",
		Password: "postgres",
		Sslmode:  "disable",
		Dbname:   "apocalypse",
	},
	ValidArray: ArrValidation{
		Rank: []string{
			"Guild Master",
			"Officer",
			"Raider",
			"Trial",
			"Social",
			"Alt",
		},
		Class: []string{
			"Death Knight",
			"Demon Hunter",
			"Druid",
			"Hunter",
			"Mage",
			"Monk",
			"Paladin",
			"Priest",
			"Rogue",
			"Shaman",
			"Warlock",
			"Warrior",
		},
	},
}