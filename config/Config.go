package config

import "os"

type Config struct {
	Hostname string
	Port string
}

func GetConfig() *Config {
	config := Config{}
	config.Hostname = os.Getenv("HOSTNAME")
	config.Port = os.Getenv("PORT")
	return &config
}