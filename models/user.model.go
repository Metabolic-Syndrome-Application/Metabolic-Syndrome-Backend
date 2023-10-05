package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID       uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	Username string    `gorm:"type:varchar(255);not null"`
	// Email     string    `gorm:"uniqueIndex;not null"`
	Password  string `gorm:"not null"`
	Role      string `gorm:"type:varchar(255);not null"`
	Verified  bool   `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Patient struct {
	ID       uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	Username string    `gorm:"type:varchar(255);not null"`
	// Email     string    `gorm:"uniqueIndex;not null"`
	Password  string `gorm:"not null"`
	Role      string `gorm:"type:varchar(255);not null"`
	Photo     string
	Verified  bool `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Staff struct { // doctor nurse admin
	ID       uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	Username string    `gorm:"type:varchar(255);not null"`
	// Email     string    `gorm:"uniqueIndex;not null"`
	Password  string `gorm:"not null"`
	Role      string `gorm:"type:varchar(255);not null"`
	Verified  bool   `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type SignUpInput struct {
	Username string `json:"name" binding:"required"`
	// Email           string `json:"email" binding:"required"`
	Password        string `json:"password" binding:"required,min=8"`
	PasswordConfirm string `json:"passwordConfirm" binding:"required"`
}

type SignInInput struct {
	Username string `json:"name" binding:"required"`
	// Email    string `json:"email"  binding:"required"`
	Password string `json:"password"  binding:"required"`
}

type UserResponse struct {
	ID       uuid.UUID `json:"id,omitempty"`
	Username string    `gorm:"type:varchar(255);not null"`
	// Email     string    `json:"email,omitempty"`
	Role      string    `json:"role,omitempty"`
	Photo     string    `json:"photo,omitempty"`
	Provider  string    `json:"provider"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
