package entity

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

var (
	ErrNameIsRequired      = errors.New("name is required")
	ErrLicenseIDIsRequired = errors.New("license id is required")
	ErrDriverAlreadyExists = errors.New("driver already exists")
	ErrDriverNotFound      = errors.New("driver not found")
	ErrDriversNotFound     = errors.New("drivers not found")
)

type Driver struct {
	ID        string
	Name      string
	LicenseID string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewDriver(name, licenseID string) *Driver {
	return &Driver{
		ID:        uuid.New().String(),
		Name:      name,
		LicenseID: licenseID,
	}
}

func (d *Driver) Validate() error {
	if d.Name == "" {
		return ErrNameIsRequired
	}
	if d.LicenseID == "" {
		return ErrLicenseIDIsRequired
	}
	return nil
}
