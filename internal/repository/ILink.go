package repository

import "github.com/sandronister/crud-truck/internal/entity"

type ILink interface {
	Save(link *entity.Link) error
	ExistsLink(driverID, truckID string) bool
	Delete(driverID, truckID string) error
	GetTrucksByDriver(driverID string) (*entity.DriverTrucks, error)
}
