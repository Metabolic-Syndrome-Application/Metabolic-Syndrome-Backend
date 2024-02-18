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

	// quiz

	//Web
	router.POST("/quiz", middleware.DeserializeUser(), cc.challengeController.CreateQuizChallenge)
	router.PUT("/quiz/:id", middleware.DeserializeUser(), cc.challengeController.UpdateQuizChallenge)
	router.GET("/quiz/:id", middleware.DeserializeUser(), cc.challengeController.GetQuizChallenge)
	router.DELETE("/quiz/:id", middleware.DeserializeUser(), cc.challengeController.DeleteQuizChallenge)
	router.GET("/quiz/all", middleware.DeserializeUser(), cc.challengeController.GetAllQuizChallenge)

	//Mobile
	router.GET("/quiz/check-today", middleware.DeserializeUser(), cc.challengeController.CheckQuizToday)
	router.GET("/quiz/random", middleware.DeserializeUser(), cc.challengeController.GetRandomQuiz)
	router.GET("/quiz/answer", middleware.DeserializeUser(), cc.challengeController.GetPointQuiz)

	// daily

	//Web
	router.POST("/daily", middleware.DeserializeUser(), cc.challengeController.CreateDailyChallenge)
	router.PUT("/daily/:id", middleware.DeserializeUser(), cc.challengeController.UpdateDailyChallenge)
	// router.GET("/daily/:id", middleware.DeserializeUser(), cc.challengeController)
	// router.DELETE("/daily/:id", middleware.DeserializeUser(), cc.challengeController)
	// router.GET("/daily/all", middleware.DeserializeUser(), cc.challengeController)
}
