package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/ployns/Metabolic-Syndrome-Backend/controllers"
	"github.com/ployns/Metabolic-Syndrome-Backend/middleware"
)

type ScreeningRouteController struct {
	screeningController controllers.ScreeningController
}

func NewScreeningRouteController(screeningController controllers.ScreeningController) ScreeningRouteController {
	return ScreeningRouteController{screeningController}
}

func (sc *ScreeningRouteController) ScreeningRoute(rg *gin.RouterGroup) {

	router := rg.Group("screening")
	router.POST("/metabolic", middleware.DeserializeUser(), sc.screeningController.MetabolicRisk)
	router.POST("/disease", middleware.DeserializeUser(), sc.screeningController.DiseaseRisk)
	router.GET("/disease", middleware.DeserializeUser(), sc.screeningController.GetDiseaseRisk)

}
