package usecase

import (
	"github.com/sandronister/crud-truck/internal/dto"
	"github.com/sandronister/crud-truck/internal/entity"
	"github.com/sandronister/crud-truck/internal/repository"
)

type LinkUseCase struct {
	linkRepository   repository.ILink
	driverRepository repository.IDriver
	truckRepository  repository.ITruck
}

func NewLinkUseCase(
	linkRepository repository.ILink,
	driverRepository repository.IDriver,
	truckRepository repository.ITruck,
) *LinkUseCase {
	return &LinkUseCase{
		linkRepository:   linkRepository,
		driverRepository: driverRepository,
		truckRepository:  truckRepository,
	}
}

func (s *LinkUseCase) Delete(driverID, truckID string) error {
	if !s.linkRepository.ExistsLink(driverID, truckID) {
		return entity.ErrLinkNotFound
	}

	return s.linkRepository.Delete(driverID, truckID)
}

func (u *LinkUseCase) ListByDriver(id string) (*dto.DriverTrucks, error) {
	_, err := u.driverRepository.FindById(id)

	if err != nil {
		return nil, entity.ErrDriverNotFound
	}

	trucksDriver, err := u.linkRepository.GetTrucksByDriver(id)

	if err != nil {
		return nil, entity.ErrNotFoundTrucks
	}

	var trucks []dto.Truck

	for _, truck := range trucksDriver.Trucks {
		trucks = append(trucks, dto.Truck{
			ID:           truck.ID,
			Brand:        truck.Brand,
			LicensePlate: truck.LicensePlate,
			Year:         truck.Year,
			CreatedAt:    truck.CreatedAt.String(),
			UpdatedAt:    truck.UpdatedAt.String(),
		})
	}

	return &dto.DriverTrucks{
		DriverID:   id,
		DriverName: trucksDriver.DriverName,
		LicenseID:  trucksDriver.LicenseID,
		Trucks:     trucks,
	}, nil

}

func (s *LinkUseCase) Save(dto *dto.Link) error {
	_, err := s.driverRepository.FindById(dto.DriverID)

	if err != nil {
		return entity.ErrDriverNotFound
	}

	_, err = s.truckRepository.FindById(dto.TruckID)

	if err != nil {
		return entity.ErrTruckNotFound
	}

	link := entity.NewLink(dto.TruckID, dto.DriverID)

	if err := link.Validate(); err != nil {
		return err
	}

	if s.linkRepository.ExistsLink(link.DriverID, link.TruckID) {
		return entity.ErrLinkAlreadyExists
	}

	dto.ID = link.ID

	return s.linkRepository.Save(link)
}
