package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/ployns/Metabolic-Syndrome-Backend/controllers"
	"github.com/ployns/Metabolic-Syndrome-Backend/middleware"
)

type RankRouteController struct {
	rankController controllers.RankController
}

func NewRankRouteController(rankController controllers.RankController) RankRouteController {
	return RankRouteController{rankController}
}

func (rc *RankRouteController) RankRoute(rg *gin.RouterGroup) {

	router := rg.Group("rank")
	router.GET("/top5", middleware.DeserializeUser(), rc.rankController.Top5)
	router.GET("/me", middleware.DeserializeUser(), rc.rankController.MyRank)

}
