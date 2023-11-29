package main

import (
	"apocalypse/database"
	"apocalypse/utils"
	_ "embed"
	"fmt"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"gopkg.in/yaml.v2"
)

type GlobalConfig struct {
	Log            LogConfig       `json:"log" yaml:"log" mapstructure:"log"`
	DataBaseConfig database.DbInfo `json:"database" yaml:"database" mapstructure:"database"`
}

type LogConfig struct {
	Level      string `json:"level" yaml:"level" mapstructure:"level"`
	EnableJSON bool   `json:"enable_json" yaml:"enable_json" mapstructure:"enable_json"`
}

var DefaultConfig = GlobalConfig{
	Log: LogConfig{
		Level:      "debug",
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

// Default config file.
//
//go:embed config.yaml
var projectConfigFile []byte

const ConfigFileEnvVar = "BOOK_STORE_BE_FILE_PATH"
const ConfigurationName = "BOOK_STORE_BE"

func ReadConfig() (*GlobalConfig, error) {

	configPath := os.Getenv(ConfigFileEnvVar)
	var cfgContent []byte
	var err error
	if configPath != "" {
		if _, err := os.Stat(configPath); err == nil {
			log.Info().Str("cfg-file-name", configPath).Msg("reading config")
			cfgContent, err = utils.ReadFileAndResolveEnvVars(configPath)
			log.Info().Msg("++++CFG:" + string(cfgContent))
			if err != nil {
				return nil, err
			}
		} else {
			return nil, fmt.Errorf("the %s env variable has been set but no file cannot be found at %s", ConfigFileEnvVar, configPath)
		}
	} else {
		log.Warn().Msgf("The config path variable %s has not been set. Reverting to bundled configuration", ConfigFileEnvVar)
		cfgContent = utils.ResolveConfigValueToByteArray(projectConfigFile)
		// return nil, fmt.Errorf("the config path variable %s has not been set; please set", ConfigFileEnvVar)
	}

	appCfg := DefaultConfig
	err = yaml.Unmarshal(cfgContent, &appCfg)
	if err != nil {
		log.Fatal().Err(err).Send()
	}

	if !appCfg.Log.EnableJSON {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}

	zerolog.SetGlobalLevel(zerolog.Level('0'))

	return &appCfg, nil
}
