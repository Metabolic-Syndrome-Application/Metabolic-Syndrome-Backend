package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ployns/Metabolic-Syndrome-Backend/models"
	"gorm.io/gorm"
)

type RankController struct {
	DB *gorm.DB
}

func NewRankController(DB *gorm.DB) RankController {
	return RankController{DB}
}
func (rc *RankController) Top5(ctx *gin.Context) {
	var patients []models.Patient
	result := rc.DB.Order("collect_points desc").Limit(5).Find(&patients)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": result.Error})
		return
	}
	type Response struct {
		Name          string `json:"name"`
		CollectPoints int    `json:"collectPoints"`
	}
	var data []Response
	for _, patient := range patients {
		name := patient.Alias
		if patient.Alias == "" {
			name = patient.FirstName
		}
		response := Response{
			Name:          name,
			CollectPoints: patient.CollectPoints,
		}
		data = append(data, response)
	}
	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": data})
}

func (rc *RankController) MyRank(ctx *gin.Context) {
	currentUser := ctx.MustGet("currentUser").(models.User)
	var patient models.Patient
	result := rc.DB.First(&patient, "id = ?", currentUser.ID)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "No post with that title exists"})
		return
	}
	var rank int64
	result2 := rc.DB.Model(&models.Patient{}).Order("collect_points desc").Where("collect_points > ?", patient.CollectPoints).Count(&rank)
	if result2.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": result.Error})
		return
	}
	type Response struct {
		Rank          int64  `json:"rank"`
		Name          string `json:"name"`
		CollectPoints int    `json:"collectPoints"`
	}
	name := patient.Alias
	if patient.Alias == "" {
		name = patient.FirstName
	}
	data := Response{
		Rank:          rank + 1,
		Name:          name,
		CollectPoints: patient.CollectPoints,
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": data})
}
