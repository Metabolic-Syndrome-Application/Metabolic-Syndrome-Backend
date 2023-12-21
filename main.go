package main

import (
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/ployns/Metabolic-Syndrome-Backend/controllers"
	"github.com/ployns/Metabolic-Syndrome-Backend/initializers"
	"github.com/ployns/Metabolic-Syndrome-Backend/routes"
)

var (
	server                   *gin.Engine
	AuthController           controllers.AuthController
	AuthRouteController      routes.AuthRouteController
	UserController           controllers.UserController
	UserRouteController      routes.UserRouteController
	ScreeningController      controllers.ScreeningController
	ScreeningRouteController routes.ScreeningRouteController
)

func init() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("? Could not load environment variables", err)
	}

	initializers.ConnectDB(&config)

	AuthController = controllers.NewAuthController(initializers.DB)
	AuthRouteController = routes.NewAuthRouteController(AuthController)

	UserController = controllers.NewUserController(initializers.DB)
	UserRouteController = routes.NewUserRouteController(UserController)

	ScreeningController = controllers.NewScreeningController(initializers.DB)
	ScreeningRouteController = routes.NewScreeningRouteController(ScreeningController)

	server = gin.Default()
}

func main() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("? Could not load environment variables", err)
	}

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"http://localhost:8000", config.ClientOrigin}
	corsConfig.AllowCredentials = true
	corsConfig.AllowHeaders = []string{"Origin", "Content-Type", "Accept", "Authorization"}

	server.Use(cors.New(corsConfig))

	router := server.Group("/api")

	AuthRouteController.AuthRoute(router)
	UserRouteController.UserRoute(router)
	ScreeningRouteController.ScreeningRoute(router)
	log.Fatal(server.Run(":" + config.ServerPort))
}
