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

// Update plan
func (pc *PlanController) UpdatePlan(ctx *gin.Context) {
	planID := ctx.Param("id")
	var plan models.Plan
	result := pc.DB.First(&plan, "id = ?", planID)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "Not have this ID"})
		return
	}

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

	updatePlan := models.Plan{
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

	if err := pc.DB.Model(&plan).Updates(updatePlan).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "Can not update plan"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "message": "Update plan success"})
}

// Get plan
func (pc *PlanController) GetPlan(ctx *gin.Context) {
	planID := ctx.Param("id")
	var plan models.Plan
	result := pc.DB.First(&plan, "id = ?", planID)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "Not have this ID"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": gin.H{"plan": plan}})

}

// Get all plan
func (pc *PlanController) GetAllPlan(ctx *gin.Context) {

}

// Delete plan
func (pc *PlanController) DeletePlan(ctx *gin.Context) {

}
