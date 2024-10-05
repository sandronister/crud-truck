package repomock

import (
	"github.com/sandronister/crud-truck/internal/entity"
	"github.com/stretchr/testify/mock"
)

type TruckRepoMock struct {
	mock.Mock
}

func (m *TruckRepoMock) FindById(id string) (*entity.Truck, error) {
	args := m.Called(id)
	return args.Get(0).(*entity.Truck), args.Error(1)
}

func (m *TruckRepoMock) Delete(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *TruckRepoMock) ExistsLicenseID(id string) bool {
	args := m.Called(id)
	return args.Bool(0)
}

func (m *TruckRepoMock) FindAll() (*[]entity.Truck, error) {
	args := m.Called()
	return args.Get(0).(*[]entity.Truck), args.Error(1)
}

func (m *TruckRepoMock) Save(driver *entity.Truck) error {
	args := m.Called(driver)
	return args.Error(0)
}

func (m *TruckRepoMock) Update(driver *entity.Truck) error {
	args := m.Called(driver)
	return args.Error(0)
}

func (m *TruckRepoMock) ExistsLicensePlate(licensePlate string) bool {
	args := m.Called(licensePlate)
	return args.Bool(0)
}
