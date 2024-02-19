package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/lib/pq"
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
			Alias       string     `json:"alias"`
			FirstName   string     `json:"firstName"`
			LastName    string     `json:"lastName"`
			YearOfBirth int        `json:"yearOfBirth"`
			Gender      string     `json:"gender"`
			Photo       string     `json:"photo"`
			ChallengeID *uuid.UUID `gorm:"type:uuid ;null" json:"challengeID"`
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

		if updateProfilePatient.ChallengeID != payload.ChallengeID {
			var dailyChallenge models.DailyChallenge
			uc.DB.First(&dailyChallenge, "id = ?", payload.ChallengeID)
			updateDaily := &models.DailyChallenge{
				Participants: dailyChallenge.Participants + 1,
			}
			uc.DB.Model(&dailyChallenge).Updates(updateDaily)
		}

		updatePatient := &models.Patient{
			Alias:       payload.Alias,
			FirstName:   payload.FirstName,
			LastName:    payload.LastName,
			YearOfBirth: payload.YearOfBirth,
			Gender:      payload.Gender,
			Photo:       payload.Photo,
			ChallengeID: payload.ChallengeID,
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
			Prefix     string `json:"prefix"`
			FirstName  string `json:"firstName"`
			LastName   string `json:"lastName"`
			Gender     string `json:"gender"`
			Department string `json:"department"`
			Specialist string `json:"specialist"`
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
			Prefix     string `json:"prefix"`
			FirstName  string `json:"firstName"`
			LastName   string `json:"lastName"`
			Gender     string `json:"gender"`
			Department string `json:"department"`
			Specialist string `json:"specialist"`
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
		result := uc.DB.Preload("Plan").
			Preload("MainDoctor").
			Preload("AssistanceDoctor").
			Preload("Challenge").
			First(&GetProfilePatient, "id = ?", currentUser.ID)
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
	currentUser := ctx.MustGet("currentUser").(models.User)
	// staff can manage all patient
	if currentUser.Role == "staff" {
		var patients []models.Patient
		result := uc.DB.Where("hn IS NOT NULL AND hn != ''").Find(&patients)
		if result.Error != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "Error fetching patient data"})
			return
		}
		type Response struct {
			ID           uuid.UUID  `json:"id"`
			HN           *string    `json:"hn"`
			FirstName    string     `json:"firstName"`
			LastName     string     `json:"lastName"`
			Gender       string     `json:"gender"`
			YearOfBirth  int        `json:"yearOfBirth"`
			Status       string     `json:"status"`
			MainDoctorID *uuid.UUID `gorm:"type:uuid ;null" json:"mainDoctorID"` //TODO:ใส่เป็นชื่อหมอ
		}
		var data []Response
		for _, patient := range patients {
			response := Response{
				ID:           patient.ID,
				HN:           patient.HN,
				FirstName:    patient.FirstName,
				LastName:     patient.LastName,
				Gender:       patient.Gender,
				YearOfBirth:  patient.YearOfBirth,
				Status:       patient.Status,
				MainDoctorID: patient.MainDoctorID,
			}
			data = append(data, response)
		}
		ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": gin.H{"users": data}})
		// admin can manage all doctor and staff
	} else if currentUser.Role == "admin" {
		var doctors []models.Doctor
		result1 := uc.DB.Find(&doctors)
		var staffs []models.Staff
		result2 := uc.DB.Find(&staffs)
		if result1.Error != nil || result2.Error != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "Error fetching patient data"})
			return
		}
		type Response struct {
			ID         uuid.UUID `json:"id"`
			Username   string    `gorm:"type:varchar(255);not null" json:"username"`
			Prefix     string    `json:"prefix"`
			FirstName  string    `json:"firstName"`
			LastName   string    `json:"lastName"`
			Gender     string    `json:"gender"`
			Department string    `json:"department"`
			Specialist string    `json:"specialist"`
			Role       string    `gorm:"type:varchar(255);not null" json:"role"`
		}
		var data []Response
		for _, doctor := range doctors {
			response := Response{
				ID:         doctor.ID,
				Username:   doctor.Username,
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
				Username:   staff.Username,
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

	} else if currentUser.Role == "doctor" {
		var patients []models.Patient
		result := uc.DB.Where("(main_doctor_id = ? OR assistance_doctor_id = ? ) AND hn IS NOT NULL AND hn != ''", currentUser.ID, currentUser.ID).Find(&patients)
		if result.Error != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "Error fetching doctor data"})
			return
		}
		type Response struct {
			ID           uuid.UUID  `json:"id"`
			HN           *string    `json:"hn"`
			FirstName    string     `json:"firstName"`
			LastName     string     `json:"lastName"`
			Gender       string     `json:"gender"`
			YearOfBirth  int        `json:"yearOfBirth"`
			Status       string     `json:"status"`
			MainDoctorID *uuid.UUID `gorm:"type:uuid ;null" json:"mainDoctorID"` //TODO:ใส่เป็นชื่อหมอ
		}
		var data []Response
		for _, patient := range patients {
			response := Response{
				ID:           patient.ID,
				HN:           patient.HN,
				FirstName:    patient.FirstName,
				LastName:     patient.LastName,
				Gender:       patient.Gender,
				YearOfBirth:  patient.YearOfBirth,
				Status:       patient.Status,
				MainDoctorID: patient.MainDoctorID,
			}
			data = append(data, response)
		}
		ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": gin.H{"users": data}})
	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "Invalid role"})
	}

}

