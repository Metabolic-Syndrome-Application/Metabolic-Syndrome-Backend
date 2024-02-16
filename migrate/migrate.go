package main

import (
	"fmt"
	"log"

	"github.com/ployns/Metabolic-Syndrome-Backend/initializers"
	"github.com/ployns/Metabolic-Syndrome-Backend/models"
)

func init() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("? Could not load environment variables", err)
	}

	initializers.ConnectDB(&config)
}

func main() {
	initializers.DB.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"")
	initializers.DB.AutoMigrate(&models.User{}, &models.Patient{}, &models.Doctor{}, &models.Staff{},
		&models.Plan{}, &models.DailyChallenge{}, &models.RecordHealth{}, &models.RecordPlan{}, &models.QuizChallenge{})
	fmt.Println("? Migration complete")
}
