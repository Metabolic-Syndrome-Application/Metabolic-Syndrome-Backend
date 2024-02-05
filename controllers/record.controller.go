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

// health
func (rc *RecordController) RecordHealth(ctx *gin.Context) {
	currentUser := ctx.MustGet("currentUser").(models.User)
	var ProfilePatient models.Patient
	result := rc.DB.First(&ProfilePatient, "id = ?", currentUser.ID)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "No post with that title exists"})
		return
	}
	var payload = struct {
		Height                 float32 `json:"height"`
		Weight                 float32 `json:"weight"`
		Waistline              float32 `json:"waistline"`
		SystolicBloodPressure  int     `json:"systolicBloodPressure"`
		DiastolicBloodPressure int     `json:"diastolicBloodPressure"`
		PulseRate              int     `json:"pulseRate"`
		BloodGlucose           float32 `json:"bloodGlucose"`
		Triglyceride           float32 `json:"triglyceride"`
		HDL                    float32 `json:"hdl"`
		LDL                    float32 `json:"ldl"`
		Cholesterol            float32 `json:"cholesterol"`
		RecordBy               string  `json:"recordBy"`
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
	result1 := rc.DB.Create(&newRecordHealth)
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
		Height                 float32 `json:"height"`
		Weight                 float32 `json:"weight"`
		Waistline              float32 `json:"waistline"`
		SystolicBloodPressure  int     `json:"systolicBloodPressure"`
		DiastolicBloodPressure int     `json:"diastolicBloodPressure"`
		PulseRate              int     `json:"pulseRate"`
		BloodGlucose           float32 `json:"bloodGlucose"`
		Triglyceride           float32 `json:"triglyceride"`
		HDL                    float32 `json:"hdl"`
		LDL                    float32 `json:"ldl"`
		Cholesterol            float32 `json:"cholesterol"`
		RecordBy               string  `json:"recordBy"`
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
	result1 := rc.DB.Create(&newRecordHealth)
	if result1.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "Can not create RecordHealth"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success"})

}

func (rc *RecordController) GetOtherRecordHealth(ctx *gin.Context) {
	userID := ctx.Param("id")
	var records []models.RecordHealth
	result := rc.DB.Where("patient_id = ?", userID).Find(&records)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "Not have this ID"})
		return
	}
	type Response struct {
		Height                 float32 `json:"height"`
		Weight                 float32 `json:"weight"`
		BMI                    float32 `json:"bmi"`
		Waistline              float32 `json:"waistline"`
		SystolicBloodPressure  int     `json:"systolicBloodPressure"`
		DiastolicBloodPressure int     `json:"diastolicBloodPressure"`
		PulseRate              int     `json:"pulseRate"`
		BloodGlucose           float32 `json:"bloodGlucose"`
		Cholesterol            float32 `json:"cholesterol"`
		HDL                    float32 `json:"hdl"`
		LDL                    float32 `json:"ldl"`
		Triglyceride           float32 `json:"triglyceride"`
		RecordBy               string  `json:"recordBy"`
		Timestamp              string
	}
	var data []Response
	for _, record := range records {
		bmi := record.Weight / ((float32(record.Height) / 100) * (float32(record.Height) / 100))
		response := Response{
			Height:                 record.Height,
			Weight:                 record.Weight,
			BMI:                    bmi,
			Waistline:              record.Waistline,
			SystolicBloodPressure:  record.SystolicBloodPressure,
			DiastolicBloodPressure: record.DiastolicBloodPressure,
			PulseRate:              record.PulseRate,
			BloodGlucose:           record.BloodGlucose,
			Cholesterol:            record.Cholesterol,
			HDL:                    record.HDL,
			LDL:                    record.LDL,
			Triglyceride:           record.Triglyceride,
			RecordBy:               record.RecordBy,
			Timestamp:              record.Timestamp.Format("2006-01-02 15:04:05"),
		}
		data = append(data, response)
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": gin.H{"record": data}})

}

