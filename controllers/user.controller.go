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
			Alias        string     `json:"alias,omitempty"`
			FirstName    string     `json:"firstName"`
			LastName     string     `json:"lastName"`
			YearOfBirth  int        `json:"yearOfBirth,omitempty"`
			Gender       string     `json:"gender,omitempty"`
			Photo        string     `json:"photo,omitempty"`
			MainDoctorId *uuid.UUID `gorm:"type:uuid ;null" json:"mainDoctorId,omitempty"`
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
			Alias:        payload.Alias,
			FirstName:    payload.FirstName,
			LastName:     payload.LastName,
			YearOfBirth:  payload.YearOfBirth,
			Gender:       payload.Gender,
			Photo:        payload.Photo,
			MainDoctorId: payload.MainDoctorId,
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
	currentUser := ctx.MustGet("currentUser").(models.User)
	// staff can manage all patient
	if currentUser.Role == "staff" {
		var patients []models.Patient
		result := uc.DB.Find(&patients)
		if result.Error != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"status": "fail", "message": "Error fetching patient data"})
			return
		}
		type Response struct {
			ID           uuid.UUID  `json:"id"`
			HN           string     `json:"hn,omitempty"`
			FirstName    string     `json:"firstName,omitempty"`
			LastName     string     `json:"lastName,omitempty"`
			Gender       string     `json:"gender,omitempty"`
			YearOfBirth  int        `json:"yearOfBirth,omitempty"`
			Status       string     `json:"status,omitempty"`
			MainDoctorId *uuid.UUID `gorm:"type:uuid ;null" json:"mainDoctorId,omitempty"` //TODO:ใส่เป็นชื่อหมอ
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
				MainDoctorId: patient.MainDoctorId,
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
			ctx.JSON(http.StatusInternalServerError, gin.H{"status": "fail", "message": "Error fetching patient data"})
			return
		}
		type Response struct {
			ID         uuid.UUID `json:"id"`
			Username   string    `gorm:"type:varchar(255);not null"`
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
		result := uc.DB.Where("main_doctor_id = ?", currentUser.ID).Find(&patients)
		if result.Error != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"status": "fail", "message": "Error fetching patient data"})
			return
		}
		type Response struct {
			ID           uuid.UUID  `json:"id"`
			HN           string     `json:"hn,omitempty"`
			FirstName    string     `json:"firstName,omitempty"`
			LastName     string     `json:"lastName,omitempty"`
			Gender       string     `json:"gender,omitempty"`
			YearOfBirth  int        `json:"yearOfBirth,omitempty"`
			Status       string     `json:"status,omitempty"`
			MainDoctorId *uuid.UUID `gorm:"type:uuid ;null" json:"mainDoctorId,omitempty"` //TODO:ใส่เป็นชื่อหมอ
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
				MainDoctorId: patient.MainDoctorId,
			}
			data = append(data, response)
		}
		ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": gin.H{"users": data}})
	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "Invalid role"})
	}

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
			HN                 string     `json:"hn,omitempty"`
			FirstName          string     `json:"firstName,omitempty"`
			LastName           string     `json:"lastName,omitempty"`
			YearOfBirth        int        `json:"yearOfBirth,omitempty"`
			Gender             string     `json:"gender,omitempty"`
			MainDoctorId       *uuid.UUID `gorm:"type:uuid ;null" json:"mainDoctorId,omitempty"`
			AssistanceDoctorId *uuid.UUID `gorm:"type:uuid ;null" json:"assistanceDoctorId,omitempty"`
			DiseaseRisk        struct {
				Diabetes       string `json:"diabetes"`
				Hyperlipidemia string `json:"hyperlipidemia"`
				Hypertension   string `json:"hypertension"`
				Obesity        string `json:"obesity"`
			} `json:"diseaseRisk"`
			PlanID *uuid.UUID `gorm:"type:uuid ;null" json:"planID,omitempty"`
			Status string     `gorm:"default:'in process' " json:"status,omitempty"`
		}{} // {} = default is null
		if err := ctx.ShouldBindJSON(&payload); err != nil {
			ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
			return
		}
		updatePatient := &models.Patient{
			HN:                 payload.HN,
			FirstName:          payload.FirstName,
			LastName:           payload.LastName,
			YearOfBirth:        payload.YearOfBirth,
			Gender:             payload.Gender,
			MainDoctorId:       payload.MainDoctorId,
			AssistanceDoctorId: payload.AssistanceDoctorId,
			DiseaseRisk:        payload.DiseaseRisk,
			PlanID:             payload.PlanID,
			Status:             payload.Status,
		}
		a := uc.DB.Model(&patient).Updates(updatePatient)
		if a.Error != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "Can not update profile patient"})
			return
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
		updateDoctor := &models.Doctor{
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
		updateStaff := &models.Staff{
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
		result := uc.DB.First(&GetProfilePatient, "id = ?", userID)
		if result.Error != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "Not have this ID"})
			return
		}

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
func (pc *UserController) DeleteUser(ctx *gin.Context) {
	userRole := ctx.Param("role")
	userID := ctx.Param("id")
	result := pc.DB.Delete(&models.User{}, "id = ?", userID)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "No user with that title exists"})
		return
	}
	if userRole == "patient" {
		result := pc.DB.Delete(&models.Patient{}, "id = ?", userID)
		if result.Error != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "No user with that title exists"})
			return
		}
	} else if userRole == "doctor" {
		result := pc.DB.Delete(&models.Doctor{}, "id = ?", userID)
		if result.Error != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "No user with that title exists"})
			return
		}
	} else if userRole == "staff" {
		result := pc.DB.Delete(&models.Staff{}, "id = ?", userID)
		if result.Error != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "No user with that title exists"})
			return
		}
	}
	ctx.JSON(http.StatusOK, gin.H{"status": "success"})
}
