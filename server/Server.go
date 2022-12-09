package server

import (
	"fmt"
	"gin/config"

	"github.com/joho/godotenv"
)

func Init(){
	godotenv.Load()
	
	router := InitRouter()

	newConfig := config.GetConfig()

	url := fmt.Sprintf("%s:%s", newConfig.Hostname, newConfig.Port)

	router.Run(url)
}