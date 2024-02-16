package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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
		Name        string        `json:"name,omitempty"`
		Description string        `json:"description,omitempty"`
		Photo       string        `json:"photo,omitempty"`
		Type        string        `json:"type,omitempty"`
		Detail      models.Detail `gorm:"type:json" json:"detail,omitempty"`
	}{}
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	var plan models.Plan
	existingPlan := pc.DB.First(&plan, "name = ?", payload.Name)
	if existingPlan.Error == nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "This plan's name is already in use"})
		return
	}
	newPlan := models.Plan{
		Name:        payload.Name,
		Description: payload.Description,
		Photo:       payload.Photo,
		Type:        payload.Type,
		Detail:      payload.Detail,
	}
	if err := pc.DB.Create(&newPlan).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "Can not create plan"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "message": "Create plan success"})
}

// Update Plan
func (pc *PlanController) UpdatePlan(ctx *gin.Context) {
	planID := ctx.Param("id")
	var plan models.Plan
	result := pc.DB.First(&plan, "id = ?", planID)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "Not have this ID"})
		return
	}
	var payload = struct {
		Name        string        `json:"name,omitempty"`
		Description string        `json:"description,omitempty"`
		Photo       string        `json:"photo,omitempty"`
		Type        string        `json:"type,omitempty"`
		Detail      models.Detail `gorm:"type:json" json:"detail,omitempty"`
	}{}
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}
	updatePlan := models.Plan{
		ID:          plan.ID,
		Name:        payload.Name,
		Description: payload.Description,
		Photo:       payload.Photo,
		Type:        payload.Type,
		Detail:      payload.Detail,
	}
	if err := pc.DB.Model(&plan).Updates(updatePlan).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "Can not update plan"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": "Update plan success"})
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
	var plans []models.Plan
	result := pc.DB.Find(&plans)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "not have plan data"})
		return
	}
	type Response struct {
		ID   uuid.UUID `json:"id"`
		Name string    `json:"name,omitempty"`
		Type string    `json:"type,omitempty"`
	}
	var data []Response
	for _, plan := range plans {
		response := Response{
			ID:   plan.ID,
			Name: plan.Name,
			Type: plan.Type,
		}
		data = append(data, response)
	}
	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": gin.H{"plan": data}})
}

// Delete plan
func (pc *PlanController) DeletePlan(ctx *gin.Context) {
	planID := ctx.Param("id")
	// Delete the references to the plan from the patient_plan table
	if err := pc.DB.Exec("DELETE FROM patient_plan WHERE plan_id = ?", planID).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "fail", "message": "Failed to delete references to the plan"})
		return
	}

	// find patienta that have this planID
	var patients []models.Patient
	if err := pc.DB.Where("plan_id @> ARRAY[?]::uuid[]", planID).Find(&patients).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "fail", "message": "error"})
		return
	}

	// delete planID from planID of patient
	for _, patient := range patients {
		indexToDelete := -1
		for i, pid := range patient.PlanID {
			if pid == planID {
				indexToDelete = i
				break
			}
		}
		if indexToDelete != -1 {
			patient.PlanID = append(patient.PlanID[:indexToDelete], patient.PlanID[indexToDelete+1:]...)
			if err := pc.DB.Save(&patient).Error; err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{"status": "fail", "message": "error"})
				return
			}
		}
	}

	// Now you can safely delete the plan
	result := pc.DB.Delete(&models.Plan{}, "id = ?", planID)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "No plan with that ID exists"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success"})

}
