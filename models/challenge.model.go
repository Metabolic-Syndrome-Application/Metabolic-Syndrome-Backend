package models

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/google/uuid"
)

// ภารกิจทั่วไป
type DairyChallenge struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"id"`
	UpdatedAt time.Time `json:"updateAt"`
}

func (DairyChallenge) TableName() string {
	return "dairyChallenge"
}

// ตอบคำถามประจำวัน
type QuizChallenge struct {
	ID        uuid.UUID       `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"id"`
	Question  string          `json:"question"`
	Choices   json.RawMessage `gorm:"type:json" json:"choices"`
	Point     int             `json:"point"`
	Time      int             `json:"time"`
	Status    string          `gorm:"default:'active' " json:"status"`
	CreatedAt time.Time       `json:"createdAt"`
	UpdatedAt time.Time       `json:"updatedAt"`
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
