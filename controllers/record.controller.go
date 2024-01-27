package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ployns/Metabolic-Syndrome-Backend/models"
	"gorm.io/gorm"
)

type RecordController struct {
	DB *gorm.DB
}

func NewRecordController(DB *gorm.DB) RecordController {
	return RecordController{DB}
}

func (rc *RecordController) RecordHealth(ctx *gin.Context) {
	currentUser := ctx.MustGet("currentUser").(models.User)
	var ProfilePatient models.Patient
	result := rc.DB.First(&ProfilePatient, "id = ?", currentUser.ID)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "No post with that title exists"})
		return
	}
	var payload = struct {
		Height                 float32 `json:"height,omitempty"`
		Weight                 float32 `json:"weight,omitempty"`
		Waistline              float32 `json:"waistline,omitempty"`
		SystolicBloodPressure  int     `json:"systolicBloodPressure,omitempty"`
		DiastolicBloodPressure int     `json:"diastolicBloodPressure,omitempty"`
		PulseRate              int     `json:"pulseRate,omitempty"`
		BloodGlucose           float32 `json:"bloodGlucose,omitempty"`
		Triglyceride           float32 `json:"triglyceride,omitempty"`
		HDL                    float32 `json:"hdl,omitempty"`
		LDL                    float32 `json:"ldl,omitempty"`
		Cholesterol            float32 `json:"cholesterol,omitempty"`
		RecordBy               string  `json:"recordBy,omitempty"`
	}{}
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}
	now := time.Now()
	newRecordHealth := &models.RecordHealth{
		PatientID:              ProfilePatient.ID,
		Height:                 payload.Height,
		Weight:                 payload.Weight,
		Waistline:              payload.Waistline,
		SystolicBloodPressure:  payload.SystolicBloodPressure,
		DiastolicBloodPressure: payload.DiastolicBloodPressure,
		PulseRate:              payload.PulseRate,
		BloodGlucose:           payload.BloodGlucose,
		Triglyceride:           payload.Triglyceride,
		HDL:                    payload.HDL,
		LDL:                    payload.LDL,
		Cholesterol:            payload.Cholesterol,
		RecordBy:               currentUser.Role,
		Timestamp:              now,
	}
	result1 := rc.DB.Save(&newRecordHealth)
	if result1.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "Can not create RecordHealth"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success"})

}

func (rc *RecordController) OtherRecordHealth(ctx *gin.Context) {
	currentUser := ctx.MustGet("currentUser").(models.User)
	userID := ctx.Param("id")
	var patient models.Patient
	result := rc.DB.First(&patient, "id = ?", userID)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "No post with that title exists"})
		return
	}
	var payload = struct {
		Height                 float32 `json:"height,omitempty"`
		Weight                 float32 `json:"weight,omitempty"`
		Waistline              float32 `json:"waistline,omitempty"`
		SystolicBloodPressure  int     `json:"systolicBloodPressure,omitempty"`
		DiastolicBloodPressure int     `json:"diastolicBloodPressure,omitempty"`
		PulseRate              int     `json:"pulseRate,omitempty"`
		BloodGlucose           float32 `json:"bloodGlucose,omitempty"`
		Triglyceride           float32 `json:"triglyceride,omitempty"`
		HDL                    float32 `json:"hdl,omitempty"`
		LDL                    float32 `json:"ldl,omitempty"`
		Cholesterol            float32 `json:"cholesterol,omitempty"`
		RecordBy               string  `json:"recordBy,omitempty"`
	}{}
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}
	now := time.Now()
	newRecordHealth := &models.RecordHealth{
		PatientID:              patient.ID,
		Height:                 payload.Height,
		Weight:                 payload.Weight,
		Waistline:              payload.Waistline,
		SystolicBloodPressure:  payload.SystolicBloodPressure,
		DiastolicBloodPressure: payload.DiastolicBloodPressure,
		PulseRate:              payload.PulseRate,
		BloodGlucose:           payload.BloodGlucose,
		Triglyceride:           payload.Triglyceride,
		HDL:                    payload.HDL,
		LDL:                    payload.LDL,
		Cholesterol:            payload.Cholesterol,
		RecordBy:               currentUser.Role,
		Timestamp:              now,
	}
	result1 := rc.DB.Save(&newRecordHealth)
	if result1.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "Can not create RecordHealth"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success"})

}
