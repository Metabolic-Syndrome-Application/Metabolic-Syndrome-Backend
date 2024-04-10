package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/ployns/Metabolic-Syndrome-Backend/controllers"
	"github.com/ployns/Metabolic-Syndrome-Backend/middleware"
)

type AuthRouteController struct {
	authController controllers.AuthController
}

func NewAuthRouteController(authController controllers.AuthController) AuthRouteController {
	return AuthRouteController{authController}
}

func (ac *AuthRouteController) AuthRoute(rg *gin.RouterGroup) {
	router := rg.Group("/auth")

	router.POST("/register", ac.authController.SignUpUser)
	router.POST("/register/other", middleware.DeserializeUser(), ac.authController.SignUpUser)
	router.POST("/login", ac.authController.SignInUser)
	router.POST("/refresh", ac.authController.RefreshAccessToken)
	router.GET("/logout", middleware.DeserializeUser(), ac.authController.LogoutUser)
	router.POST("/change-password", middleware.DeserializeUser(), ac.authController.ChangePassword)

}
