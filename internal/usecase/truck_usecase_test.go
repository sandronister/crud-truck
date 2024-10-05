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

func TestDeleteTruckSuccess(t *testing.T) {
	mockRepo := new(repomock.TruckRepoMock)
	mockRepo.On("FindById", "123").Return(&entity.Truck{}, nil)
	mockRepo.On("Delete", "123").Return(nil)

	ut := NewTruckUseCase(mockRepo)
	err := ut.Delete("123")
	assert.Nil(t, err)
}

func TestDeleteTruckFail(t *testing.T) {
	mockRepo := new(repomock.TruckRepoMock)
	mockRepo.On("FindById", "123").Return(&entity.Truck{}, errors.New("failed"))
	mockRepo.On("Delete", "123").Return(nil)

	ut := NewTruckUseCase(mockRepo)
	err := ut.Delete("123")
	assert.NotEmpty(t, err)
}

func TestFindAllTruckSuccess(t *testing.T) {
	mockRepo := new(repomock.TruckRepoMock)
	mockRepo.On("FindAll").Return(&[]entity.Truck{}, nil)

	ud := NewTruckUseCase(mockRepo)
	result, _ := ud.FindAll()
	assert.Empty(t, result)
}

func TestFindAllTruckFail(t *testing.T) {
	mockRepo := new(repomock.TruckRepoMock)
	mockRepo.On("FindAll").Return(&[]entity.Truck{}, entity.ErrTruckNotFound)

	ud := NewTruckUseCase(mockRepo)
	result, err := ud.FindAll()
	assert.NotEmpty(t, err)
	assert.Empty(t, result)
}

func TestFindAllTruckNotFound(t *testing.T) {
	mockRepo := new(repomock.TruckRepoMock)
	mockRepo.On("FindAll").Return(&[]entity.Truck{}, nil)

	ud := NewTruckUseCase(mockRepo)
	result, err := ud.FindAll()
	assert.NotEmpty(t, err)
	assert.Empty(t, result)
}

func TestFindIdSuccess(t *testing.T) {
	mockRepo := new(repomock.TruckRepoMock)
	mockRepo.On("FindById", "123").Return(&entity.Truck{
		ID:           "123",
		Brand:        "Scania",
		Model:        "R440",
		Year:         2015,
		LicensePlate: "ABC1234",
	}, nil)

	ut := NewTruckUseCase(mockRepo)
	result, err := ut.FindByID("123")
	assert.Nil(t, err)
	assert.NotEmpty(t, result)
}

func TestFindIdTruckFail(t *testing.T) {
	mockRepo := new(repomock.TruckRepoMock)
	mockRepo.On("FindById", "123").Return(&entity.Truck{}, errors.New("failed"))

	ut := NewTruckUseCase(mockRepo)
	result, err := ut.FindByID("123")
	assert.NotEmpty(t, err)
	assert.Nil(t, result)
}

func TestSaveTruckSuccess(t *testing.T) {
	mockRepo := new(repomock.TruckRepoMock)
	ut := NewTruckUseCase(mockRepo)

	mockRepo.On("ExistsLicensePlate", "LICENSE123").Return(false)
	mockRepo.On("Save", mock.AnythingOfType("*entity.Truck")).Return(nil)

	err := ut.Save(&dto.Truck{
		Brand:        "Volvo",
		Model:        "FH16",
		Year:         2020,
		LicensePlate: "LICENSE123",
	})
	assert.NoError(t, err)

}

func TestSaveTruckFail(t *testing.T) {
	mockRepo := new(repomock.TruckRepoMock)
	ut := NewTruckUseCase(mockRepo)

	mockRepo.On("ExistsLicenseID", "LICENSE123").Return(true)

	err := ut.Save(&dto.Truck{})
	assert.Error(t, err)
}

func TestUpdateSuccess(t *testing.T) {
	mockRepo := new(repomock.TruckRepoMock)
	mockRepo.On("FindById", "123").Return(&entity.Truck{}, nil)
	mockRepo.On("Update", mock.AnythingOfType("*entity.Truck")).Return(nil)

	ut := NewTruckUseCase(mockRepo)
	err := ut.Update(&dto.Truck{
		ID:    "123",
		Brand: "Volvo",
		Model: "FH16",
		Year:  2020,
	})
	assert.Nil(t, err)
}

func TestUpdateFail(t *testing.T) {
	mockRepo := new(repomock.TruckRepoMock)
	mockRepo.On("FindById", "123").Return(&entity.Truck{}, errors.New("error"))
	mockRepo.On("Update", mock.AnythingOfType("*entity.Truck")).Return(nil)

	ut := NewTruckUseCase(mockRepo)
	err := ut.Update(&dto.Truck{
		ID:    "123",
		Brand: "Volvo",
		Model: "FH16",
		Year:  2020,
	})
	assert.NotEmpty(t, err)
}
