package models

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/google/uuid"
)

type RecordHealth struct {
	ID                     uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"id"`
	PatientID              uuid.UUID `gorm:"type:uuid ;null" json:"patientID"`
	Height                 float32   `json:"height"`
	Weight                 float32   `json:"weight"`
	Waistline              float32   `json:"waistline"`
	SystolicBloodPressure  int       `json:"systolicBloodPressure"`
	DiastolicBloodPressure int       `json:"diastolicBloodPressure"`
	PulseRate              int       `json:"pulseRate"`
	BloodGlucose           float32   `json:"bloodGlucose"`
	Cholesterol            float32   `json:"cholesterol"`
	HDL                    float32   `json:"hdl"`
	LDL                    float32   `json:"ldl"`
	Triglyceride           float32   `json:"triglyceride"`
	RecordBy               string    `json:"recordBy"`
	Timestamp              time.Time
}

func (RecordHealth) TableName() string {
	return "recordHealth"
}

type RecordPlan struct {
	ID        uuid.UUID       `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"id"`
	PatientID uuid.UUID       `gorm:"type:uuid ;null" json:"patientID"`
	List      json.RawMessage `gorm:"type:json" json:"list"`
	Mood      *string         `json:"mood"`
	GetPoint  bool            `gorm:"default:false" json:"getPoint"`
	CreatedAt time.Time       `json:"createAt"`
	UpdatedAt time.Time       `json:"updateAt"`
}

type List struct {
	Name  string `json:"name"`
	Check string `gorm:"default:false" json:"check"`
	Type  string `json:"type"`
}

func (dr *List) Scan(value interface{}) error {
	if data, ok := value.([]byte); ok {
		return json.Unmarshal(data, dr)
	}
	return errors.New("failed to unmarshal List")
}

func (RecordPlan) TableName() string {
	return "recordPlan"
}
