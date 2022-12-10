package main

import (
	"gin/server"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	server.Init()
}

