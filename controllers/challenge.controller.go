package controllers

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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

//web

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
func (cc *ChallengeController) UpdateQuizChallenge(ctx *gin.Context) {
	quizID := ctx.Param("id")
	var quiz models.QuizChallenge
	result := cc.DB.First(&quiz, "id = ?", quizID)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "Not have this ID"})
		return
	}
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
	updateQuizChallenge := &models.QuizChallenge{
		Question:  payload.Question,
		Choices:   payload.Choices,
		Points:    payload.Points,
		LimitTime: payload.LimitTime,
	}
	result1 := cc.DB.Model(&quiz).Updates(updateQuizChallenge)
	if result1.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "Can not update Quiz Challenge"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": "Update Quiz Challenge success"})

}

// delete quiz challenge
func (cc *ChallengeController) DeleteQuizChallenge(ctx *gin.Context) {
	quizID := ctx.Param("id")
	result := cc.DB.Delete(&models.QuizChallenge{}, "id = ?", quizID)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "No quiz with that ID exists"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": "Delete Quiz Challenge success"})
}

// get 1 quiz challenge
func (cc *ChallengeController) GetQuizChallenge(ctx *gin.Context) {
	quizID := ctx.Param("id")
	var quiz models.QuizChallenge
	result := cc.DB.First(&quiz, "id = ?", quizID)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "Not have this ID"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": gin.H{"quiz": quiz}})

}

// get all quiz challenge
func (cc *ChallengeController) GetAllQuizChallenge(ctx *gin.Context) {
	var quizs []models.QuizChallenge
	result := cc.DB.Find(&quizs)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "not have plan data"})
		return
	}
	type Response struct {
		ID        uuid.UUID `json:"id"`
		Question  string    `json:"question"`
		Points    int       `json:"points"`
		LimitTime int       `json:"limitTime"`
	}
	var data []Response
	for _, quiz := range quizs {
		response := Response{
			ID:        quiz.ID,
			Question:  quiz.Question,
			Points:    quiz.Points,
			LimitTime: quiz.LimitTime,
		}
		data = append(data, response)
	}
	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": gin.H{"quiz": data}})
}

// mobile

// Check Quiz Today
func (cc *ChallengeController) CheckQuizToday(ctx *gin.Context) {
	currentUser := ctx.MustGet("currentUser").(models.User)
	var recordQuiz models.RecordQuiz
	date := time.Now().UTC().Truncate(24 * time.Hour)
	result := cc.DB.First(&recordQuiz, "patient_id = ? AND created_at >= ? AND created_at < ?", currentUser.ID, date, date.Add(24*time.Hour))
	if result.Error != nil {
		// not fond this row
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": "You haven't done the quiz today", "check": false})
		} else {
			ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "error"})
		}
	} else {
		ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": "You have already done today's quiz", "check": true})

	}
}

// get random quiz
func (cc *ChallengeController) GetRandomQuiz(ctx *gin.Context) {
	currentUser := ctx.MustGet("currentUser").(models.User)
	var quiz models.QuizChallenge
	cc.DB.Order("RANDOM()").Limit(1).Find(&quiz)
	createQuizToday := &models.RecordQuiz{
		PatientID:       currentUser.ID,
		QuizChallengeID: quiz.ID,
	}
	cc.DB.Create(&createQuizToday)
	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": gin.H{"quiz": quiz}})
}

// get point quiz challenge
func (cc *ChallengeController) GetPointQuiz(ctx *gin.Context) {
	currentUser := ctx.MustGet("currentUser").(models.User)
	var patient models.Patient
	result := cc.DB.First(&patient, "id = ?", currentUser.ID)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "Not have this ID"})
		return
	}
	var payload = struct {
		IsCorrect bool `json:"isCorrect"`
	}{}
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}
	if payload.IsCorrect == true {
		updatePoint := &models.Patient{
			CollectPoints: patient.CollectPoints + 150,
		}
		result = cc.DB.Model(&patient).Updates(updatePoint)
		if result.Error != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "Can not update point"})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": "correct, get 150 points", "result": "correct"})

	} else {
		ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": "incorrect, don't get point", "result": "incorrect"})

	}
}

//////////////////////////////////////////////////////////////////////////////////////////////////

// daily

// web

// Create daily
func (cc *ChallengeController) CreateDailyChallenge(ctx *gin.Context) {
	var payload = struct {
		Name        string          `json:"name"`
		Description string          `json:"description"`
		Photo       string          `json:"photo"`
		Detail      json.RawMessage `gorm:"type:json" json:"detail"`
		Points      int             `json:"points"`
		NumDays     int             `json:"numDays"`
	}{}
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	var daily models.DailyChallenge
	existingDaily := cc.DB.First(&daily, "name = ?", payload.Name)
	if existingDaily.Error == nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "This daily's name is already in use"})
		return
	}
	newDaily := models.DailyChallenge{
		Name:        payload.Name,
		Description: payload.Description,
		Photo:       payload.Photo,
		Detail:      payload.Detail,
		Points:      payload.Points,
		NumDays:     payload.NumDays,
	}
	if err := cc.DB.Create(&newDaily).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "Can not create daily"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "message": "Create daily success"})
}

// Update Pdaily
func (cc *ChallengeController) UpdateDailyChallenge(ctx *gin.Context) {
	dailyID := ctx.Param("id")
	var daily models.DailyChallenge
	result := cc.DB.First(&daily, "id = ?", dailyID)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "Not have this ID"})
		return
	}
	var payload = struct {
		Name        string          `json:"name"`
		Description string          `json:"description"`
		Photo       string          `json:"photo"`
		Detail      json.RawMessage `gorm:"type:json" json:"detail"`
		Points      int             `json:"points"`
		NumDays     int             `json:"numDays"`
		Status      string          `json:"status"`
	}{}
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}
	updateDaily := models.DailyChallenge{
		Name:        payload.Name,
		Description: payload.Description,
		Photo:       payload.Photo,
		Detail:      payload.Detail,
		Points:      payload.Points,
		NumDays:     payload.NumDays,
		Status:      payload.Status,
	}
	if err := cc.DB.Model(&daily).Updates(updateDaily).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "Can not update daily"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": "Update daily success"})
}
