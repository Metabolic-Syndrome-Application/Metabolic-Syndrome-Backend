package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/ployns/Metabolic-Syndrome-Backend/controllers"
	"github.com/ployns/Metabolic-Syndrome-Backend/middleware"
)

type KnowledgeRouteController struct {
	knowledgeController controllers.KnowledgeController
}

func NewKnowledgeRouteController(knowledgeController controllers.KnowledgeController) KnowledgeRouteController {
	return KnowledgeRouteController{knowledgeController}
}

func (kc *KnowledgeRouteController) KnowledgeRoute(rg *gin.RouterGroup) {
	router := rg.Group("/knowledge")

	// Mobile
	router.GET("/:disease", middleware.DeserializeUser(), kc.knowledgeController.GetKnowledge)

}
