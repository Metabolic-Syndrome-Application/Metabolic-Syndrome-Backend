package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ployns/Metabolic-Syndrome-Backend/models"
	"gorm.io/gorm"
)

type ScreeningController struct {
	DB *gorm.DB
}

func NewScreeningController(DB *gorm.DB) ScreeningController {
	return ScreeningController{DB}
}

func (sc *ScreeningController) MetabolicRisk(ctx *gin.Context) {
	currentUser := ctx.MustGet("currentUser").(models.User)
	var ProfilePatient models.Patient
	result := sc.DB.First(&ProfilePatient, "id = ?", currentUser.ID)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "No post with that title exists"})
		return
	}
	var payload = struct {
		Occupation             string
		Height                 float32 `json:"height,omitempty"`
		Weight                 float32 `json:"weight,omitempty"`
		BMI                    float32 `json:"bmi,omitempty"`
		Waistline              float32 `json:"waistline,omitempty"`
		SystolicBloodPressure  int     `json:"systolicBloodPressure,omitempty"`
		DiastolicBloodPressure int     `json:"diastolicBloodPressure,omitempty"`
		PulseRate              int     `json:"pulseRate,omitempty"`
		BloodGlucose           float32 `json:"bloodGlucose,omitempty"`
		Triglyceride           float32 `json:"triglyceride,omitempty"`
		HDL                    float32 `json:"hdl,omitempty"`
	}{}
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}
	var risk = 0

	// screening Waistline and HDL
	if ProfilePatient.Gender == "female" {
		if payload.Waistline >= 80 {
			risk = risk + 1
		}
		if payload.HDL < 50 {
			risk = risk + 1
		}
	} else if ProfilePatient.Gender == "male" {
		if payload.Waistline >= 90 {
			risk = risk + 1
		}
		if payload.HDL < 40 {
			risk = risk + 1
		}
	}

	// screening BloodGlucose
	if payload.BloodGlucose > 100 {
		risk = risk + 1
	}

	// screening BloodPressure
	if payload.SystolicBloodPressure > 130 || payload.DiastolicBloodPressure > 85 {
		risk = risk + 1
	}

	// screening Triglyceride
	if payload.Triglyceride > 150 {
		risk = risk + 1
	}

	// result
	var metabolicRisk = ""
	if risk <= 1 {
		metabolicRisk = "low"

		updateRisk := &models.Patient{
			DiseaseRisk: "metabolic",
		}

		result := sc.DB.Model(&ProfilePatient).Updates(updateRisk)
		if result.Error != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "Can not update profile Staff"})
			return
		}
	} else if risk == 2 {
		metabolicRisk = "medium"
	} else if risk >= 3 {
		metabolicRisk = "high"
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": gin.H{"metabolicRisk": metabolicRisk}})

}
