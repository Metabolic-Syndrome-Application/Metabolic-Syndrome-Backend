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

	// health

	// Create record health (Mobile)
	router.POST("/health", middleware.DeserializeUser(), rc.recordController.RecordHealth)
	// Create record health by hospital (Web)
	router.POST("/health/:id", middleware.DeserializeUser(), rc.recordController.OtherRecordHealth)
	// Get other record health (Web)  รวมทั้ง record_by patient / hospital
	router.GET("/health/:id", middleware.DeserializeUser(), rc.recordController.GetOtherRecordHealth)
	// Get record health record_by hospital latest(Web)
	router.GET("/health/hospital/latest/:id", middleware.DeserializeUser(), rc.recordController.GetRecordHealthByHospitalLatest)
	// Get record health record_by patient latest(Mobile)
	router.GET("/health/patient/latest", middleware.DeserializeUser(), rc.recordController.GetRecordHealthByPatientLatest)
	// Get other record health record_by patient (Web)
	router.GET("/health/patient/:id", middleware.DeserializeUser(), rc.recordController.GetOtherRecordHealthByPatient)
	// Get other record health record_by hospital (Web)
	router.GET("/health/hospital/:id", middleware.DeserializeUser(), rc.recordController.GetOtherRecordHealthByHospital)
	// Get record health record_by patient type (Mobile)
	router.GET("/health/patient/me/:type", middleware.DeserializeUser(), rc.recordController.GetRecordHealthByPatientType)
	// Get other record health record_by patient type (Web)
	router.GET("/health/patient/:id/:type", middleware.DeserializeUser(), rc.recordController.GetOtherRecordHealthByPatientType)

	/////////////////////////////////////////////////////////////////////////////////////////////

	// plan

	// today
	// get record plan (Mobile)
	router.GET("/plan", middleware.DeserializeUser(), rc.recordController.GetRecordPlan)
	// get record plan list (Mobile)
	router.GET("/plan/list", middleware.DeserializeUser(), rc.recordController.GetRecordPlanList)
	// get record plan mood (Mobile)
	router.GET("/plan/mood", middleware.DeserializeUser(), rc.recordController.GetRecordPlanMood)
	// get record plan mood all (Mobile)
	router.GET("/plan/mood/all", middleware.DeserializeUser(), rc.recordController.GetRecordPlanMoodAll)

	// update record plan list (Mobile)
	router.PUT("/plan/list", middleware.DeserializeUser(), rc.recordController.UpdateRecordPlanList)
	// update record plan mood (Mobile)
	router.PUT("/plan/mood", middleware.DeserializeUser(), rc.recordController.UpdateRecordPlanMood)

	// by date (not today)
	// get record plan list (Mobile)
	router.GET("/plan/list/:date", middleware.DeserializeUser(), rc.recordController.GetRecordPlanListByDate)
	// update record plan list (Mobile)
	router.PUT("/plan/list/:date", middleware.DeserializeUser(), rc.recordController.UpdateRecordPlanListByDate)

	// get record plan (web)
	router.GET("/plan/:id", middleware.DeserializeUser(), rc.recordController.GetOtherRecordPlan)

	/////////////////////////////////////////////////////////////////////////////////////////////

	// daily

	// mobile
	router.GET("/daily/list", middleware.DeserializeUser(), rc.recordController.GetRecordDailyList)
	router.PUT("/daily/list", middleware.DeserializeUser(), rc.recordController.UpdateRecordDailyList)

}
