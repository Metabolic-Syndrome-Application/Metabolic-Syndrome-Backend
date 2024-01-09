package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/ployns/Metabolic-Syndrome-Backend/controllers"
	"github.com/ployns/Metabolic-Syndrome-Backend/middleware"
)

type PlanRouteController struct {
	planController controllers.PlanController
}

func NewPlanRouteController(planController controllers.PlanController) PlanRouteController {
	return PlanRouteController{planController}
}

func (pc *PlanRouteController) PlanRoute(rg *gin.RouterGroup) {

	router := rg.Group("plan")
	router.POST("/create", middleware.DeserializeUser(), pc.planController.CreatePlan)

}