func (rc *RecordController) GetRecordHealthByPatient(ctx *gin.Context) {
	currentUser := ctx.MustGet("currentUser").(models.User)
	var records []models.RecordHealth
	result := rc.DB.Where("patient_id = ? AND record_by = 'patient'", currentUser.ID).Find(&records)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "Not have this ID"})
		return
	}
	type Response struct {
		Height                 float32 `json:"height"`
		Weight                 float32 `json:"weight"`
		BMI                    float32 `json:"bmi"`
		Waistline              float32 `json:"waistline"`
		SystolicBloodPressure  int     `json:"systolicBloodPressure"`
		DiastolicBloodPressure int     `json:"diastolicBloodPressure"`
		PulseRate              int     `json:"pulseRate"`
		BloodGlucose           float32 `json:"bloodGlucose"`
		Cholesterol            float32 `json:"cholesterol"`
		HDL                    float32 `json:"hdl"`
		LDL                    float32 `json:"ldl"`
		Triglyceride           float32 `json:"triglyceride"`
		RecordBy               string  `json:"recordBy"`
		Timestamp              string
	}
	var data []Response
	for _, record := range records {
		bmi := record.Weight / ((float32(record.Height) / 100) * (float32(record.Height) / 100))
		response := Response{
			Height:                 record.Height,
			Weight:                 record.Weight,
			BMI:                    bmi,
			Waistline:              record.Waistline,
			SystolicBloodPressure:  record.SystolicBloodPressure,
			DiastolicBloodPressure: record.DiastolicBloodPressure,
			PulseRate:              record.PulseRate,
			BloodGlucose:           record.BloodGlucose,
			Cholesterol:            record.Cholesterol,
			HDL:                    record.HDL,
			LDL:                    record.LDL,
			Triglyceride:           record.Triglyceride,
			RecordBy:               record.RecordBy,
			Timestamp:              record.Timestamp.Format("2006-01-02 15:04:05"),
		}
		data = append(data, response)
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": gin.H{"record": data}})

}

func (rc *RecordController) GetOtherRecordHealthByPatient(ctx *gin.Context) {
	userID := ctx.Param("id")
	var records []models.RecordHealth
	result := rc.DB.Where("patient_id = ? AND record_by = 'patient'", userID).Find(&records)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "Not have this ID"})
		return
	}
	type Response struct {
		Height                 float32 `json:"height"`
		Weight                 float32 `json:"weight"`
		BMI                    float32 `json:"bmi"`
		Waistline              float32 `json:"waistline"`
		SystolicBloodPressure  int     `json:"systolicBloodPressure"`
		DiastolicBloodPressure int     `json:"diastolicBloodPressure"`
		PulseRate              int     `json:"pulseRate"`
		BloodGlucose           float32 `json:"bloodGlucose"`
		Cholesterol            float32 `json:"cholesterol"`
		HDL                    float32 `json:"hdl"`
		LDL                    float32 `json:"ldl"`
		Triglyceride           float32 `json:"triglyceride"`
		RecordBy               string  `json:"recordBy"`
		Timestamp              string
	}
	var data []Response
	for _, record := range records {
		bmi := record.Weight / ((float32(record.Height) / 100) * (float32(record.Height) / 100))
		response := Response{
			Height:                 record.Height,
			Weight:                 record.Weight,
			BMI:                    bmi,
			Waistline:              record.Waistline,
			SystolicBloodPressure:  record.SystolicBloodPressure,
			DiastolicBloodPressure: record.DiastolicBloodPressure,
			PulseRate:              record.PulseRate,
			BloodGlucose:           record.BloodGlucose,
			Cholesterol:            record.Cholesterol,
			HDL:                    record.HDL,
			LDL:                    record.LDL,
			Triglyceride:           record.Triglyceride,
			RecordBy:               record.RecordBy,
			Timestamp:              record.Timestamp.Format("2006-01-02 15:04:05"),
		}
		data = append(data, response)
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": gin.H{"record": data}})

}

func (rc *RecordController) GetOtherRecordHealthByHospital(ctx *gin.Context) {
	userID := ctx.Param("id")
	var records []models.RecordHealth
	result := rc.DB.Where("patient_id = ? AND (record_by = 'doctor' OR record_by = 'staff')", userID).Find(&records)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "Not have this ID"})
		return
	}
	type Response struct {
		Height                 float32 `json:"height"`
		Weight                 float32 `json:"weight"`
		BMI                    float32 `json:"bmi"`
		Waistline              float32 `json:"waistline"`
		SystolicBloodPressure  int     `json:"systolicBloodPressure"`
		DiastolicBloodPressure int     `json:"diastolicBloodPressure"`
		PulseRate              int     `json:"pulseRate"`
		BloodGlucose           float32 `json:"bloodGlucose"`
		Cholesterol            float32 `json:"cholesterol"`
		HDL                    float32 `json:"hdl"`
		LDL                    float32 `json:"ldl"`
		Triglyceride           float32 `json:"triglyceride"`
		RecordBy               string  `json:"recordBy"`
		Timestamp              string
	}
	var data []Response
	for _, record := range records {
		bmi := record.Weight / ((float32(record.Height) / 100) * (float32(record.Height) / 100))
		response := Response{
			Height:                 record.Height,
			Weight:                 record.Weight,
			BMI:                    bmi,
			Waistline:              record.Waistline,
			SystolicBloodPressure:  record.SystolicBloodPressure,
			DiastolicBloodPressure: record.DiastolicBloodPressure,
			PulseRate:              record.PulseRate,
			BloodGlucose:           record.BloodGlucose,
			Cholesterol:            record.Cholesterol,
			HDL:                    record.HDL,
			LDL:                    record.LDL,
			Triglyceride:           record.Triglyceride,
			RecordBy:               record.RecordBy,
			Timestamp:              record.Timestamp.Format("2006-01-02 15:04:05"),
		}
		data = append(data, response)
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": gin.H{"record": data}})

}

