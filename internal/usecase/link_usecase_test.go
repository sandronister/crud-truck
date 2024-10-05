package usecase

import (
	"errors"
	"testing"

	"github.com/sandronister/crud-truck/internal/dto"
	"github.com/sandronister/crud-truck/internal/entity"
	repomock "github.com/sandronister/crud-truck/internal/usecase/repo_mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func GetRepoMock() (*repomock.LinkRepoMock, *repomock.TruckRepoMock, *repomock.DriveRepositoryMock) {
	linkRepo := new(repomock.LinkRepoMock)
	truckRepo := new(repomock.TruckRepoMock)
	driveRepo := new(repomock.DriveRepositoryMock)

	return linkRepo, truckRepo, driveRepo
}

func TestDeleteLinkSuccess(t *testing.T) {
	linkRepo, truckRepo, driveRepo := GetRepoMock()

	linkRepo.On("Delete", "123", "456").Return(nil)
	linkRepo.On("ExistsLink", "123", "456").Return(true)

	ud := NewLinkUseCase(linkRepo, driveRepo, truckRepo)
	err := ud.Delete("123", "456")
	assert.Nil(t, err)
}

func TestDeleteLinkFail(t *testing.T) {
	linkRepo, truckRepo, driveRepo := GetRepoMock()

	linkRepo.On("Delete", "123", "456").Return(errors.New("error"))
	linkRepo.On("ExistsLink", "123", "456").Return(false)

	ud := NewLinkUseCase(linkRepo, driveRepo, truckRepo)
	err := ud.Delete("123", "456")
	assert.NotEmpty(t, err)
}

func TestLisByDriverSuccess(t *testing.T) {
	linkRepo, truckRepo, driveRepo := GetRepoMock()

	driveRepo.On("FindById", "123").Return(&entity.Driver{}, nil)
	linkRepo.On("GetTrucksByDriver", "123").Return(&entity.DriverTrucks{}, nil)

	ud := NewLinkUseCase(linkRepo, driveRepo, truckRepo)
	links, err := ud.ListByDriver("123")
	assert.Nil(t, err)
	assert.NotEmpty(t, links)
}

func TestLisByDriverFail(t *testing.T) {
	linkRepo, truckRepo, driveRepo := GetRepoMock()

	driveRepo.On("FindById", "123").Return(&entity.Driver{}, errors.New("error"))
	linkRepo.On("GetTrucksByDriver", "123").Return(&entity.DriverTrucks{}, nil)

	ud := NewLinkUseCase(linkRepo, driveRepo, truckRepo)
	_, err := ud.ListByDriver("123")
	assert.Error(t, err)
}

func TestSaveLinkSuccess(t *testing.T) {
	linkRepo, truckRepo, driveRepo := GetRepoMock()

	driveRepo.On("FindById", "123").Return(&entity.Driver{}, nil)
	linkRepo.On("Save", mock.AnythingOfType("*entity.Link")).Return(nil)
	linkRepo.On("ExistsLink", "123", "123").Return(false)
	truckRepo.On("FindById", "123").Return(&entity.Truck{}, nil)

	ud := NewLinkUseCase(linkRepo, driveRepo, truckRepo)
	err := ud.Save(&dto.Link{
		DriverID: "123",
		TruckID:  "123",
	})
	assert.Nil(t, err)
}

func TestSaveDriverNotExists(t *testing.T) {
	linkRepo, truckRepo, driveRepo := GetRepoMock()

	driveRepo.On("FindById", "123").Return(&entity.Driver{}, errors.New("not found"))
	linkRepo.On("Save", mock.AnythingOfType("*entity.Link")).Return(nil)
	linkRepo.On("ExistsLink", "123", "123").Return(false)
	truckRepo.On("FindById", "123").Return(&entity.Truck{}, nil)

	ud := NewLinkUseCase(linkRepo, driveRepo, truckRepo)
	err := ud.Save(&dto.Link{
		DriverID: "123",
		TruckID:  "123",
	})
	assert.Error(t, err)
}

func TestSaveTruckNotExists(t *testing.T) {
	linkRepo, truckRepo, driveRepo := GetRepoMock()

	driveRepo.On("FindById", "123").Return(&entity.Driver{}, nil)
	linkRepo.On("Save", mock.AnythingOfType("*entity.Link")).Return(nil)
	linkRepo.On("ExistsLink", "123", "123").Return(false)
	truckRepo.On("FindById", "123").Return(&entity.Truck{}, errors.New("not found"))

	ud := NewLinkUseCase(linkRepo, driveRepo, truckRepo)
	err := ud.Save(&dto.Link{
		DriverID: "123",
		TruckID:  "123",
	})
	assert.Error(t, err)
}

func TestSaveLinkExists(t *testing.T) {
	linkRepo, truckRepo, driveRepo := GetRepoMock()

	driveRepo.On("FindById", "123").Return(&entity.Driver{}, nil)
	linkRepo.On("Save", mock.AnythingOfType("*entity.Link")).Return(nil)
	linkRepo.On("ExistsLink", "123", "123").Return(true)
	truckRepo.On("FindById", "123").Return(&entity.Truck{}, nil)

	ud := NewLinkUseCase(linkRepo, driveRepo, truckRepo)
	err := ud.Save(&dto.Link{
		DriverID: "123",
		TruckID:  "123",
	})
	assert.Error(t, err)
}
