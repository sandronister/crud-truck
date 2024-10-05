package connection

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/sandronister/crud-truck/config"
)

func GetConnection(conf *config.Conf) (IConnection, error) {

	switch conf.DBDriver {
	case "mysql":
		return NewMysqlConnection(conf)
	case "cassandra":
		return NewCassandraConnection(conf)
	}

	return nil, fmt.Errorf("invalid driver")
}
