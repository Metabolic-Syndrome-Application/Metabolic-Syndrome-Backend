package models

import (
	"time"

	"github.com/google/uuid"
)

type Connect struct {
	ID                 uuid.UUID  `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"id"`
	OTP                string     `json:"otp"`
	HN                 *string    `json:"hn"`
	FirstName          string     `json:"firstName"`
	LastName           string     `json:"lastName"`
	YearOfBirth        int        `json:"yearOfBirth"`
	Gender             string     `json:"gender"`
	MainDoctorID       *uuid.UUID `gorm:"type:uuid" json:"mainDoctorID"`
	AssistanceDoctorID *uuid.UUID `gorm:"type:uuid" json:"assistanceDoctorID"`
	Disease            *string    `json:"disease"`
	ExpiresIn          time.Time  `json:"expiresIn"`
	CreatedAt          time.Time  `json:"createdAt"`
	UpdatedAt          time.Time  `json:"updatedAt"`
}

func (Connect) TableName() string {
	return "connect"
}
