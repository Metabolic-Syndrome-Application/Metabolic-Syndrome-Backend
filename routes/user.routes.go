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

	router.GET("/profile/all", middleware.DeserializeUser(), uc.userController.GetAllUserProfile) // get profile all (depend on currentUser.Role)

	router.POST("/profile/:role/:id", uc.userController.PostOtherProfile) // post profile other profile
	router.GET("/profile/:role/:id", uc.userController.GetOtherProfile)   // get profile other profile
	router.DELETE("/profile/:role/:id", uc.userController.DeleteUser)     // delete user

}
