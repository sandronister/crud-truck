package dto

type Link struct {
	ID        string `json:"id"`
	TruckID   string `json:"truck_id"`
	DriverID  string `json:"driver_id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
