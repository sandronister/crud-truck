package dto

type DriverTrucks struct {
	DriverID   string  `json:"driver_id"`
	DriverName string  `json:"driver_name"`
	LicenseID  string  `json:"license_id"`
	Trucks     []Truck `json:"trucks"`
}
