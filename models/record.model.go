package models

import (
	"time"

	"github.com/google/uuid"
)

type RecordHealth struct {
	ID                     uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"id"`
	PatientID              uuid.UUID `gorm:"type:uuid ;null" json:"patientID,omitempty"`
	Height                 float32   `json:"height,omitempty"`
	Weight                 float32   `json:"weight,omitempty"`
	Waistline              float32   `json:"waistline,omitempty"`
	SystolicBloodPressure  int       `json:"systolicBloodPressure,omitempty"`
	DiastolicBloodPressure int       `json:"diastolicBloodPressure,omitempty"`
	PulseRate              int       `json:"pulseRate,omitempty"`
	BloodGlucose           float32   `json:"bloodGlucose,omitempty"`
	Cholesterol            float32   `json:"cholesterol,omitempty"`
	HDL                    float32   `json:"hdl,omitempty"`
	LDL                    float32   `json:"ldl,omitempty"`
	Triglyceride           float32   `json:"triglyceride,omitempty"`
	RecordBy               string    `json:"recordBy,omitempty"`
	Timestamp              time.Time
}

func (RecordHealth) TableName() string {
	return "recordHealth"
}

// type RecordPlan struct {
// 	ID        uuid.UUID       `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"id"`
// 	PatientID uuid.UUID       `gorm:"type:uuid ;null" json:"patientID,omitempty"`
// 	Patient   Patient         `gorm:"foreignKey:PatientID; " json:"patient,omitempty"`
// 	Detail    json.RawMessage `gorm:"type:jsonb" json:"detail,omitempty"`
// 	Mood      string          `json:"mood,omitempty"`
// 	GetPoint  bool            `gorm:"default:false" json:"getPoint,omitempty"`
// 	Timestamp time.Time
// }

// func (RecordPlan) TableName() string {
// 	return "recordPlan"
// }
