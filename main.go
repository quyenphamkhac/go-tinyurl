package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/quyenphamkhac/go-tinyurl/config"
	"github.com/quyenphamkhac/go-tinyurl/server"
)

func main() {
	if os.Getenv("ENV") == "" {
		if err := godotenv.Load(".env.dev"); err != nil {
			log.Fatalf("[Error] loading .env file failed")
		}
	}
	config.Init(os.Getenv("ENV"))
	server.Serve()
}
