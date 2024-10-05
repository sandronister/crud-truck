package di

import (
	"github.com/sandronister/crud-truck/config"
	"github.com/sandronister/crud-truck/internal/infra/database/connection"
	"github.com/sandronister/crud-truck/internal/infra/database/repository"
	"github.com/sandronister/crud-truck/internal/infra/handler"
	"github.com/sandronister/crud-truck/internal/usecase"
)

func NewLinkHandler(connection connection.IConnection, conf *config.Conf) (*handler.LinkHandler, error) {

	linkRepo, err := repository.NewLinkRepository(connection, conf)
	if err != nil {
		return nil, err
	}

	driveRepo, err := repository.NewDriverRepository(connection, conf)
	if err != nil {
		return nil, err
	}

	truckRepo, err := repository.NewTruckRepository(connection, conf)
	if err != nil {
		return nil, err
	}

	usecase := usecase.NewLinkUseCase(linkRepo, driveRepo, truckRepo)
	return handler.NewLinkHandler(usecase), nil
}
