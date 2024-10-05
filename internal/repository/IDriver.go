package repository

import "github.com/sandronister/crud-truck/internal/entity"

type IDriver interface {
	Save(driver *entity.Driver) error
	FindAll() (*[]entity.Driver, error)
	FindById(id string) (*entity.Driver, error)
	ExistsLicenseID(license string) bool
	Update(driver *entity.Driver) error
	Delete(id string) error
}
