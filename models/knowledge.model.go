package models

import (
	"time"

	"github.com/google/uuid"
)

type Knowledge struct {
	ID          uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"id"`
	Disease     string    `json:"disease"`
	Name        string    `json:"name"`
	Details     string    `json:"details"`
	Symptoms    string    `json:"symptoms"`
	Medications string    `json:"medications"`
	Behaviors   string    `json:"behaviors"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

func (Knowledge) TableName() string {
	return "knowledge"
}
