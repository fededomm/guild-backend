package config

import (
	"guild-be/src/database"
	"guild-be/src/observability"
)

type GlobalConfig struct {
	Log            LogConfig                   `json:"log" yaml:"log" mapstructure:"log"`
	DataBaseConfig database.DbInfo             `json:"database" yaml:"database" mapstructure:"database"`
	Observability  observability.Observability `json:"observability" yaml:"observability" mapstructure:"observability"`
	App            AppConfig                   `json:"app" yaml:"app" mapstructure:"app"`
}

type AppConfig struct {
	Server ServerConfig `json:"server" yaml:"server" mapstructure:"server"`
	CORS   CORSConfig   `json:"cors" yaml:"cors" mapstructure:"cors"`
}

type CORSConfig struct {
	AllowOrigins     []string `json:"allow_origins" yaml:"allow_origins" mapstructure:"allow_origins"`
	AllowMethods     []string `json:"allow_methods" yaml:"allow_methods" mapstructure:"allow_methods"`
	AllowHeaders     []string `json:"allow_headers" yaml:"allow_headers" mapstructure:"allow_headers"`
	AllowCredentials bool     `json:"allow_credentials" yaml:"allow_credentials" mapstructure:"allow_credentials"`
}

type ServerConfig struct {
	Port                string `json:"port" yaml:"port" mapstructure:"port"`
	Host                string `json:"host" yaml:"host" mapstructure:"host"`
	HealthCheckEndpoint string `json:"health_check_endpoint" yaml:"health_check_endpoint" mapstructure:"health_check_endpoint"`
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
	App: AppConfig{
		Server: ServerConfig{
			Port:                ":8080",
			Host:                "0.0.0.0",
			HealthCheckEndpoint: "/healthcheck",
		},
	},
}
