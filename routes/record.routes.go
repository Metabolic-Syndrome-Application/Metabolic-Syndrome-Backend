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
	router.POST("/health", middleware.DeserializeUser(), rc.recordController.RecordHealth)
	router.POST("/health/:id", middleware.DeserializeUser(), rc.recordController.OtherRecordHealth)

}
