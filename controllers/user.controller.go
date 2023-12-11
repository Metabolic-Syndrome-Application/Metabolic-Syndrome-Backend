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

// update and create profile myself

func (uc *UserController) UpdateProfile(ctx *gin.Context) {
	currentUser := ctx.MustGet("currentUser").(models.User)

	// patient
	if currentUser.Role == "patient" {
		var payload = struct {
			Alias       string `json:"alias,omitempty"`
			FirstName   string `json:"firstName"`
			LastName    string `json:"lastName"`
			YearOfBirth string `json:"yearOfBirth,omitempty"`
			Gender      string `json:"gender,omitempty"`
			Photo       string `json:"photo,omitempty"`
		}{} // {} = default is null
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

		// doctor
	} else if currentUser.Role == "doctor" {
		var payload = struct {
			Prefix     string `json:"prefix,omitempty"`
			FirstName  string `json:"firstName,omitempty"`
			LastName   string `json:"lastName,omitempty"`
			Gender     string `json:"gender,omitempty"`
			Department string `json:"department,omitempty"`
			Specialist string `json:"specialist,omitempty"`
		}{} // {} = default is null
		if err := ctx.ShouldBindJSON(&payload); err != nil {
			ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
			return
		}
		var updateProfileDoctor models.Doctor
		result := uc.DB.First(&updateProfileDoctor, "id = ?", currentUser.ID)
		if result.Error != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "No post with that title exists"})
			return
		}
		updateDoctor := &models.Doctor{
			Prefix:     payload.Prefix,
			FirstName:  payload.FirstName,
			LastName:   payload.LastName,
			Gender:     payload.Gender,
			Department: payload.Department,
			Specialist: payload.Specialist,
		}

		a := uc.DB.Model(&updateProfileDoctor).Updates(updateDoctor)
		if a.Error != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "Can not update profile Doctor"})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": "Update profile Doctor success"})

		// staff
	} else if currentUser.Role == "staff" {
		var payload = struct {
			Prefix     string `json:"prefix,omitempty"`
			FirstName  string `json:"firstName,omitempty"`
			LastName   string `json:"lastName,omitempty"`
			Gender     string `json:"gender,omitempty"`
			Department string `json:"department,omitempty"`
			Specialist string `json:"specialist,omitempty"`
		}{} // {} = default is null
		if err := ctx.ShouldBindJSON(&payload); err != nil {
			ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
			return
		}
		var updateProfileStaff models.Staff
		result := uc.DB.First(&updateProfileStaff, "id = ?", currentUser.ID)
		if result.Error != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "No post with that title exists"})
			return
		}
		updateStaff := &models.Staff{
			Prefix:     payload.Prefix,
			FirstName:  payload.FirstName,
			LastName:   payload.LastName,
			Gender:     payload.Gender,
			Department: payload.Department,
			Specialist: payload.Specialist,
		}

		a := uc.DB.Model(&updateProfileStaff).Updates(updateStaff)
		if a.Error != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "Can not update profile Staff"})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": "Update profile Staff success"})
	}

}

// get data profile myself

func (uc *UserController) GetProfile(ctx *gin.Context) {
	currentUser := ctx.MustGet("currentUser").(models.User)

	if currentUser.Role == "patient" {
		var GetProfilePatient models.Patient
		result := uc.DB.First(&GetProfilePatient, "id = ?", currentUser.ID)
		if result.Error != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "Not have this ID"})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": gin.H{"user": GetProfilePatient}})

	} else if currentUser.Role == "doctor" {
		var GetProfileDoctor models.Doctor
		result := uc.DB.First(&GetProfileDoctor, "id = ?", currentUser.ID)
		if result.Error != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "Not have this ID"})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": gin.H{"user": GetProfileDoctor}})

	} else if currentUser.Role == "staff" {
		var GetProfileStaff models.Staff
		result := uc.DB.First(&GetProfileStaff, "id = ?", currentUser.ID)
		if result.Error != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "Not have this ID"})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": gin.H{"user": GetProfileStaff}})

	}

}
