package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

type User struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	Username  string    `gorm:"type:varchar(255);not null" json:"username"`
	Password  string    `gorm:"not null"`
	Role      string    `gorm:"type:varchar(255);not null" json:"role"`
	Verified  bool      `gorm:"not null" json:"verified"`
	Type      string    `gorm:"not null" json:"type"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (User) TableName() string {
	return "user"
}

type Doctor struct {
	ID         uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"id"`
	Username   string    `gorm:"type:varchar(255);not null" json:"username"`
	Prefix     string    `json:"prefix"`
	FirstName  string    `json:"firstName"`
	LastName   string    `json:"lastName"`
	Gender     string    `json:"gender"`
	Department string    `json:"department"`
	Specialist string    `json:"specialist"`
	UpdatedAt  time.Time `json:"updateAt"`
}

func (Doctor) TableName() string {
	return "doctor"
}

type Staff struct {
	ID         uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"id"`
	Username   string    `gorm:"type:varchar(255);not null" json:"username"`
	Prefix     string    `json:"prefix"`
	FirstName  string    `json:"firstName"`
	LastName   string    `json:"lastName"`
	Gender     string    `json:"gender"`
	Department string    `json:"department"`
	Specialist string    `json:"specialist"`
	UpdatedAt  time.Time `json:"updateAt"`
}

func (Staff) TableName() string {
	return "staff"
}

type Patient struct {
	ID                 uuid.UUID      `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"id"`
	HN                 *string        `json:"hn"`
	Alias              string         `json:"alias"`
	FirstName          string         `json:"firstName"`
	LastName           string         `json:"lastName"`
	YearOfBirth        int            `json:"yearOfBirth"`
	Gender             string         `json:"gender"`
	Occupation         string         `json:"occupation"`
	Photo              string         `json:"photo"`
	MainDoctorID       *uuid.UUID     `gorm:"type:uuid ;null" json:"mainDoctorID"`
	MainDoctor         Doctor         `gorm:"foreignKey:MainDoctorID; " json:"mainDoctor"`
	AssistanceDoctorID *uuid.UUID     `gorm:"type:uuid ;null" json:"assistanceDoctorID"`
	AssistanceDoctor   Doctor         `gorm:"foreignKey:AssistanceDoctorID;" json:"assistanceDoctor"`
	DiseaseRisk        DiseaseRisk    `gorm:"type:jsonb" json:"diseaseRisk"`
	PlanID             pq.StringArray `gorm:"type:uuid[];column:plan_id" json:"planID"`
	Plan               []Plan         `gorm:"many2many:patient_plan;association_foreignkey:ID;joinForeignKey:PatientID;References:ID;joinReferences:PlanID" json:"Plan"`
	ChallengeID        *uuid.UUID     `gorm:"type:uuid ;null" json:"challengeID"`
	Challenge          DairyChallenge `gorm:"foreignKey:ChallengeID;" json:"challenge"`
	CollectPoints      int            `json:"collectPoints"`
	Status             string         `gorm:"default:'in process' " json:"status"`
	UpdatedAt          time.Time      `json:"updateAt"`
}
type DiseaseRisk struct {
	Diabetes       string `json:"diabetes"`
	Hyperlipidemia string `json:"hyperlipidemia"`
	Hypertension   string `json:"hypertension"`
	Obesity        string `json:"obesity"`
}

func (dr *DiseaseRisk) Scan(value interface{}) error {
	if data, ok := value.([]byte); ok {
		return json.Unmarshal(data, dr)
	}
	return errors.New("failed to unmarshal DiseaseRisk")
}

func (patient *Patient) Scan(value interface{}) error {
	if value == nil {
		return nil
	}

	switch v := value.(type) {
	case pq.StringArray:
		patient.PlanID = v
		return nil
	default:
		return errors.New("unsupported type for PlanID")
	}
}

func (patient *Patient) Value() (driver.Value, error) {
	if patient.PlanID == nil {
		return nil, nil
	}
	return patient.PlanID.Value()
}

func (Patient) TableName() string {
	return "patient"
}

type SignUpInput struct {
	Username        string `json:"username" binding:"required"`
	Password        string `json:"password" binding:"required"`
	PasswordConfirm string `json:"passwordConfirm" binding:"required"`
	Role            string `json:"role" binding:"required"`
}

type SignInInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password"  binding:"required"`
}

type UserResponse struct {
	ID        uuid.UUID `json:"id"`
	Username  string    `gorm:"type:varchar(255);not null" json:"username"`
	Role      string    `json:"role"`
	Photo     string    `json:"photo"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
