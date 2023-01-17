package config

import (
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	Server ServerConfig
}

type ServerConfig struct {
	IPAddr string
	Port   int
}

func LoadConfig() Config {
	var config Config
	viper.SetConfigName("config")
	viper.AddConfigPath("./config")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Config file read error: %s", err.Error())
	}

	if err := viper.Unmarshal(&config); err != nil {
		log.Fatalf("Unable to decode into struct: %v", err.Error())
	}

	return config
}
