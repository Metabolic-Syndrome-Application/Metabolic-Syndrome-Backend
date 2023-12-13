package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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

// get profile all
func (uc *UserController) GetAllUserProfile(ctx *gin.Context) {
	userRole := ctx.Param("role")
	if userRole == "patient" {
		var patients []models.Patient
		result := uc.DB.Find(&patients)
		if result.Error != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"status": "fail", "message": "Error fetching patient data"})
			return
		}
		type Response struct {
			ID          uuid.UUID `json:"id"`
			HN          string    `json:"hn,omitempty"`
			FirstName   string    `json:"firstName,omitempty"`
			LastName    string    `json:"lastName,omitempty"`
			Gender      string    `json:"gender,omitempty"`
			YearOfBirth string    `json:"yearOfBirth,omitempty"`
			Status      string    `json:"status,omitempty"`
			// MainDoctorId *uuid.UUID `gorm:"type:uuid ;null" json:"mainDoctorId,omitempty"` //TODO:ใส่เป็นชื่อหมอ
		}
		var data []Response
		for _, patient := range patients {
			response := Response{
				ID:          patient.ID,
				HN:          patient.HN,
				FirstName:   patient.FirstName,
				LastName:    patient.LastName,
				Gender:      patient.Gender,
				YearOfBirth: patient.YearOfBirth,
				Status:      patient.Status,
			}
			data = append(data, response)
		}
		ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": gin.H{"users": data}})
	} else if userRole == "doctor-staff" {
		var doctors []models.Doctor
		result1 := uc.DB.Find(&doctors)
		var staffs []models.Staff
		result2 := uc.DB.Find(&staffs)
		if result1.Error != nil || result2.Error != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"status": "fail", "message": "Error fetching patient data"})
			return
		}
		type Response struct {
			ID         uuid.UUID `json:"id"`
			Prefix     string    `json:"prefix,omitempty"`
			FirstName  string    `json:"firstName,omitempty"`
			LastName   string    `json:"lastName,omitempty"`
			Gender     string    `json:"gender,omitempty"`
			Department string    `json:"department,omitempty"`
			Specialist string    `json:"specialist,omitempty"`
			Role       string    `gorm:"type:varchar(255);not null"`
		}
		var data []Response
		for _, doctor := range doctors {
			response := Response{
				ID:         doctor.ID,
				Prefix:     doctor.Prefix,
				FirstName:  doctor.FirstName,
				LastName:   doctor.LastName,
				Gender:     doctor.Gender,
				Department: doctor.Department,
				Specialist: doctor.Specialist,
				Role:       "doctor",
			}
			data = append(data, response)
		}
		for _, staff := range staffs {
			response := Response{
				ID:         staff.ID,
				Prefix:     staff.Prefix,
				FirstName:  staff.FirstName,
				LastName:   staff.LastName,
				Gender:     staff.Gender,
				Department: staff.Department,
				Specialist: staff.Specialist,
				Role:       "staff",
			}
			data = append(data, response)
		}
		ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": gin.H{"users": data}})

	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "Invalid role"})
	}

}

// 	userRole := ctx.Param("role")
// 	userID := ctx.Param("id")

// 	var patient models.Patient
// 	var doctor models.Doctor
// 	var staff models.Staff

// 	if err := uc.DB.First(&patient, "id = ?", userID).Error; err == nil {
// 		var payload = struct {
// 			Alias       string `json:"alias,omitempty"`
// 			FirstName   string `json:"firstName"`
// 			LastName    string `json:"lastName"`
// 			YearOfBirth string `json:"yearOfBirth,omitempty"`
// 			Gender      string `json:"gender,omitempty"`
// 			Photo       string `json:"photo,omitempty"`
// 		}{} // {} = default is null
// 		if err := ctx.ShouldBindJSON(&payload); err != nil {
// 			ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
// 			return
// 		}

// 		updatePatient := &models.Patient{
// 			Alias:       payload.Alias,
// 			FirstName:   payload.FirstName,
// 			LastName:    payload.LastName,
// 			YearOfBirth: payload.YearOfBirth,
// 			Gender:      payload.Gender,
// 			Photo:       payload.Photo,
// 		}

// 		a := uc.DB.Model(&patient).Updates(updatePatient)
// 		if a.Error != nil {
// 			ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "Can not update profile patient"})
// 			return
// 		}
// 		ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": "Update profile patient success"})
// 		return
// 	} else if err := uc.DB.First(&doctor, "id = ?", userID).Error; err == nil {
// 		ctx.JSON(http.StatusOK, gin.H{"user": doctor, "role": "doctor"})
// 		return
// 	} else if err := uc.DB.First(&staff, "id = ?", userID).Error; err == nil {
// 		ctx.JSON(http.StatusOK, gin.H{"user": staff, "role": "staff"})
// 		return
// 	} else {
// 		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "No post with that title exists"})
// 	}

