package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

type DbConfig struct {
	Host     string
	Port     string
	KeySpace string
}

type RedisConfig struct {
	Host string
	Port string
	Addr string
}

type ServerConfig struct {
	Host string
	Port string
}

type JwtConfig struct {
	Secret string
	Issuer string
}

type Configuration struct {
	DbConfig
	ServerConfig
	JwtConfig
	RedisConfig
}

var config *Configuration

func Init(env string) {
	fileConfig := fmt.Sprintf("config/%s.yaml", env)
	viper.SetConfigFile(fileConfig)
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("[Error] Parsing config file failed: %s", err.Error())
	}
	config = &Configuration{
		DbConfig: DbConfig{
			Host:     viper.GetString("cassandra.host"),
			Port:     viper.GetString("cassandra.port"),
			KeySpace: viper.GetString("cassandra.key_space"),
		},
		ServerConfig: ServerConfig{
			Host: viper.GetString("server.host"),
			Port: viper.GetString("server.port"),
		},
		JwtConfig: JwtConfig{
			Secret: viper.GetString("jwt.secret"),
			Issuer: viper.GetString("jwt.issuer"),
		},
		RedisConfig: RedisConfig{
			Host: viper.GetString("redis.host"),
			Port: viper.GetString("redis.port"),
			Addr: viper.GetString("redis.addr"),
		},
	}
}

func GetConfig() *Configuration {
	return config
}