// get profile doctor all (dropdown maindoctor/assistancedoctor)

func (uc *UserController) GetAllUserProfileDoctor(ctx *gin.Context) {
	var doctors []models.Doctor
	result := uc.DB.Find(&doctors)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "Error fetching patient data"})
		return
	}
	type Response struct {
		ID        uuid.UUID `json:"id"`
		Prefix    string    `json:"prefix"`
		FirstName string    `json:"firstName"`
		LastName  string    `json:"lastName"`
	}
	var data []Response
	for _, doctor := range doctors {
		response := Response{
			ID:        doctor.ID,
			Prefix:    doctor.Prefix,
			FirstName: doctor.FirstName,
			LastName:  doctor.LastName,
		}
		data = append(data, response)
	}
	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": gin.H{"users": data}})

}

// update and create profile other user

func (uc *UserController) UpdateOtherProfile(ctx *gin.Context) {
	userRole := ctx.Param("role")
	userID := ctx.Param("id")

	// update and create profile role patient

	if userRole == "patient" {
		var patient models.Patient
		result := uc.DB.First(&patient, "id = ?", userID)
		if result.Error != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "Not have this ID"})
			return
		}

		var payload = struct {
			HN                 *string    `json:"hn"`
			FirstName          string     `json:"firstName"`
			LastName           string     `json:"lastName"`
			YearOfBirth        int        `json:"yearOfBirth"`
			Gender             string     `json:"gender"`
			MainDoctorID       *uuid.UUID `gorm:"type:uuid;null" json:"mainDoctorID"`
			AssistanceDoctorID *uuid.UUID `gorm:"type:uuid;null" json:"assistanceDoctorID"`
			DiseaseRisk        struct {
				Diabetes       string `json:"diabetes"`
				Hyperlipidemia string `json:"hyperlipidemia"`
				Hypertension   string `json:"hypertension"`
				Obesity        string `json:"obesity"`
			} `json:"diseaseRisk"`
			PlanID pq.StringArray `gorm:"type:uuid[];column:plan_id" json:"planID"`
			Status string         `gorm:"default:'in process'" json:"status"`
		}{}

		if err := ctx.ShouldBindJSON(&payload); err != nil {
			ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
			return
		}

		// planID = null => planDefault (Original plan)
		var planID pq.StringArray
		if payload.PlanID == nil {
			planID = patient.PlanID
		} else {
			planID = payload.PlanID
		}

		updatePatient := &models.Patient{
			ID:                 patient.ID,
			HN:                 payload.HN,
			FirstName:          payload.FirstName,
			LastName:           payload.LastName,
			YearOfBirth:        payload.YearOfBirth,
			Gender:             payload.Gender,
			MainDoctorID:       payload.MainDoctorID,
			AssistanceDoctorID: payload.AssistanceDoctorID,
			DiseaseRisk:        payload.DiseaseRisk,
			PlanID:             planID,
			Status:             payload.Status,
		}

		result = uc.DB.Model(&patient).Updates(updatePatient)
		if result.Error != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "Can not update profile patient"})
			return
		}

		// table patient_plan fotr patient.Plan
		var plans []models.Plan
		var plansToRemove []models.Plan

		uc.DB.Find(&plansToRemove)
		uc.DB.Model(&patient).Association("Plan").Delete(&plansToRemove)

		for _, planID := range patient.PlanID {

			uc.DB.Where("id = ?", planID).Find(&plans)
			uc.DB.Model(&patient).Association("Plan").Append(&plans)

		}

		ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": "Update profile patient success"})

		// update and create profile role doctor

	} else if userRole == "doctor" {
		var doctor models.Doctor
		result := uc.DB.First(&doctor, "id = ?", userID)
		if result.Error != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "Not have this ID"})
			return
		}
		var payload = struct {
			Prefix     string `json:"prefix"`
			FirstName  string `json:"firstName"`
			LastName   string `json:"lastName"`
			Gender     string `json:"gender"`
			Department string `json:"department"`
			Specialist string `json:"specialist"`
		}{} // {} = default is null
		if err := ctx.ShouldBindJSON(&payload); err != nil {
			ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
			return
		}
		updateDoctor := &models.Doctor{
			ID:         doctor.ID,
			Prefix:     payload.Prefix,
			FirstName:  payload.FirstName,
			LastName:   payload.LastName,
			Gender:     payload.Gender,
			Department: payload.Department,
			Specialist: payload.Specialist,
		}
		a := uc.DB.Model(&doctor).Updates(updateDoctor)
		if a.Error != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "Can not update profile doctor"})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": "Update profile doctor success"})

	} else if userRole == "staff" {
		var staff models.Staff
		result := uc.DB.First(&staff, "id = ?", userID)
		if result.Error != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "Not have this ID"})
			return
		}
		var payload = struct {
			Prefix     string `json:"prefix"`
			FirstName  string `json:"firstName"`
			LastName   string `json:"lastName"`
			Gender     string `json:"gender"`
			Department string `json:"department"`
			Specialist string `json:"specialist"`
		}{} // {} = default is null
		if err := ctx.ShouldBindJSON(&payload); err != nil {
			ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
			return
		}
		updateStaff := &models.Staff{
			ID:         staff.ID,
			Prefix:     payload.Prefix,
			FirstName:  payload.FirstName,
			LastName:   payload.LastName,
			Gender:     payload.Gender,
			Department: payload.Department,
			Specialist: payload.Specialist,
		}
		a := uc.DB.Model(&staff).Updates(updateStaff)
		if a.Error != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "Can not update profile staff"})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": "Update profile staff success"})

	} else {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "Not have this role"})
	}

}

