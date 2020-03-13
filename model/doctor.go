package model

type Doctor struct {
	Model
	DoctorId      string `json:"doctor_id"`
	FirstName     string `json:"first_name"`
	LastName      string `json:"last_name"`
	Specialist    string `json:"specialist"`
	LicenseNumber string `json:"license_number"`
	Phone         string `json:"phone"`
	ImageUrl      string `json:"image_url"`
}