func (rc *RecordController) GetRecordHealthByPatientBmi(ctx *gin.Context) {
	currentUser := ctx.MustGet("currentUser").(models.User)
	var records []models.RecordHealth
	result := rc.DB.Where("patient_id = ? AND record_by = 'patient'", currentUser.ID).Find(&records)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "Not have this ID"})
		return
	}
	type Response struct {
		BMI       float32 `json:"bmi"`
		RecordBy  string  `json:"recordBy"`
		Timestamp string
	}
	var data []Response
	for _, record := range records {
		bmi := record.Weight / ((float32(record.Height) / 100) * (float32(record.Height) / 100))
		response := Response{
			BMI:       bmi,
			RecordBy:  record.RecordBy,
			Timestamp: record.Timestamp.Format("2006-01-02 15:04:05"),
		}
		data = append(data, response)
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": gin.H{"record": data}})

}

func (rc *RecordController) GetRecordHealthByPatientWaistline(ctx *gin.Context) {
	currentUser := ctx.MustGet("currentUser").(models.User)
	var records []models.RecordHealth
	result := rc.DB.Where("patient_id = ? AND record_by = 'patient'", currentUser.ID).Find(&records)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "Not have this ID"})
		return
	}
	type Response struct {
		Waistline float32 `json:"waistline"`
		RecordBy  string  `json:"recordBy"`
		Timestamp string
	}
	var data []Response
	for _, record := range records {
		response := Response{
			Waistline: record.Waistline,
			RecordBy:  record.RecordBy,
			Timestamp: record.Timestamp.Format("2006-01-02 15:04:05"),
		}
		data = append(data, response)
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": gin.H{"record": data}})

}

func (rc *RecordController) GetRecordHealthByPatientBloodGlucose(ctx *gin.Context) {
	currentUser := ctx.MustGet("currentUser").(models.User)
	var records []models.RecordHealth
	result := rc.DB.Where("patient_id = ? AND record_by = 'patient'", currentUser.ID).Find(&records)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "Not have this ID"})
		return
	}
	type Response struct {
		BloodGlucose float32 `json:"bloodGlucose"`
		RecordBy     string  `json:"recordBy"`
		Timestamp    string
	}
	var data []Response
	for _, record := range records {
		response := Response{
			BloodGlucose: record.BloodGlucose,
			RecordBy:     record.RecordBy,
			Timestamp:    record.Timestamp.Format("2006-01-02 15:04:05"),
		}
		data = append(data, response)
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": gin.H{"record": data}})

}

func (rc *RecordController) GetRecordHealthByPatientBloodLipids(ctx *gin.Context) {
	currentUser := ctx.MustGet("currentUser").(models.User)
	var records []models.RecordHealth
	result := rc.DB.Where("patient_id = ? AND record_by = 'patient'", currentUser.ID).Find(&records)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "Not have this ID"})
		return
	}
	type Response struct {
		Cholesterol  float32 `json:"cholesterol"`
		HDL          float32 `json:"hdl"`
		LDL          float32 `json:"ldl"`
		Triglyceride float32 `json:"triglyceride"`
		RecordBy     string  `json:"recordBy"`
		Timestamp    string
	}
	var data []Response
	for _, record := range records {
		response := Response{
			Cholesterol:  record.Cholesterol,
			HDL:          record.HDL,
			LDL:          record.LDL,
			Triglyceride: record.Triglyceride,
			RecordBy:     record.RecordBy,
			Timestamp:    record.Timestamp.Format("2006-01-02 15:04:05"),
		}
		data = append(data, response)
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": gin.H{"record": data}})

}

func (rc *RecordController) GetRecordHealthByPatientBloodPressure(ctx *gin.Context) {
	currentUser := ctx.MustGet("currentUser").(models.User)
	var records []models.RecordHealth
	result := rc.DB.Where("patient_id = ? AND record_by = 'patient'", currentUser.ID).Find(&records)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "Not have this ID"})
		return
	}
	type Response struct {
		SystolicBloodPressure  int    `json:"systolicBloodPressure"`
		DiastolicBloodPressure int    `json:"diastolicBloodPressure"`
		PulseRate              int    `json:"pulseRate"`
		RecordBy               string `json:"recordBy"`
		Timestamp              string
	}
	var data []Response
	for _, record := range records {
		response := Response{
			SystolicBloodPressure:  record.SystolicBloodPressure,
			DiastolicBloodPressure: record.DiastolicBloodPressure,
			PulseRate:              record.PulseRate,
			RecordBy:               record.RecordBy,
			Timestamp:              record.Timestamp.Format("2006-01-02 15:04:05"),
		}
		data = append(data, response)
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": gin.H{"record": data}})

}

