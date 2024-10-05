package repository

import "github.com/sandronister/crud-truck/internal/entity"

type ITruck interface {
	Save(truck *entity.Truck) error
	FindAll() (*[]entity.Truck, error)
	FindById(id string) (*entity.Truck, error)
	ExistsLicensePlate(licensePlate string) bool
	Update(truck *entity.Truck) error
	Delete(id string) error
}
