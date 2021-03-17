package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/quyenphamkhac/go-tinyurl/config"
	"github.com/quyenphamkhac/go-tinyurl/server"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("[Error] loading .env file failed")
	}
	config.Init(os.Getenv("ENV"))
	server.Serve()
}
