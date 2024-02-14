// package models

// import (
// 	"encoding/json"
// 	"errors"
// 	"time"

// 	"github.com/google/uuid"
// )

// type Plan struct {
// 	ID          uuid.UUID       `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"id"`
// 	Name        string          `json:"name,omitempty"`
// 	DiseaseRisk string          `json:"diseaseRisk,omitempty"`
// 	Description string          `json:"description,omitempty"`
// 	Photo       string          `json:"photo,omitempty"`
// 	Monday      json.RawMessage `gorm:"type:jsonb" json:"monday,omitempty"`
// 	Tuesday     json.RawMessage `gorm:"type:jsonb" json:"tuesday,omitempty"`
// 	Wednesday   json.RawMessage `gorm:"type:jsonb" json:"wednesday,omitempty"`
// 	Thursday    json.RawMessage `gorm:"type:jsonb" json:"thursday,omitempty"`
// 	Friday      json.RawMessage `gorm:"type:jsonb" json:"friday,omitempty"`
// 	Saturday    json.RawMessage `gorm:"type:jsonb" json:"saturday,omitempty"`
// 	Sunday      json.RawMessage `gorm:"type:jsonb" json:"sunday,omitempty"`
// 	CreatedAt   time.Time
// 	UpdatedAt   time.Time
// }

// type Detail struct {
// 	Name   string `json:"name,omitempty"`
// 	Type   string `json:"type,omitempty"`
// 	Status *bool  `gorm:"default:false" json:"status,omitempty"`
// }

// func (dr *Detail) Scan(value interface{}) error {
// 	if data, ok := value.([]byte); ok {
// 		return json.Unmarshal(data, dr)
// 	}
// 	return errors.New("failed to unmarshal Detail")
// }
// func (Plan) TableName() string {
// 	return "plan"
// }

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
