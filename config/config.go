package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// App config struct
type (
	Config struct {
		Server  ServerConfig  `mapstructure:"Server"`
		MongoDB MongoDBConfig `mapstructure:"MongoDB"`
	}

	ServerConfig struct {
		Port string `mapstructure:"Port"`
	}

	MongoDBConfig struct {
		EnableURI bool   `mapstructure:"EnableURI"`
		URI       string `mapstructure:"URI"`
	}
)

var APP Config

// Load config file from given path
func Load(path string, fileName string, configType string) error {
	v := viper.New()

	v.AddConfigPath(path)
	v.SetConfigName(fileName)
	v.SetConfigType(configType)

	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return fmt.Errorf("config file not found")
		}
	}
	// Unmarshal the config into the struct
	if err := v.Unmarshal(&APP); err != nil {
		return fmt.Errorf("error unmarshaling config: %s", err)
	}

	return nil
}
