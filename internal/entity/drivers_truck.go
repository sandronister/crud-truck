package entity

type DriverTrucks struct {
	DriverID   string
	DriverName string
	LicenseID  string
	Trucks     []Truck
}

func New(driverId, driverName, licenseID string, trucks []Truck) *DriverTrucks {
	return &DriverTrucks{
		DriverID:   driverId,
		DriverName: driverName,
		LicenseID:  licenseID,
		Trucks:     trucks,
	}
}
