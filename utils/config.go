package utils

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	Port string `mapstructure:"port"`
}

func LoadConfig() (config Config, err error) {
	viper.SetConfigName("default")   // name of config file (without extension)
	viper.SetConfigType("json")      // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("./config/") // path to look for the config file in

	err = viper.ReadInConfig() // Find and read the config file
	if err != nil {            // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
	err = viper.Unmarshal(&config)
	return
}
