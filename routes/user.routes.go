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
	router.GET("/profile", middleware.DeserializeUser(), uc.userController.GetProfile)    // get profile me
	router.PUT("/profile", middleware.DeserializeUser(), uc.userController.UpdateProfile) // post profile me

	router.GET("/profile/all", middleware.DeserializeUser(), uc.userController.GetAllUserProfile)              // get profile all (depend on currentUser.Role)
	router.GET("/profile/all/doctor", middleware.DeserializeUser(), uc.userController.GetAllUserProfileDoctor) // get profile all doctor

	router.PUT("/profile/:role/:id", middleware.DeserializeUser(), uc.userController.UpdateOtherProfile) // post profile other profile
	router.GET("/profile/:role/:id", middleware.DeserializeUser(), uc.userController.GetOtherProfile)    // get profile other profile
	router.DELETE("/profile/:role/:id", middleware.DeserializeUser(), uc.userController.DeleteUser)      // delete user

}
