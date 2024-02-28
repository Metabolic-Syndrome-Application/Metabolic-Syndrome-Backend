package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ployns/Metabolic-Syndrome-Backend/models"
	"gorm.io/gorm"
)

type KnowledgeController struct {
	DB *gorm.DB
}

func NewKnowledgeController(DB *gorm.DB) KnowledgeController {
	return KnowledgeController{DB}
}

// Get Knowledge
func (kc *KnowledgeController) GetKnowledge(ctx *gin.Context) {
	disease := ctx.Param("disease")
	var knowledge models.Knowledge
	result := kc.DB.First(&knowledge, "disease = ?", disease)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "Not have this disease"})
		return
	}

	type Response struct {
		Name        string `json:"name"`
		Details     string `json:"details"`
		Symptoms    string `json:"symptoms"`
		Medications string `json:"medications"`
		Behaviors   string `json:"behaviors"`
	}

	data := Response{
		Name:        knowledge.Name,
		Details:     knowledge.Details,
		Symptoms:    knowledge.Symptoms,
		Medications: knowledge.Medications,
		Behaviors:   knowledge.Behaviors,
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": data})

}

// Get all Knowledge
func (kc *KnowledgeController) GetAllKnowledge(ctx *gin.Context) {
	var knowledges []models.Knowledge
	result := kc.DB.Find(&knowledges)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "Not have disease"})
		return
	}

	type Response struct {
		Disease string `json:"disease"`
		Photo   string `json:"photo"`
		Name    string `json:"name"`
		Details string `json:"details"`
	}

	var data []Response
	for _, knowledge := range knowledges {
		response := Response{
			Disease: knowledge.Disease,
			Photo:   knowledge.Photo,
			Name:    knowledge.Name,
			Details: knowledge.Details,
		}
		data = append(data, response)
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": data})

}
