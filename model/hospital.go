package model

type Hospital struct {
	Model
	HospitalId string  `json:"hospital_id"`
	Name       string  `json:"name"`
	Type       string  `json:"type"`
	Province   string  `json:"province"`
	City       string  `json:"city"`
	Address    string  `json:"address"`
	Latitude   float64 `json:"latitude"`
	Longitude  float64 `json:"longitude"`
}
