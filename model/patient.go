package model

import "time"

type Patient struct {
	Model
	PatientId      string           `json:"patient_id"`
	FirstName      string           `json:"first_name"`
	LastName       string           `json:"last_name"`
	DateOfBirth    time.Time        `json:"date_of_birth"`
	Gender         string           `json:"gender"`
	Phone          string           `json:"phone"`
	Email          string           `json:"email"`
	Address        string           `json:"address"`
	ImageUrl       string           `json:"image_url"`
	MedicalHistory []MedicalHistory `gorm:"-" json:"medical_history"`
}
