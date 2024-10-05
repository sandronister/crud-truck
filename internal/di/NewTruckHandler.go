package di

import (
	"github.com/sandronister/crud-truck/config"
	"github.com/sandronister/crud-truck/internal/infra/database/connection"
	"github.com/sandronister/crud-truck/internal/infra/database/repository"
	"github.com/sandronister/crud-truck/internal/infra/handler"
	"github.com/sandronister/crud-truck/internal/usecase"
)

func NewTruckHandler(connection connection.IConnection, conf *config.Conf) (*handler.TruckHandler, error) {
	repository, err := repository.NewTruckRepository(connection, conf)
	if err != nil {
		return nil, err
	}
	usecase := usecase.NewTruckUseCase(repository)
	return handler.NewTruckHandler(usecase), nil
}
