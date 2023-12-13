package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/ployns/Metabolic-Syndrome-Backend/controllers"
	"github.com/ployns/Metabolic-Syndrome-Backend/middleware"
)

type UserRouteController struct {
	userController controllers.UserController
}

func NewUserRouteController(userController controllers.UserController) UserRouteController {
	return UserRouteController{userController}
}

func (uc *UserRouteController) UserRoute(rg *gin.RouterGroup) {

	router := rg.Group("user")
	router.GET("/profile", middleware.DeserializeUser(), uc.userController.GetProfile)     // get profile me
	router.POST("/profile", middleware.DeserializeUser(), uc.userController.UpdateProfile) // post profile me

	router.GET("/profile/:role", uc.userController.GetAllUserProfile) // get profile with role
	// router.GET("/profile/:role/:id", uc.userController.GetOtherProfile) // get profile other profile
}
