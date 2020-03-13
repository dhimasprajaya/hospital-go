package model

import "github.com/golang/protobuf/ptypes/timestamp"

type MedicalHistory struct {
	Model
	PatientId  int                 `json:"patient_id"`
	HospitalId int                 `json:"hospital_id"`
	DoctorId   int                 `json:"doctor_id"`
	Timestamp  timestamp.Timestamp `json:"timestamp"`
}
