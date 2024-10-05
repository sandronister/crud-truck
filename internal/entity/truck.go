package entity

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Truck struct {
	ID           string
	Brand        string
	Model        string
	Year         int
	LicensePlate string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

var (
	ErrModelIsRequired        = errors.New("model is required")
	ErrYearIsRequired         = errors.New("year is required")
	ErrBrandIsRequired        = errors.New("brand is required")
	ErrLicensePlateIsRequired = errors.New("license plate is required")
	ErrTruckNotFound          = errors.New("truck not found")
)

func NewTruck(brand, model, licensePlate string, year int) *Truck {
	return &Truck{
		ID:           uuid.New().String(),
		Brand:        brand,
		Model:        model,
		Year:         year,
		LicensePlate: licensePlate,
	}
}

func (t *Truck) Validate() error {
	if t.Brand == "" {
		return ErrBrandIsRequired
	}

	if t.Model == "" {
		return ErrModelIsRequired
	}

	if t.LicensePlate == "" {
		return ErrLicensePlateIsRequired
	}

	if t.Year == 0 {
		return ErrYearIsRequired
	}

	return nil
}
