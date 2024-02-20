package models

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/google/uuid"
)

type Plan struct {
	ID          uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Photo       string    `json:"photo"`
	Type        string    `json:"type"`
	Detail      Detail    `gorm:"type:json" json:"detail"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
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
