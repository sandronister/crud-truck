package repomock

import (
	"github.com/sandronister/crud-truck/internal/entity"
	"github.com/stretchr/testify/mock"
)

type DriveRepositoryMock struct {
	mock.Mock
}

func (m *DriveRepositoryMock) FindById(id string) (*entity.Driver, error) {
	args := m.Called(id)
	return args.Get(0).(*entity.Driver), args.Error(1)
}

func (m *DriveRepositoryMock) Delete(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *DriveRepositoryMock) ExistsLicenseID(id string) bool {
	args := m.Called(id)
	return args.Bool(0)
}

func (m *DriveRepositoryMock) FindAll() (*[]entity.Driver, error) {
	args := m.Called()
	return args.Get(0).(*[]entity.Driver), args.Error(1)
}

func (m *DriveRepositoryMock) Save(driver *entity.Driver) error {
	args := m.Called(driver)
	return args.Error(0)
}

func (m *DriveRepositoryMock) Update(driver *entity.Driver) error {
	args := m.Called(driver)
	return args.Error(0)
}
