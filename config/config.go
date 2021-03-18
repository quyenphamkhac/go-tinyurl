package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

type DbConfig struct {
	Host         string
	Port         string
	DatabaseName string
}

type ServerConfig struct {
	Host string
	Port string
}

type Configuration struct {
	DbConfig
	ServerConfig
}

var config *Configuration

func Init(env string) {
	fileConfig := fmt.Sprintf("config/%s.yaml", env)
	fmt.Println(fileConfig)
	viper.SetConfigFile(fileConfig)
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("[Error] Parsing config file failed: %s", err.Error())
	}
	config = &Configuration{
		DbConfig: DbConfig{
			Host:         viper.GetString("cassandra.host"),
			Port:         viper.GetString("cassandra.port"),
			DatabaseName: viper.GetString("canssandra.database_name"),
		},
		ServerConfig: ServerConfig{
			Host: viper.GetString("server.host"),
			Port: viper.GetString("server.port"),
		},
	}
}

func GetConfig() *Configuration {
	return config
}
