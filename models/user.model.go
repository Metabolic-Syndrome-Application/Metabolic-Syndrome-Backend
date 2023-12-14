package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	Username  string    `gorm:"type:varchar(255);not null"`
	Password  string    `gorm:"not null"`
	Role      string    `gorm:"type:varchar(255);not null"`
	Verified  bool      `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (User) TableName() string {
	return "user"
}

type Doctor struct {
	ID         uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"id"`
	Prefix     string    `json:"prefix,omitempty"`
	FirstName  string    `json:"firstName,omitempty"`
	LastName   string    `json:"lastName,omitempty"`
	Gender     string    `json:"gender,omitempty"`
	Department string    `json:"department,omitempty"`
	Specialist string    `json:"specialist,omitempty"`
}

func (Doctor) TableName() string {
	return "doctor"
}

type Staff struct {
	ID         uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"id"`
	Prefix     string    `json:"prefix,omitempty"`
	FirstName  string    `json:"firstName,omitempty"`
	LastName   string    `json:"lastName,omitempty"`
	Gender     string    `json:"gender,omitempty"`
	Department string    `json:"department,omitempty"`
	Specialist string    `json:"specialist,omitempty"`
}

func (Staff) TableName() string {
	return "staff"
}

type Patient struct {
	ID                 uuid.UUID  `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"id"`
	HN                 string     `json:"hn,omitempty"`
	Alias              string     `json:"alias,omitempty"`
	FirstName          string     `json:"firstName,omitempty"`
	LastName           string     `json:"lastName,omitempty"`
	YearOfBirth        int        `json:"yearOfBirth,omitempty"`
	Gender             string     `json:"gender,omitempty"`
	Photo              string     `json:"photo,omitempty"`
	MainDoctorId       *uuid.UUID `gorm:"type:uuid ;null" json:"mainDoctorId,omitempty"`
	MainDoctor         Doctor     `gorm:"foreignKey:MainDoctorId; " json:"mainDoctor,omitempty"`
	AssistanceDoctorId *uuid.UUID `gorm:"type:uuid ;null" json:"assistanceDoctorId,omitempty"`
	AssistanceDoctor   Doctor     `gorm:"foreignKey:AssistanceDoctorId;" json:"assistanceDoctor,omitempty"`
	DiseaseRisk        string     `json:"diseaseRisk,omitempty"`
	PlanID             *uuid.UUID `gorm:"type:uuid ;null" json:"planID,omitempty"`
	Plan               Plan       `gorm:"foreignKey:PlanID;" json:"plan,omitempty"`
	ChallengeID        *uuid.UUID `gorm:"type:uuid ;null" json:"challengeID,omitempty"`
	Challenge          Challenge  `gorm:"foreignKey:ChallengeID;" json:"challenge,omitempty"`
	CollectPoints      int        `json:"collectPoints,omitempty"`
	Status             string     `gorm:"default:'in process' " json:"status,omitempty"`
}

func (Patient) TableName() string {
	return "patient"
}

type Plan struct {
	PlanID uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"id"`
}

func (Plan) TableName() string {
	return "plan"
}

type Challenge struct {
	ChallengeID uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"id"`
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
	Username  string    `gorm:"type:varchar(255);not null"`
	Role      string    `json:"role,omitempty"`
	Photo     string    `json:"photo,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