// get profile other user

func (uc *UserController) GetOtherProfile(ctx *gin.Context) {
	userRole := ctx.Param("role")
	userID := ctx.Param("id")
	if userRole == "patient" {
		var GetProfilePatient models.Patient

		result := uc.DB.Preload("Plan").
			Preload("MainDoctor").
			Preload("AssistanceDoctor").
			Preload("Challenge").
			First(&GetProfilePatient, "id = ?", userID)
		if result.Error != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "Not have this ID"})
			return
		}
		// uc.DB.Model(&GetProfilePatient).Association("Plan").Find(&GetProfilePatient.Plan)

		ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": gin.H{"user": GetProfilePatient}})

	} else if userRole == "doctor" {
		var GetProfileDoctor models.Doctor
		result := uc.DB.First(&GetProfileDoctor, "id = ?", userID)
		if result.Error != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "Not have this ID"})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": gin.H{"user": GetProfileDoctor}})

	} else if userRole == "staff" {
		var GetProfileStaff models.Staff
		result := uc.DB.First(&GetProfileStaff, "id = ?", userID)
		if result.Error != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "Not have this ID"})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": gin.H{"user": GetProfileStaff}})
	} else {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "Not have this role"})
	}
}
func (uc *UserController) DeleteUser(ctx *gin.Context) {
	userRole := ctx.Param("role")
	userID := ctx.Param("id")
	result := uc.DB.Delete(&models.User{}, "id = ?", userID)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "No user with that title exists"})
		return
	}
	if userRole == "patient" {
		// Delete the references to the patient from the patient_plan table
		if err := uc.DB.Exec("DELETE FROM patient_plan WHERE patient_id = ?", userID).Error; err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"status": "fail", "message": "Failed to delete references to the plan"})
			return
		}
		result := uc.DB.Delete(&models.Patient{}, "id = ?", userID)
		if result.Error != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "No user with that title exists"})
			return
		}
	} else if userRole == "doctor" {
		if err := uc.DB.Model(&models.Patient{}).Where("main_doctor_id = ?", userID).Update("main_doctor_id", nil).Error; err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"status": "fail", "message": "error"})
			return
		}
		if err := uc.DB.Model(&models.Patient{}).Where("assistance_doctor_id = ?", userID).Update("assistance_doctor_id", nil).Error; err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"status": "fail", "message": "error"})
			return
		}

		result := uc.DB.Delete(&models.Doctor{}, "id = ?", userID)
		if result.Error != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "No user with that title exists"})
			return
		}

	} else if userRole == "staff" {
		result := uc.DB.Delete(&models.Staff{}, "id = ?", userID)
		if result.Error != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "No user with that title exists"})
			return
		}
	}
	ctx.JSON(http.StatusOK, gin.H{"status": "success"})
}
