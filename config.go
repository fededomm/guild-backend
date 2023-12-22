package main

import (
	_ "embed"
	"fmt"
	"guild-be/src/config"
	"guild-be/src/rest/utils"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"gopkg.in/yaml.v2"
)

// Default config file.
//
//go:embed config.yaml
var projectConfigFile []byte

const ConfigFileEnvVar = "GUILD_BACKEND_FILE_PATH"
const ConfigurationName = "GUILD_BACKEND"

func ReadConfig() (*config.GlobalConfig, error) {

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

	appCfg := config.DefaultConfig
	err = yaml.Unmarshal(cfgContent, &appCfg)
	
	if err != nil {
		log.Fatal().Err(err).Msgf("Error unmarshalling config: %q", err)
	}
	if !appCfg.Log.EnableJSON {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}
	zerolog.SetGlobalLevel(zerolog.Level(appCfg.Log.Level))
	return &appCfg, nil
}
