package main

import (
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/ployns/Metabolic-Syndrome-Backend/controllers"
	"github.com/ployns/Metabolic-Syndrome-Backend/initializers"
	"github.com/ployns/Metabolic-Syndrome-Backend/migrate"
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
	PlanController           controllers.PlanController
	PlanRouteController      routes.PlanRouteController
	RecordController         controllers.RecordController
	RecordRouteController    routes.RecordRouteController
	ChallengeController      controllers.ChallengeController
	ChallengeRouteController routes.ChallengeRouteController
	ConnectController        controllers.ConnectController
	ConnectRouteController   routes.ConnectRouteController
	RankController           controllers.RankController
	RankRouteController      routes.RankRouteController
	KnowledgeController      controllers.KnowledgeController
	KnowledgeRouteController routes.KnowledgeRouteController
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

	PlanController = controllers.NewPlanController(initializers.DB)
	PlanRouteController = routes.NewPlanRouteController(PlanController)

	RecordController = controllers.NewRecordController(initializers.DB)
	RecordRouteController = routes.NewRecordRouteController(RecordController)

	ChallengeController = controllers.NewChallengeController(initializers.DB)
	ChallengeRouteController = routes.NewChallengeRouteController(ChallengeController)

	ConnectController = controllers.NewConnectController(initializers.DB)
	ConnectRouteController = routes.NewConnectRouteController(ConnectController)

	RankController = controllers.NewRankController(initializers.DB)
	RankRouteController = routes.NewRankRouteController(RankController)

	KnowledgeController = controllers.NewKnowledgeController(initializers.DB)
	KnowledgeRouteController = routes.NewKnowledgeRouteController(KnowledgeController)

	server = gin.Default()
}

func main() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("? Could not load environment variables", err)
	}

	migrate.Migrate()

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"http://localhost:8000", config.ClientOrigin}
	corsConfig.AllowCredentials = true
	corsConfig.AllowHeaders = []string{"Origin", "Content-Type", "Accept", "Authorization", "Cookie"}
	corsConfig.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE"}
	corsConfig.ExposeHeaders = []string{"Content-Length"}

	server.Use(cors.New(corsConfig))

	router := server.Group("/api")

	AuthRouteController.AuthRoute(router)
	UserRouteController.UserRoute(router)
	ScreeningRouteController.ScreeningRoute(router)
	PlanRouteController.PlanRoute(router)
	RecordRouteController.RecordRoute(router)
	ChallengeRouteController.ChallengeRoute(router)
	ConnectRouteController.ConnectRoute(router)
	RankRouteController.RankRoute(router)
	KnowledgeRouteController.KnowledgeRoute(router)
	log.Fatal(server.Run(":" + config.ServerPort))
}
