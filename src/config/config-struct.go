package config

import (
	"guild-be/src/database"
	"guild-be/src/observability"
)
type GlobalConfig struct {
	Log            LogConfig       `json:"log" yaml:"log" mapstructure:"log"`
	DataBaseConfig database.DbInfo `json:"database" yaml:"database" mapstructure:"database"`
	Observability  observability.Observability   `json:"observability" yaml:"observability" mapstructure:"observability"`
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
	Observability: observability.Observability{
		Enable:      false,
		Endpoint:    "127.0.0.1/4317",
		ServiceName: "guild-be",
	},
}
