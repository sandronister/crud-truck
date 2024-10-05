package entity

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

var (
	ErrTruckIDIsRequired  = errors.New("truck id is required")
	ErrDriverIDIsRequired = errors.New("driver id is required")
	ErrLinkAlreadyExists  = errors.New("link already exists")
	ErrLinkNotFound       = errors.New("link not found")
	ErrNotFoundTrucks     = errors.New("trucks not found")
)

type Link struct {
	ID        string
	TruckID   string
	DriverID  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewLink(truckID, driverID string) *Link {
	return &Link{
		ID:       uuid.New().String(),
		TruckID:  truckID,
		DriverID: driverID,
	}
}

func (l *Link) Validate() error {
	if l.TruckID == "" {
		return ErrTruckIDIsRequired
	}

	if l.DriverID == "" {
		return ErrDriverIDIsRequired
	}

	return nil
}
