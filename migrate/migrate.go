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
	initializers.DB.AutoMigrate(&models.User{})
	fmt.Println("? Migration complete")
}