func (rc *RecordController) GetOtherRecordHealthByPatientBmi(ctx *gin.Context) {
	userID := ctx.Param("id")
	var records []models.RecordHealth
	result := rc.DB.Where("patient_id = ? AND record_by = 'patient'", userID).Find(&records)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "Not have this ID"})
		return
	}
	type Response struct {
		BMI       float32 `json:"bmi"`
		RecordBy  string  `json:"recordBy"`
		Timestamp string
	}
	var data []Response
	for _, record := range records {
		bmi := record.Weight / ((float32(record.Height) / 100) * (float32(record.Height) / 100))
		response := Response{
			BMI:       bmi,
			RecordBy:  record.RecordBy,
			Timestamp: record.Timestamp.Format("2006-01-02 15:04:05"),
		}
		data = append(data, response)
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": gin.H{"record": data}})

}

func (rc *RecordController) GetOtherRecordHealthByPatientWaistline(ctx *gin.Context) {
	userID := ctx.Param("id")
	var records []models.RecordHealth
	result := rc.DB.Where("patient_id = ? AND record_by = 'patient'", userID).Find(&records)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "Not have this ID"})
		return
	}
	type Response struct {
		Waistline float32 `json:"waistline"`
		RecordBy  string  `json:"recordBy"`
		Timestamp string
	}
	var data []Response
	for _, record := range records {
		response := Response{
			Waistline: record.Waistline,
			RecordBy:  record.RecordBy,
			Timestamp: record.Timestamp.Format("2006-01-02 15:04:05"),
		}
		data = append(data, response)
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": gin.H{"record": data}})

}

func (rc *RecordController) GetOtherRecordHealthByPatientBloodGlucose(ctx *gin.Context) {
	userID := ctx.Param("id")
	var records []models.RecordHealth
	result := rc.DB.Where("patient_id = ? AND record_by = 'patient'", userID).Find(&records)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "Not have this ID"})
		return
	}
	type Response struct {
		BloodGlucose float32 `json:"bloodGlucose"`
		RecordBy     string  `json:"recordBy"`
		Timestamp    string
	}
	var data []Response
	for _, record := range records {
		response := Response{
			BloodGlucose: record.BloodGlucose,
			RecordBy:     record.RecordBy,
			Timestamp:    record.Timestamp.Format("2006-01-02 15:04:05"),
		}
		data = append(data, response)
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": gin.H{"record": data}})

}

func (rc *RecordController) GetOtherRecordHealthByPatientBloodLipids(ctx *gin.Context) {
	userID := ctx.Param("id")
	var records []models.RecordHealth
	result := rc.DB.Where("patient_id = ? AND record_by = 'patient'", userID).Find(&records)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "Not have this ID"})
		return
	}
	type Response struct {
		Cholesterol  float32 `json:"cholesterol"`
		HDL          float32 `json:"hdl"`
		LDL          float32 `json:"ldl"`
		Triglyceride float32 `json:"triglyceride"`
		RecordBy     string  `json:"recordBy"`
		Timestamp    string
	}
	var data []Response
	for _, record := range records {
		response := Response{
			Cholesterol:  record.Cholesterol,
			HDL:          record.HDL,
			LDL:          record.LDL,
			Triglyceride: record.Triglyceride,
			RecordBy:     record.RecordBy,
			Timestamp:    record.Timestamp.Format("2006-01-02 15:04:05"),
		}
		data = append(data, response)
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": gin.H{"record": data}})

}

func (rc *RecordController) GetOtherRecordHealthByPatientBloodPressure(ctx *gin.Context) {
	userID := ctx.Param("id")
	var records []models.RecordHealth
	result := rc.DB.Where("patient_id = ? AND record_by = 'patient'", userID).Find(&records)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "Not have this ID"})
		return
	}
	type Response struct {
		SystolicBloodPressure  int    `json:"systolicBloodPressure"`
		DiastolicBloodPressure int    `json:"diastolicBloodPressure"`
		PulseRate              int    `json:"pulseRate"`
		RecordBy               string `json:"recordBy"`
		Timestamp              string
	}
	var data []Response
	for _, record := range records {
		response := Response{
			SystolicBloodPressure:  record.SystolicBloodPressure,
			DiastolicBloodPressure: record.DiastolicBloodPressure,
			PulseRate:              record.PulseRate,
			RecordBy:               record.RecordBy,
			Timestamp:              record.Timestamp.Format("2006-01-02 15:04:05"),
		}
		data = append(data, response)
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": gin.H{"record": data}})
}
