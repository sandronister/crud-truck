package dto

type Truck struct {
	ID           string `json:"id,omitempty"`
	Brand        string `json:"brand"`
	Model        string `json:"model"`
	Year         int    `json:"year"`
	LicensePlate string `json:"license_plate"`
	CreatedAt    string `json:"created_at,omitempty" `
	UpdatedAt    string `json:"updated_at,omitempty"`
}
