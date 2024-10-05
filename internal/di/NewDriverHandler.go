package di

import (
	"github.com/sandronister/crud-truck/config"
	"github.com/sandronister/crud-truck/internal/infra/database/connection"
	"github.com/sandronister/crud-truck/internal/infra/database/repository"
	"github.com/sandronister/crud-truck/internal/infra/handler"
	"github.com/sandronister/crud-truck/internal/usecase"
)

func NewDriverHandler(connection connection.IConnection, conf *config.Conf) (*handler.DriverHandler, error) {
	repository, err := repository.NewDriverRepository(connection, conf)
	if err != nil {
		return nil, err
	}
	usecase := usecase.NewDriverUseCase(repository)
	return handler.NewDriverHandler(usecase), nil
}
