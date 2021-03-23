package datasources

import (
	"log"

	"github.com/gocql/gocql"
	"github.com/quyenphamkhac/go-tinyurl/config"
)

var session *gocql.Session

func InitDatabase() {
	cluster := gocql.NewCluster(config.GetConfig().DbConfig.Host)
	cluster.Keyspace = config.GetConfig().KeySpace
	var err error
	session, err = cluster.CreateSession()
	if err != nil {
		log.Fatal("[Error] Create session failed")
	}
}

func GetSession() *gocql.Session {
	return session
}
