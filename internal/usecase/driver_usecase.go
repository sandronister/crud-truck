package usecase

import (
	"database/sql"

	"github.com/sandronister/crud-truck/internal/dto"
	"github.com/sandronister/crud-truck/internal/entity"
	"github.com/sandronister/crud-truck/internal/repository"
)

type DriverUsecase struct {
	repository repository.IDriver
}

func NewDriverUseCase(repository repository.IDriver) *DriverUsecase {
	return &DriverUsecase{
		repository: repository,
	}
}

func (d *DriverUsecase) Delete(id string) error {
	_, err := d.repository.FindById(id)

	if err != nil {
		return entity.ErrDriverNotFound
	}

	return d.repository.Delete(id)
}

func (d *DriverUsecase) FindAll() ([]*dto.Driver, error) {
	drivers, err := d.repository.FindAll()

	if err != nil {
		if err.Error() == sql.ErrNoRows.Error() {
			return nil, entity.ErrDriversNotFound
		}
		return nil, err
	}

	var driversDTO []*dto.Driver

	for _, driver := range *drivers {
		driversDTO = append(driversDTO, &dto.Driver{
			ID:        driver.ID,
			Name:      driver.Name,
			LicenseID: driver.LicenseID,
			CreatedAt: driver.CreatedAt.String(),
			UpdatedAt: driver.UpdatedAt.String(),
		})
	}
	return driversDTO, nil
}

func (d *DriverUsecase) FindByID(id string) (*dto.Driver, error) {
	driver, err := d.repository.FindById(id)

	if err != nil {
		return nil, entity.ErrDriverNotFound
	}

	return &dto.Driver{
		ID:        driver.ID,
		Name:      driver.Name,
		LicenseID: driver.LicenseID,
		CreatedAt: driver.CreatedAt.String(),
		UpdatedAt: driver.UpdatedAt.String(),
	}, nil
}

func (d *DriverUsecase) Save(name, licenseID string) (*dto.Driver, error) {
	driver := entity.NewDriver(name, licenseID)
	if err := driver.Validate(); err != nil {
		return nil, err
	}

	if exists := d.repository.ExistsLicenseID(driver.LicenseID); exists {
		return nil, entity.ErrDriverAlreadyExists
	}

	err := d.repository.Save(driver)

	if err != nil {
		return nil, err
	}

	driverDTO := &dto.Driver{
		ID:        driver.ID,
		Name:      driver.Name,
		LicenseID: driver.LicenseID,
	}

	return driverDTO, nil
}

func (d *DriverUsecase) Update(dto *dto.Driver) error {
	driver := entity.NewDriver(dto.Name, dto.LicenseID)
	driver.ID = dto.ID

	actualDriver, err := d.repository.FindById(driver.ID)

	if err != nil {
		if err.Error() == sql.ErrNoRows.Error() {
			return entity.ErrDriverNotFound
		}

		return err
	}

	if driver.Name == "" {
		driver.Name = actualDriver.Name
	}

	if driver.LicenseID == "" {
		driver.LicenseID = actualDriver.LicenseID
	}

	return d.repository.Update(driver)
}
