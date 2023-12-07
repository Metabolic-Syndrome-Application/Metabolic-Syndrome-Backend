package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ployns/Metabolic-Syndrome-Backend/models"
	"gorm.io/gorm"
)

type UserController struct {
	DB *gorm.DB
}

func NewUserController(DB *gorm.DB) UserController {
	return UserController{DB}
}
func (uc *UserController) UpdateProfile(ctx *gin.Context) {
	currentUser := ctx.MustGet("currentUser").(models.User)

	var payload *models.UpdateProfilePatient
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}
	var updateProfilePatient models.Patient
	result := uc.DB.First(&updateProfilePatient, "id = ?", currentUser.ID)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "No post with that title exists"})
		return
	}
	updatePatient := &models.Patient{
		Alias:       payload.Alias,
		FirstName:   payload.FirstName,
		LastName:    payload.LastName,
		YearOfBirth: payload.YearOfBirth,
		Gender:      payload.Gender,
		Photo:       payload.Photo,
	}

	a := uc.DB.Model(&updateProfilePatient).Updates(updatePatient)
	if a.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "Can not update profile patient"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": "Update profile patient success"})

}

func (uc *UserController) GetProfile(ctx *gin.Context) {
	currentUser := ctx.MustGet("currentUser").(models.User)
	var GetProfilePatient models.Patient
	result := uc.DB.First(&GetProfilePatient, "id = ?", currentUser.ID)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "Not have this ID"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": gin.H{"user": GetProfilePatient}})
}
