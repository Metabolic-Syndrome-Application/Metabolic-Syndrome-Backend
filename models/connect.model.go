package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

type Connect struct {
	ID                 uuid.UUID      `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"id"`
	OTP                string         `json:"otp"`
	HN                 *string         `json:"hn"`
	FirstName          string         `json:"firstName"`
	LastName           string         `json:"lastName"`
	YearOfBirth        int            `json:"yearOfBirth"`
	Gender             string         `json:"gender"`
	MainDoctorID       *uuid.UUID     `gorm:"type:uuid" json:"mainDoctorID"`
	AssistanceDoctorID *uuid.UUID     `gorm:"type:uuid" json:"assistanceDoctorID"`
	DiseaseRisk        DiseaseRisk    `gorm:"type:jsonb" json:"diseaseRisk"`
	PlanID             pq.StringArray `gorm:"type:uuid[];column:plan_id" json:"planID"`
	ExpiresIn          time.Time      `json:"expiresIn"`
	CreatedAt          time.Time      `json:"createdAt"`
	UpdatedAt          time.Time      `json:"updatedAt"`
}

func (Connect) TableName() string {
	return "connect"
}
