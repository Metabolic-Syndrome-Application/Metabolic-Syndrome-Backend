package models

import (
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
