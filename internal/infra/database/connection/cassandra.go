package connection

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/gocql/gocql"
	"github.com/sandronister/crud-truck/config"
)

type CassandraConnection struct {
	Session *gocql.Session
}

func NewCassandraConnection(conf *config.Conf) (*CassandraConnection, error) {
	cluster := gocql.NewCluster(conf.Ips...)
	cluster.Consistency = gocql.Quorum
	cluster.ProtoVersion = 4
	cluster.ConnectTimeout = time.Second * 5
	session, err := cluster.CreateSession()

	if err != nil {
		return nil, err
	}

	return &CassandraConnection{Session: session}, nil
}

func (c *CassandraConnection) Close() {
	c.Session.Close()
}

func (c *CassandraConnection) Ping() error {
	return c.Session.Query("SELECT now() FROM system.local").Exec()
}

func (c *CassandraConnection) GetSQLConnection() (*sql.DB, error) {
	return nil, fmt.Errorf("not sql connection")
}

func (c *CassandraConnection) GetCassandraConnection() (*gocql.Session, error) {
	return c.Session, nil
}
