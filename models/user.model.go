package models

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	Username  string    `gorm:"type:varchar(255);not null" json:"username"`
	Password  string    `gorm:"not null"`
	Role      string    `gorm:"type:varchar(255);not null" json:"role"`
	Verified  bool      `gorm:"not null" json:"verified"`
	Type      string    `gorm:"not null" json:"type"`
	CreatedAt time.Time `json:"createAt"`
	UpdatedAt time.Time `json:"updateAt"`
}

func (User) TableName() string {
	return "user"
}

type Doctor struct {
	ID         uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"id"`
	Username   string    `gorm:"type:varchar(255);not null" json:"username"`
	Prefix     string    `json:"prefix,omitempty"`
	FirstName  string    `json:"firstName,omitempty"`
	LastName   string    `json:"lastName,omitempty"`
	Gender     string    `json:"gender,omitempty"`
	Department string    `json:"department,omitempty"`
	Specialist string    `json:"specialist,omitempty"`
	UpdatedAt  time.Time `json:"updateAt"`
}

func (Doctor) TableName() string {
	return "doctor"
}

type Staff struct {
	ID         uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"id"`
	Username   string    `gorm:"type:varchar(255);not null" json:"username"`
	Prefix     string    `json:"prefix,omitempty"`
	FirstName  string    `json:"firstName,omitempty"`
	LastName   string    `json:"lastName,omitempty"`
	Gender     string    `json:"gender,omitempty"`
	Department string    `json:"department,omitempty"`
	Specialist string    `json:"specialist,omitempty"`
	UpdatedAt  time.Time `json:"updateAt"`
}

func (Staff) TableName() string {
	return "staff"
}

type Patient struct {
	ID                 uuid.UUID   `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"id"`
	HN                 string      `json:"hn,omitempty"`
	Alias              string      `json:"alias,omitempty"`
	FirstName          string      `json:"firstName,omitempty"`
	LastName           string      `json:"lastName,omitempty"`
	YearOfBirth        int         `json:"yearOfBirth,omitempty"`
	Gender             string      `json:"gender,omitempty"`
	Occupation         string      `json:"occupation,omitempty"`
	Photo              string      `json:"photo,omitempty"`
	MainDoctorID       *uuid.UUID  `gorm:"type:uuid ;null" json:"mainDoctorID,omitempty"`
	MainDoctor         Doctor      `gorm:"foreignKey:MainDoctorID; " json:"mainDoctor,omitempty"`
	AssistanceDoctorID *uuid.UUID  `gorm:"type:uuid ;null" json:"assistanceDoctorID,omitempty"`
	AssistanceDoctor   Doctor      `gorm:"foreignKey:AssistanceDoctorID;" json:"assistanceDoctor,omitempty"`
	DiseaseRisk        DiseaseRisk `gorm:"type:jsonb" json:"diseaseRisk,omitempty"`
	PlanID             *uuid.UUID  `gorm:"type:uuid ;null" json:"planID,omitempty"`
	Plan               Plan        `gorm:"foreignKey:PlanID;" json:"plan,omitempty"`
	ChallengeID        *uuid.UUID  `gorm:"type:uuid ;null" json:"challengeID,omitempty"`
	Challenge          Challenge   `gorm:"foreignKey:ChallengeID;" json:"challenge,omitempty"`
	CollectPoints      int         `json:"collectPoints,omitempty"`
	Status             string      `gorm:"default:'in process' " json:"status,omitempty"`
	UpdatedAt          time.Time   `json:"updateAt"`
}
type DiseaseRisk struct {
	Diabetes       string `json:"diabetes"`
	Hyperlipidemia string `json:"hyperlipidemia"`
	Hypertension   string `json:"hypertension"`
	Obesity        string `json:"obesity"`
}

func (dr *DiseaseRisk) Scan(value interface{}) error {
	if data, ok := value.([]byte); ok {
		return json.Unmarshal(data, dr)
	}
	return errors.New("failed to unmarshal DiseaseRisk")
}

func (Patient) TableName() string {
	return "patient"
}

type Challenge struct {
	ChallengeID uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"id"`
	UpdatedAt   time.Time `json:"updateAt"`
}

func (Challenge) TableName() string {
	return "challenge"
}

type SignUpInput struct {
	Username        string `json:"username" binding:"required"`
	Password        string `json:"password" binding:"required"`
	PasswordConfirm string `json:"passwordConfirm" binding:"required"`
	Role            string `json:"role" binding:"required"`
}

type SignInInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password"  binding:"required"`
}

type UserResponse struct {
	ID        uuid.UUID `json:"id,omitempty"`
	Username  string    `gorm:"type:varchar(255);not null" json:"username"`
	Role      string    `json:"role,omitempty"`
	Photo     string    `json:"photo,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
