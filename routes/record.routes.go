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

	//create record health by patient
	router.POST("/health", middleware.DeserializeUser(), rc.recordController.RecordHealth)
	//create record health by hospital
	router.POST("/health/:id", middleware.DeserializeUser(), rc.recordController.OtherRecordHealth)
	// get other record health (record by patient / hospital)
	router.GET("/health/:id", middleware.DeserializeUser(), rc.recordController.GetOtherRecordHealth)
	// get my record health (record by patient)
	router.GET("/health/patient", middleware.DeserializeUser(), rc.recordController.GetRecordHealthByPatient)

	// // get other record health (record by patient)
	// router.GET("/health/patient/:id", middleware.DeserializeUser(), rc.recordController.GetRecordHealthByPatient)
	// // get other record health (record by hospital)
	// router.GET("/health/hospital/:id", middleware.DeserializeUser(), rc.recordController.GetRecordHealthByPatient)

}
