package models

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/google/uuid"
)

type RecordHealth struct {
	ID                     uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"id"`
	PatientID              uuid.UUID `gorm:"type:uuid" json:"patientID"`
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
	Timestamp              time.Time `json:"timestamp"`
}

func (RecordHealth) TableName() string {
	return "recordHealth"
}

type RecordPlan struct {
	ID        uuid.UUID       `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"id"`
	PatientID uuid.UUID       `gorm:"type:uuid" json:"patientID"`
	List      json.RawMessage `gorm:"type:json" json:"list"`
	Mood      *string         `json:"mood"`
	GetPoint  bool            `gorm:"default:false" json:"getPoint"`
	CreatedAt time.Time       `json:"createdAt"`
	UpdatedAt time.Time       `json:"updatedAt"`
}

type List struct {
	Name  string `json:"name"`
	Check bool   `gorm:"default:false" json:"check"`
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

type RecordQuiz struct {
	ID              uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"id"`
	PatientID       uuid.UUID `gorm:"type:uuid" json:"patientID"`
	QuizChallengeID uuid.UUID `gorm:"type:uuid" json:"quizChallengeID"`
	CreatedAt       time.Time `json:"createdAt"`
}

func (RecordQuiz) TableName() string {
	return "recordQuiz"
}

type RecordDaily struct {
	ID               uuid.UUID       `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"id"`
	PatientID        uuid.UUID       `gorm:"type:uuid" json:"patientID"`
	DailyChallengeID uuid.UUID       `gorm:"type:uuid" json:"dailyChallengeID"`
	Day              int             `json:"day"`
	List             json.RawMessage `gorm:"type:json" json:"list"`
	StartDate        string          `json:"startDate"`
	EndDate          string          `json:"endDate"`
	UpdatedAt        time.Time       `json:"updatedAt"`
	CreatedAt        time.Time       `json:"createdAt"`
}

func (RecordDaily) TableName() string {
	return "recordDaily"
}
