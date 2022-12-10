package server

import (
	"fmt"
	"gin/config"
)

func Init(){
	router := InitRouter()

	newConfig := config.GetConfig()

	url := fmt.Sprintf("%s:%s", newConfig.Hostname, newConfig.Port)

	router.Run(url)
}