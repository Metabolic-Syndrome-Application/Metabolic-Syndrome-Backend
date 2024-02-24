package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/ployns/Metabolic-Syndrome-Backend/controllers"
	"github.com/ployns/Metabolic-Syndrome-Backend/middleware"
)

type ConnectRouteController struct {
	connectController controllers.ConnectController
}

func NewConnectRouteController(connectController controllers.ConnectController) ConnectRouteController {
	return ConnectRouteController{connectController}
}

func (cc *ConnectRouteController) ConnectRoute(rg *gin.RouterGroup) {
	router := rg.Group("/connect")

	// Web
	router.POST("/generate-otp", middleware.DeserializeUser(), cc.connectController.GenerateOTP)
	router.GET("/refresh-otp/:id", middleware.DeserializeUser(), cc.connectController.RefreshOTP)

}
