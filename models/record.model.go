package models

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/google/uuid"
)

type RecordHealth struct {
	ID                     uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();not null;primary_key" json:"id"`
	PatientID              uuid.UUID `gorm:"type:uuid ;not null" json:"patientID"`
	Height                 float32   `gorm:"not null" json:"height"`
	Weight                 float32   `gorm:"not null" json:"weight"`
	Waistline              float32   `gorm:"not null" json:"waistline"`
	SystolicBloodPressure  int       `gorm:"not null" json:"systolicBloodPressure"`
	DiastolicBloodPressure int       `gorm:"not null" json:"diastolicBloodPressure"`
	PulseRate              int       `gorm:"not null" json:"pulseRate"`
	BloodGlucose           float32   `gorm:"not null" json:"bloodGlucose"`
	Cholesterol            float32   `gorm:"not null" json:"cholesterol"`
	HDL                    float32   `gorm:"not null" json:"hdl"`
	LDL                    float32   `gorm:"not null" json:"ldl"`
	Triglyceride           float32   `gorm:"not null" json:"triglyceride"`
	RecordBy               string    `gorm:"not null" json:"recordBy"`
	Timestamp              time.Time `gorm:"not null"`
}

func (RecordHealth) TableName() string {
	return "recordHealth"
}

type RecordPlan struct {
	ID        uuid.UUID       `gorm:"type:uuid;default:uuid_generate_v4();not null;primary_key" json:"id"`
	PatientID uuid.UUID       `gorm:"type:uuid ;not null" json:"patientID"`
	List      json.RawMessage `gorm:"type:json" json:"list"`
	Mood      *string         `json:"mood"`
	GetPoint  bool            `gorm:"default:false;not null" json:"getPoint"`
	CreatedAt time.Time       `gorm:"not null" json:"createdAt"`
	UpdatedAt time.Time       `gorm:"not null" json:"updatedAt"`
}

type List struct {
	Name  string `gorm:"not null" json:"name"`
	Check bool   `gorm:"default:false;not null" json:"check"`
	Type  string `gorm:"not null" json:"type"`
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
	ID              uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();not null;primary_key" json:"id"`
	PatientID       uuid.UUID `gorm:"type:uuid ;not null" json:"patientID"`
	QuizChallengeID uuid.UUID `gorm:"type:uuid ;not null" json:"quizChallengeID"`
	CreatedAt       time.Time `gorm:"not null" json:"createdAt"`
}

func (RecordQuiz) TableName() string {
	return "recordQuiz"
}

type RecordDaily struct {
	ID               uuid.UUID       `gorm:"type:uuid;not null;primary_key" json:"id"`
	PatientID        uuid.UUID       `gorm:"type:uuid ;not null" json:"patientID"`
	DailyChallengeID uuid.UUID       `gorm:"type:uuid ;not null" json:"dailyChallengeID"`
	Day              int             `gorm:"not null" json:"day"`
	List             json.RawMessage `gorm:"type:json;not null" json:"list"`
	StartDate        string          `gorm:"not null" json:"startDate"`
	EndDate          string          `gorm:"not null" json:"endDate"`
	UpdatedAt        time.Time       `gorm:"not null" json:"updatedAt"`
	CreatedAt        time.Time       `gorm:"not null" json:"createdAt"`
}

func (RecordDaily) TableName() string {
	return "recordDaily"
}