// 	currentUser := ctx.MustGet("currentUser").(models.User)

// 	// patient
// 	if currentUser.Role == "patient" {
// 		var payload = struct {
// 			Alias       string `json:"alias,omitempty"`
// 			FirstName   string `json:"firstName"`
// 			LastName    string `json:"lastName"`
// 			YearOfBirth string `json:"yearOfBirth,omitempty"`
// 			Gender      string `json:"gender,omitempty"`
// 			Photo       string `json:"photo,omitempty"`
// 		}{} // {} = default is null
// 		if err := ctx.ShouldBindJSON(&payload); err != nil {
// 			ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
// 			return
// 		}
// 		var updateProfilePatient models.Patient
// 		result := uc.DB.First(&updateProfilePatient, "id = ?", currentUser.ID)
// 		if result.Error != nil {
// 			ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "No post with that title exists"})
// 			return
// 		}
// 		updatePatient := &models.Patient{
// 			Alias:       payload.Alias,
// 			FirstName:   payload.FirstName,
// 			LastName:    payload.LastName,
// 			YearOfBirth: payload.YearOfBirth,
// 			Gender:      payload.Gender,
// 			Photo:       payload.Photo,
// 		}

// 		a := uc.DB.Model(&updateProfilePatient).Updates(updatePatient)
// 		if a.Error != nil {
// 			ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "Can not update profile patient"})
// 			return
// 		}
// 		ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": "Update profile patient success"})

// 		// doctor
// 	} else if currentUser.Role == "doctor" {
// 		var payload = struct {
// 			Prefix     string `json:"prefix,omitempty"`
// 			FirstName  string `json:"firstName,omitempty"`
// 			LastName   string `json:"lastName,omitempty"`
// 			Gender     string `json:"gender,omitempty"`
// 			Department string `json:"department,omitempty"`
// 			Specialist string `json:"specialist,omitempty"`
// 		}{} // {} = default is null
// 		if err := ctx.ShouldBindJSON(&payload); err != nil {
// 			ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
// 			return
// 		}
// 		var updateProfileDoctor models.Doctor
// 		result := uc.DB.First(&updateProfileDoctor, "id = ?", currentUser.ID)
// 		if result.Error != nil {
// 			ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "No post with that title exists"})
// 			return
// 		}
// 		updateDoctor := &models.Doctor{
// 			Prefix:     payload.Prefix,
// 			FirstName:  payload.FirstName,
// 			LastName:   payload.LastName,
// 			Gender:     payload.Gender,
// 			Department: payload.Department,
// 			Specialist: payload.Specialist,
// 		}

// 		a := uc.DB.Model(&updateProfileDoctor).Updates(updateDoctor)
// 		if a.Error != nil {
// 			ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "Can not update profile Doctor"})
// 			return
// 		}
// 		ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": "Update profile Doctor success"})

// 		// staff
// 	} else if currentUser.Role == "staff" {
// 		var payload = struct {
// 			Prefix     string `json:"prefix,omitempty"`
// 			FirstName  string `json:"firstName,omitempty"`
// 			LastName   string `json:"lastName,omitempty"`
// 			Gender     string `json:"gender,omitempty"`
// 			Department string `json:"department,omitempty"`
// 			Specialist string `json:"specialist,omitempty"`
// 		}{} // {} = default is null
// 		if err := ctx.ShouldBindJSON(&payload); err != nil {
// 			ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
// 			return
// 		}
// 		var updateProfileStaff models.Staff
// 		result := uc.DB.First(&updateProfileStaff, "id = ?", currentUser.ID)
// 		if result.Error != nil {
// 			ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "No post with that title exists"})
// 			return
// 		}
// 		updateStaff := &models.Staff{
// 			Prefix:     payload.Prefix,
// 			FirstName:  payload.FirstName,
// 			LastName:   payload.LastName,
// 			Gender:     payload.Gender,
// 			Department: payload.Department,
// 			Specialist: payload.Specialist,
// 		}

// 		a := uc.DB.Model(&updateProfileStaff).Updates(updateStaff)
// 		if a.Error != nil {
// 			ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "Can not update profile Staff"})
// 			return
// 		}
// 		ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": "Update profile Staff success"})
// 	}

// }
