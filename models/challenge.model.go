package models

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/google/uuid"
)

// ภารกิจทั่วไป
type DailyChallenge struct {
	ID           uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"id"`
	Name         string    `gorm:"not null" json:"name"`
	Description  string    `gorm:"not null" json:"description"`
	Photo        string    `gorm:"not null" json:"photo"`
	Detail       Detail    `gorm:"type:json" json:"detail"`
	Points       int       `gorm:"not null" json:"points"`
	NumDays      int       `gorm:"not null" json:"numDays"`
	Participants int       `gorm:"not null" json:"participants"`
	Status       string    `gorm:"not null" gorm:"default:'active' " json:"status"`
	CreatedAt    time.Time `gorm:"not null" json:"createdAt"`
	UpdatedAt    time.Time `gorm:"not null" json:"updatedAt"`
}

func (DailyChallenge) TableName() string {
	return "dailyChallenge"
}

// ตอบคำถามประจำวัน
type QuizChallenge struct {
	ID        uuid.UUID       `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"id"`
	Question  string          `gorm:"not null" json:"question"`
	Choices   json.RawMessage `gorm:"type:json" json:"choices"`
	Points    int             `gorm:"not null" json:"points"`
	LimitTime int             `gorm:"not null" json:"limitTime"`
	CreatedAt time.Time       `gorm:"not null" json:"createdAt"`
	UpdatedAt time.Time       `gorm:"not null" json:"updatedAt"`
}

type Choices struct {
	Option    string `json:"option"`
	IsCorrect bool   `json:"isCorrect"`
}

func (dr *Choices) Scan(value interface{}) error {
	if data, ok := value.([]byte); ok {
		return json.Unmarshal(data, dr)
	}
	return errors.New("failed to unmarshal Choices")
}

func (QuizChallenge) TableName() string {
	return "quizChallenge"
}
