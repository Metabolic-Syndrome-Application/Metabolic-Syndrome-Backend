package models

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/google/uuid"
)

type Plan struct {
	ID          uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"id"`
	Name        string    `gorm:"not null" json:"name"`
	Description string    `gorm:"not null" json:"description"`
	Photo       string    `gorm:"not null" json:"photo"`
	Type        string    `gorm:"not null" json:"type"`
	Detail      Detail    `gorm:"type:json" json:"detail"`
	CreatedAt   time.Time `gorm:"not null" json:"createdAt"`
	UpdatedAt   time.Time `gorm:"not null" json:"updatedAt"`
}

type Detail struct {
	Name []string `json:"name"`
	Day  []string `json:"day"`
}

func (dr *Detail) Scan(value interface{}) error {
	if data, ok := value.([]byte); ok {
		return json.Unmarshal(data, dr)
	}
	return errors.New("failed to unmarshal Detail")
}

func (Plan) TableName() string {
	return "plan"
}
