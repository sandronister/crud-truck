package dto

type Driver struct {
	ID        string `json:"id,omitempty"`
	Name      string `json:"name"`
	LicenseID string `json:"license_id"`
	CreatedAt string `json:"created_at,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`
}
