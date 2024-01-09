package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ployns/Metabolic-Syndrome-Backend/models"
	"gorm.io/gorm"
)

type PlanController struct {
	DB *gorm.DB
}

func NewPlanController(DB *gorm.DB) PlanController {
	return PlanController{DB}
}

// Create Plan
func (pc *PlanController) CreatePlan(ctx *gin.Context) {
	var payload = struct {
		Name        string          `json:"name,omitempty"`
		DiseaseRisk string          `json:"diseaseRisk,omitempty"`
		Description string          `json:"description,omitempty"`
		Photo       string          `json:"photo,omitempty"`
		Monday      []models.Detail `json:"monday,omitempty"`
		Tuesday     []models.Detail `json:"tuesday,omitempty"`
		Wednesday   []models.Detail `json:"wednesday,omitempty"`
		Thursday    []models.Detail `json:"thursday,omitempty"`
		Friday      []models.Detail `json:"friday,omitempty"`
		Saturday    []models.Detail `json:"saturday,omitempty"`
		Sunday      []models.Detail `json:"sunday,omitempty"`
	}{}

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	// Marshal the nested structures to JSON
	mondayJSON, _ := json.Marshal(payload.Monday)
	tuesdayJSON, _ := json.Marshal(payload.Tuesday)
	wednesdayJSON, _ := json.Marshal(payload.Wednesday)
	thursdayJSON, _ := json.Marshal(payload.Thursday)
	fridayJSON, _ := json.Marshal(payload.Friday)
	saturdayJSON, _ := json.Marshal(payload.Saturday)
	sundayJSON, _ := json.Marshal(payload.Sunday)

	// Create a new Plan with the JSON data
	newPlan := models.Plan{
		Name:        payload.Name,
		DiseaseRisk: payload.DiseaseRisk,
		Description: payload.Description,
		Photo:       payload.Photo,
		Monday:      json.RawMessage(mondayJSON),
		Tuesday:     json.RawMessage(tuesdayJSON),
		Wednesday:   json.RawMessage(wednesdayJSON),
		Thursday:    json.RawMessage(thursdayJSON),
		Friday:      json.RawMessage(fridayJSON),
		Saturday:    json.RawMessage(saturdayJSON),
		Sunday:      json.RawMessage(sundayJSON),
	}

	// Save the new Plan to the database
	if err := pc.DB.Create(&newPlan).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "Can not create plan"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "message": "Create plan success"})
}
