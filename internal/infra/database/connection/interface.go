package connection

import (
	"database/sql"

	"github.com/gocql/gocql"
)

type IConnection interface {
	Close()
	Ping() error
	GetSQLConnection() (*sql.DB, error)
	GetCassandraConnection() (*gocql.Session, error)
}
