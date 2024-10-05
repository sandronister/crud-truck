package usecase

import (
	"database/sql"
	"errors"
	"testing"

	"github.com/sandronister/crud-truck/internal/dto"
	"github.com/sandronister/crud-truck/internal/entity"
	repomock "github.com/sandronister/crud-truck/internal/usecase/repo_mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestDeleteDriverSuccess(t *testing.T) {
	mockRepo := new(repomock.DriveRepositoryMock)
	mockRepo.On("FindById", "123").Return(&entity.Driver{}, nil)
	mockRepo.On("Delete", "123").Return(nil)

	ud := NewDriverUseCase(mockRepo)
	err := ud.Delete("123")
	assert.Nil(t, err)
}

func TestDeleteDriverFail(t *testing.T) {
	mockRepo := new(repomock.DriveRepositoryMock)
	mockRepo.On("FindById", "123").Return(&entity.Driver{}, errors.New("failed"))
	mockRepo.On("Delete", "123").Return(nil)

	ud := NewDriverUseCase(mockRepo)
	err := ud.Delete("123")
	assert.NotEmpty(t, err)
}

func TestFindAllDriverSuccess(t *testing.T) {
	mockRepo := new(repomock.DriveRepositoryMock)
	mockRepo.On("FindAll").Return(&[]entity.Driver{}, nil)

	ud := NewDriverUseCase(mockRepo)
	result, _ := ud.FindAll()
	assert.Empty(t, result)
}

func TestFindAllDriverFail(t *testing.T) {
	mockRepo := new(repomock.DriveRepositoryMock)
	mockRepo.On("FindAll").Return(&[]entity.Driver{}, entity.ErrDriversNotFound)

	ud := NewDriverUseCase(mockRepo)
	result, err := ud.FindAll()
	assert.NotEmpty(t, err)
	assert.Empty(t, result)
}

func TestFindAllDriverNotFound(t *testing.T) {
	mockRepo := new(repomock.DriveRepositoryMock)
	mockRepo.On("FindAll").Return(&[]entity.Driver{}, sql.ErrNoRows)

	ud := NewDriverUseCase(mockRepo)
	_, err := ud.FindAll()
	assert.NotEmpty(t, err)
}

func TestFindIdDriverSuccess(t *testing.T) {
	mockRepo := new(repomock.DriveRepositoryMock)
	mockRepo.On("FindById", "123").Return(&entity.Driver{}, nil)

	ud := NewDriverUseCase(mockRepo)
	result, err := ud.FindByID("123")
	assert.Nil(t, err)
	assert.NotEmpty(t, result)
}

func TestFindIdDriverFail(t *testing.T) {
	mockRepo := new(repomock.DriveRepositoryMock)
	mockRepo.On("FindById", "123").Return(&entity.Driver{}, errors.New("failed"))

	ud := NewDriverUseCase(mockRepo)
	result, err := ud.FindByID("123")
	assert.NotEmpty(t, err)
	assert.Nil(t, result)
}

func TestSaveDriverSuccess(t *testing.T) {
	mockRepo := new(repomock.DriveRepositoryMock)
	usecase := DriverUsecase{repository: mockRepo}

	mockRepo.On("ExistsLicenseID", "LICENSE123").Return(false)
	mockRepo.On("Save", mock.AnythingOfType("*entity.Driver")).Return(nil)

	result, err := usecase.Save("John Doe", "LICENSE123")
	assert.NoError(t, err)
	assert.NotEmpty(t, result)

}

func TestSaveDriverFail(t *testing.T) {
	mockRepo := new(repomock.DriveRepositoryMock)
	usecase := DriverUsecase{repository: mockRepo}

	mockRepo.On("ExistsLicenseID", "LICENSE123").Return(true)

	result, err := usecase.Save("John Doe", "LICENSE123")
	assert.Error(t, err)
	assert.Empty(t, result)
}

func TestUpdateDriverSuccess(t *testing.T) {
	mockRepo := new(repomock.DriveRepositoryMock)
	mockRepo.On("FindById", "123").Return(&entity.Driver{}, nil)
	mockRepo.On("Update", mock.AnythingOfType("*entity.Driver")).Return(nil)

	ud := NewDriverUseCase(mockRepo)
	err := ud.Update(&dto.Driver{
		ID:        "123",
		Name:      "John Doe",
		LicenseID: "LICENSE123",
	})
	assert.Nil(t, err)
}

func TestUpdateDriverFail(t *testing.T) {
	mockRepo := new(repomock.DriveRepositoryMock)
	mockRepo.On("FindById", "123").Return(&entity.Driver{}, errors.New("error"))
	mockRepo.On("Update", mock.AnythingOfType("*entity.Driver")).Return(nil)

	ud := NewDriverUseCase(mockRepo)
	err := ud.Update(&dto.Driver{
		ID:        "123",
		Name:      "John Doe",
		LicenseID: "LICENSE123",
	})
	assert.NotEmpty(t, err)
}
