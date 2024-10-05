package repomock

import (
	"github.com/sandronister/crud-truck/internal/entity"
	"github.com/stretchr/testify/mock"
)

type LinkRepoMock struct {
	mock.Mock
}

func (m *LinkRepoMock) Delete(driverID, truckID string) error {
	args := m.Called(driverID, truckID)
	return args.Error(0)
}

func (m *LinkRepoMock) ExistsLink(driverID, truckID string) bool {
	args := m.Called(driverID, truckID)
	return args.Bool(0)
}

func (m *LinkRepoMock) GetTrucksByDriver(driverID string) (*entity.DriverTrucks, error) {
	args := m.Called(driverID)
	return args.Get(0).(*entity.DriverTrucks), args.Error(1)
}

func (m *LinkRepoMock) Save(link *entity.Link) error {
	args := m.Called(link)
	return args.Error(0)
}
