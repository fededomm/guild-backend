package config

import "guild-be/src/database"

type GlobalConfig struct {
	Log            LogConfig       `json:"log" yaml:"log" mapstructure:"log"`
	DataBaseConfig database.DbInfo `json:"database" yaml:"database" mapstructure:"database"`
	Ranking        []string        `json:"ranking" yaml:"ranking" mapstructure:"ranking"`
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
}
