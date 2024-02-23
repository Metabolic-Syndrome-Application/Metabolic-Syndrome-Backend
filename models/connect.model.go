package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

type Connect struct {
	ID                 uuid.UUID      `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"id"`
	OTP                string         `gorm:"not null" json:"otp"`
	HN                 string         `gorm:"not null" json:"hn"`
	FirstName          string         `gorm:"not null" json:"firstName"`
	LastName           string         `gorm:"not null" json:"lastName"`
	YearOfBirth        int            `gorm:"not null" json:"yearOfBirth"`
	Gender             string         `gorm:"not null" json:"gender"`
	MainDoctorID       *uuid.UUID     `gorm:"type:uuid;not null" json:"mainDoctorID"`
	AssistanceDoctorID *uuid.UUID     `gorm:"type:uuid" json:"assistanceDoctorID"`
	DiseaseRisk        DiseaseRisk    `gorm:"type:jsonb" json:"diseaseRisk"`
	PlanID             pq.StringArray `gorm:"type:uuid[];column:plan_id;null" json:"planID"`
	ExpiresIn          time.Time      `gorm:"not null" json:"expiresIn"`
	CreatedAt          time.Time      `gorm:"not null" json:"createdAt"`
	UpdatedAt          time.Time      `gorm:"not null" json:"updatedAt"`
}

func (Connect) TableName() string {
	return "connect"
}
