package usecase

import (
	"fmt"

	"github.com/sandronister/crud-truck/internal/dto"
	"github.com/sandronister/crud-truck/internal/entity"
	"github.com/sandronister/crud-truck/internal/repository"
)

type TruckUseCase struct {
	repository repository.ITruck
}

func NewTruckUseCase(repository repository.ITruck) *TruckUseCase {
	return &TruckUseCase{
		repository: repository,
	}
}

func (u *TruckUseCase) Delete(id string) error {
	_, err := u.repository.FindById(id)

	if err != nil {
		return entity.ErrTruckNotFound
	}

	return u.repository.Delete(id)
}

func (u *TruckUseCase) FindAll() ([]*dto.Truck, error) {
	trucks, err := u.repository.FindAll()

	if err != nil {
		return nil, err
	}

	if len(*trucks) == 0 {
		return nil, entity.ErrNotFoundTrucks
	}

	var trucksDTO []*dto.Truck

	for _, truck := range *trucks {
		trucksDTO = append(trucksDTO, &dto.Truck{
			ID:           truck.ID,
			Brand:        truck.Brand,
			Model:        truck.Model,
			Year:         truck.Year,
			LicensePlate: truck.LicensePlate,
		})
	}

	return trucksDTO, nil
}

func (u *TruckUseCase) FindByID(id string) (*dto.Truck, error) {

	truck, err := u.repository.FindById(id)
	if err != nil {
		return nil, err
	}
	truckDTO := &dto.Truck{
		ID:           truck.ID,
		Brand:        truck.Brand,
		Model:        truck.Model,
		Year:         truck.Year,
		LicensePlate: truck.LicensePlate,
	}
	return truckDTO, nil
}

func (u *TruckUseCase) Save(dto *dto.Truck) error {

	truck := entity.NewTruck(dto.Brand, dto.Model, dto.LicensePlate, dto.Year)

	if err := truck.Validate(); err != nil {
		return fmt.Errorf("error validating truck: %w", err)
	}

	if u.repository.ExistsLicensePlate(truck.LicensePlate) {
		return fmt.Errorf("license plate already exists")
	}

	err := u.repository.Save(truck)
	if err != nil {
		return err
	}

	dto.ID = truck.ID

	return nil

}

func (u *TruckUseCase) Update(dto *dto.Truck) error {
	truck := entity.NewTruck(dto.Brand, dto.Model, dto.LicensePlate, dto.Year)
	truck.ID = dto.ID

	oldTruck, err := u.repository.FindById(truck.ID)

	if err != nil {
		return entity.ErrTruckNotFound
	}

	if truck.Brand == "" {
		truck.Brand = oldTruck.Brand
	}

	if truck.Model == "" {
		truck.Model = oldTruck.Model
	}

	if truck.LicensePlate == "" {
		truck.LicensePlate = oldTruck.LicensePlate
	}

	if truck.Year == 0 {
		truck.Year = oldTruck.Year
	}

	return u.repository.Update(truck)

}
