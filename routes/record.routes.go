package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/ployns/Metabolic-Syndrome-Backend/controllers"
	"github.com/ployns/Metabolic-Syndrome-Backend/middleware"
)

type RecordRouteController struct {
	recordController controllers.RecordController
}

func NewRecordRouteController(recordController controllers.RecordController) RecordRouteController {
	return RecordRouteController{recordController}
}

func (rc *RecordRouteController) RecordRoute(rg *gin.RouterGroup) {

	router := rg.Group("record")

	// Create record health (Mobile)
	router.POST("/health", middleware.DeserializeUser(), rc.recordController.RecordHealth)
	// Create record health by hospital (Web)
	router.POST("/health/:id", middleware.DeserializeUser(), rc.recordController.OtherRecordHealth)
	// Get other record health (Web)  รวมทั้ง record_by patient / hospital
	router.GET("/health/:id", middleware.DeserializeUser(), rc.recordController.GetOtherRecordHealth)
	// Get record health record_by patient (Mobile)
	router.GET("/health/patient", middleware.DeserializeUser(), rc.recordController.GetRecordHealthByPatient)
	// Get other record health record_by patient (Web)
	router.GET("/health/patient/:id", middleware.DeserializeUser(), rc.recordController.GetOtherRecordHealthByPatient)
	// Get other record health record_by hospital (Web)
	router.GET("/health/hospital/:id", middleware.DeserializeUser(), rc.recordController.GetOtherRecordHealthByHospital)
	// Get record health record_by patient bmi (Mobile)
	router.GET("/health/patient/bmi", middleware.DeserializeUser(), rc.recordController.GetRecordHealthByPatientBmi)
	// Get record health record_by patient waistline (Mobile)
	router.GET("/health/patient/waistline", middleware.DeserializeUser(), rc.recordController.GetRecordHealthByPatientWaistline)
	// Get record health record_by patient bloodGlucose (Mobile)
	router.GET("/health/patient/bloodGlucose", middleware.DeserializeUser(), rc.recordController.GetRecordHealthByPatientBloodGlucose)
	// Get record health record_by patient bloodLipids (Mobile)
	router.GET("/health/patient/bloodLipids", middleware.DeserializeUser(), rc.recordController.GetRecordHealthByPatientBloodLipids)
	// Get record health record_by patient bloodPressure (Mobile)
	router.GET("/health/patient/bloodPressure", middleware.DeserializeUser(), rc.recordController.GetRecordHealthByPatientBloodPressure)
	// Get other record health record_by patient bmi (Web)
	router.GET("/health/patient/bmi/:id", middleware.DeserializeUser(), rc.recordController.GetOtherRecordHealthByPatientBmi)
	// Get other record health record_by patient waistline (Web)
	router.GET("/health/patient/waistline/:id", middleware.DeserializeUser(), rc.recordController.GetOtherRecordHealthByPatientWaistline)
	// Get other record health record_by patient bloodGlucose (Web)
	router.GET("/health/patient/bloodGlucose/:id", middleware.DeserializeUser(), rc.recordController.GetOtherRecordHealthByPatientBloodGlucose)
	// Get other record health record_by patient bloodLipids (Web)
	router.GET("/health/patient/bloodLipids/:id", middleware.DeserializeUser(), rc.recordController.GetOtherRecordHealthByPatientBloodLipids)
	// Get other record health record_by patient bloodPressure (Web)
	router.GET("/health/patient/bloodPressure/:id", middleware.DeserializeUser(), rc.recordController.GetOtherRecordHealthByPatientBloodPressure)

}
