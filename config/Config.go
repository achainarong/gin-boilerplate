package config

import (
	"fmt"
	"os"
)

type Config struct {
	Hostname string
	Port string
	DatabaseName string
	ConnectionString string
}

var (
	configInstance Config
)

func GetConfig() *Config {

	if configInstance == (Config{}) {
		configInstance.Hostname = os.Getenv("HOSTNAME")
		configInstance.Port = os.Getenv("PORT")
		configInstance.DatabaseName = os.Getenv("DB_NAME")
		configInstance.ConnectionString = fmt.Sprintf(
			"mongodb://%s:%s@%s:%s", 
				os.Getenv("DB_USER"),
				os.Getenv("DB_PASSWORD"),
				os.Getenv("DB_HOSTNAME"),
				os.Getenv("DB_PORT"))	
	}
	

	return &configInstance
}