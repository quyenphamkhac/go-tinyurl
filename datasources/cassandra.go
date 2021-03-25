package datasources

import (
	"sync"

	"github.com/gocql/gocql"
	"github.com/quyenphamkhac/go-tinyurl/config"
)

var (
	initCanssandraOnce sync.Once
	session            *gocql.Session
)

func GetCassandraSession() *gocql.Session {
	initCanssandraOnce.Do(func() {
		var err error
		cluster := gocql.NewCluster(config.GetConfig().DbConfig.Host)
		cluster.Keyspace = config.GetConfig().KeySpace
		session, err = cluster.CreateSession()
		if err != nil {
			panic(err)
		}
	})
	return session
}
