package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/ployns/Metabolic-Syndrome-Backend/controllers"
	"github.com/ployns/Metabolic-Syndrome-Backend/middleware"
)

type ChallengeRouteController struct {
	challengeController controllers.ChallengeController
}

func NewChallengeRouteController(challengeController controllers.ChallengeController) ChallengeRouteController {
	return ChallengeRouteController{challengeController}
}

func (cc *ChallengeRouteController) ChallengeRoute(rg *gin.RouterGroup) {

	router := rg.Group("challenge")
	router.POST("/quiz", middleware.DeserializeUser(), cc.challengeController.CreateQuizChallenge)

}
