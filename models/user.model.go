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
	Username   string    `gorm:"type:varchar(255);not null" json:"username"`
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
	Username   string    `gorm:"type:varchar(255);not null" json:"username"`
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
	ID                 uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"id"`
	HN                 string    `json:"hn,omitempty"`
	Alias              string    `json:"alias,omitempty"`
	FirstName          string    `json:"firstName,omitempty"`
	LastName           string    `json:"lastName,omitempty"`
	YearOfBirth        string    `json:"yearOfBirth,omitempty"`
	Gender             string    `json:"gender,omitempty"`
	MainDoctorId       uuid.UUID `gorm:"type:uuid" json:"mainDoctorId,omitempty"`
	MainDoctor         Doctor    `gorm:"foreignKey:MainDoctorId;" json:"mainDoctor,omitempty"`
	AssistanceDoctorId uuid.UUID `gorm:"type:uuid" json:"assistanceDoctorId,omitempty"`
	AssistanceDoctor   Doctor    `gorm:"foreignKey:AssistanceDoctorId;" json:"assistanceDoctor,omitempty"`
	Photo              string    `json:"photo,omitempty"`
}

func (Patient) TableName() string {
	return "patient"
}

type SignUpInput struct {
	Username        string `json:"username" binding:"required"`
	Password        string `json:"password" binding:"required,min=8"`
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

type UpdateProfilePatient struct {
	HN          string `json:"hn,omitempty"`
	Alias       string `json:"alias,omitempty"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	YearOfBirth string `json:"yearOfBirth,omitempty"`
	Gender      string `json:"gender,omitempty"`
	Photo       string `json:"photo,omitempty"`
}
