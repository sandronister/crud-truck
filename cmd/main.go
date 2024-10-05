package main

import (
	"github.com/sandronister/crud-truck/config"
	"github.com/sandronister/crud-truck/internal/di"
	"github.com/sandronister/crud-truck/internal/infra/database/connection"
	"github.com/sandronister/crud-truck/internal/infra/web"
	"github.com/sandronister/crud-truck/pkg/exception"
)

func main() {
	conf, err := config.LoadConfig(".")
	exception.Catch(err)

	db, err := connection.GetConnection(conf)
	exception.Catch(err)

	defer db.Close()

	err = db.Ping()
	exception.Catch(err)

	driverHandler, err := di.NewDriverHandler(db, conf)
	exception.Catch(err)

	truckHandler, err := di.NewTruckHandler(db, conf)
	exception.Catch(err)

	linkHandler, err := di.NewLinkHandler(db, conf)
	exception.Catch(err)

	server := web.NewServer(conf.WebPort)
	server.AddDriverHandler(driverHandler)
	server.AddTruckHandler(truckHandler)
	server.AddLinkHandler(linkHandler)

	server.Run()

}
