package config

import (
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
	viper.SetConfigName(env)
	viper.SetConfigType("yaml")
	viper.AddConfigPath("config/")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("[Error] Parsing config file failed")
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
