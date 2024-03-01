package main

import (
	"fmt"
	"log"

	"github.com/ployns/Metabolic-Syndrome-Backend/initializers"
	"github.com/ployns/Metabolic-Syndrome-Backend/models"
	"github.com/ployns/Metabolic-Syndrome-Backend/seeds"
	"gorm.io/gorm"
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
		&models.Plan{}, &models.DailyChallenge{}, &models.RecordHealth{}, &models.RecordPlan{},
		&models.QuizChallenge{}, &models.RecordQuiz{}, &models.RecordDaily{}, &models.Connect{}, &models.Knowledge{},
	)
	fmt.Println("? Migration complete")

	if !hasData(initializers.DB) {
		seeds.SeedKnowledge(initializers.DB)
		seeds.SeedQuiz(initializers.DB)
	}
}
func hasData(db *gorm.DB) bool {
	var count int64
	db.Model(&models.Knowledge{}).Count(&count)
	return count > 0
}
