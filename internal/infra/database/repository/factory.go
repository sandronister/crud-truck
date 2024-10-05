package repository

import (
	"fmt"

	"github.com/sandronister/crud-truck/config"
	"github.com/sandronister/crud-truck/internal/infra/database/connection"
	"github.com/sandronister/crud-truck/internal/infra/database/repository/mysql"
	"github.com/sandronister/crud-truck/internal/repository"
)

func NewDriverRepository(db connection.IConnection, conf *config.Conf) (repository.IDriver, error) {
	switch conf.DBDriver {
	case "mysql":
		db, err := db.GetSQLConnection()
		if err != nil {
			return nil, err
		}
		return mysql.NewDriverRepository(db), nil
	default:
		return nil, fmt.Errorf("invalid database driver")
	}
}

func NewTruckRepository(db connection.IConnection, conf *config.Conf) (repository.ITruck, error) {
	switch conf.DBDriver {
	case "mysql":
		db, err := db.GetSQLConnection()
		if err != nil {
			return nil, err
		}
		return mysql.NewTruckRepository(db), nil
	default:
		return nil, fmt.Errorf("invalid database driver")
	}
}

func NewLinkRepository(db connection.IConnection, conf *config.Conf) (repository.ILink, error) {
	switch conf.DBDriver {
	case "mysql":
		db, err := db.GetSQLConnection()
		if err != nil {
			return nil, err
		}
		return mysql.NewLinkRepository(db), nil
	default:
		return nil, fmt.Errorf("invalid database driver")
	}
}
