package connection

import (
	"database/sql"
	"fmt"

	"github.com/gocql/gocql"
	"github.com/sandronister/crud-truck/config"
)

type MysqlConnection struct {
	*sql.DB
}

func getMysqlString(conf *config.Conf) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", conf.DBUser, conf.DBPassword, conf.DBHost, conf.DBPort, conf.DBName)
}

func NewMysqlConnection(conf *config.Conf) (*MysqlConnection, error) {
	db, err := sql.Open("mysql", getMysqlString(conf))
	if err != nil {
		return nil, err
	}

	return &MysqlConnection{DB: db}, nil
}

func (m *MysqlConnection) Close() {
	m.DB.Close()
}

func (m *MysqlConnection) Ping() error {
	return m.DB.Ping()
}

func (m *MysqlConnection) GetCassandraConnection() (*gocql.Session, error) {
	return nil, fmt.Errorf("not cassandra connection")
}

func (m *MysqlConnection) GetSQLConnection() (*sql.DB, error) {
	return m.DB, nil
}
