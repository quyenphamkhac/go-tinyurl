package db

import (
	"log"

	"github.com/gocql/gocql"
	"github.com/quyenphamkhac/go-tinyurl/config"
)

var db *gocql.Session

func InitDatabase() {
	cluster := gocql.NewCluster(config.GetConfig().DbConfig.Host)
	cluster.Keyspace = config.GetConfig().KeySpace
	session, err := cluster.CreateSession()
	if err != nil {
		log.Fatal("[Error] Create session failed")
	}
	db = session
	defer db.Close()
}

func GetDb() *gocql.Session {
	return db
}
