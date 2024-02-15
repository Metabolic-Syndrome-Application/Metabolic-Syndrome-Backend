package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ployns/Metabolic-Syndrome-Backend/models"
	"gorm.io/gorm"
)

type ChallengeController struct {
	DB *gorm.DB
}

func NewChallengeController(DB *gorm.DB) ChallengeController {
	return ChallengeController{DB}
}

// quiz

// create quiz challenge
func (cc *ChallengeController) CreateQuizChallenge(ctx *gin.Context) {
	var payload = struct {
		Question  string          `json:"question"`
		Choices   json.RawMessage `gorm:"type:json" json:"choices"`
		Points    int             `json:"points"`
		LimitTime int             `json:"limitTime"`
	}{}
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}
	newQuizChallenge := &models.QuizChallenge{
		Question:  payload.Question,
		Choices:   payload.Choices,
		Points:    payload.Points,
		LimitTime: payload.LimitTime,
	}
	result1 := cc.DB.Create(&newQuizChallenge)
	if result1.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "Can not create Quiz Challenge"})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "message": "Create Quiz Challenge success"})

}

// update quiz challenge

// get 1 quiz challenge

// get all quiz challenge

// get random quiz

// get point quiz challenge
