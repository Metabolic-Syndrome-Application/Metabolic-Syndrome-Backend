package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/ployns/Metabolic-Syndrome-Backend/models"
	"github.com/ployns/Metabolic-Syndrome-Backend/utils"
	"gorm.io/gorm"
)

type ConnectController struct {
	DB *gorm.DB
}

func NewConnectController(DB *gorm.DB) ConnectController {
	return ConnectController{DB}
}

func (cc *ConnectController) GenerateOTP(ctx *gin.Context) {
	var payload = struct {
		HN                 *string    `json:"hn"`
		FirstName          string     `json:"firstName"`
		LastName           string     `json:"lastName"`
		YearOfBirth        int        `json:"yearOfBirth"`
		Gender             string     `json:"gender"`
		MainDoctorID       *uuid.UUID `gorm:"type:uuid ;null" json:"mainDoctorID"`
		AssistanceDoctorID *uuid.UUID `gorm:"type:uuid ;null" json:"assistanceDoctorID"`
		DiseaseRisk        struct {
			Diabetes       string `json:"diabetes"`
			Hyperlipidemia string `json:"hyperlipidemia"`
			Hypertension   string `json:"hypertension"`
			Obesity        string `json:"obesity"`
		} `json:"diseaseRisk"`
	}{}
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	otp := utils.GenerateOTP(4)
	expiredIn := time.Now().Add(3 * time.Minute)

	newConnect := models.Connect{
		OTP:                otp,
		ExpiresIn:          expiredIn,
		HN:                 payload.HN,
		FirstName:          payload.FirstName,
		LastName:           payload.LastName,
		YearOfBirth:        payload.YearOfBirth,
		Gender:             payload.Gender,
		MainDoctorID:       payload.MainDoctorID,
		AssistanceDoctorID: payload.AssistanceDoctorID,
		DiseaseRisk:        payload.DiseaseRisk,
	}
	if err := cc.DB.Create(&newConnect).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "Can not create connect"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "message": "Create connect success", "data": gin.H{"id": newConnect.ID, "otp": otp}})
}

func (cc *ConnectController) RefreshOTP(ctx *gin.Context) {
	connectID := ctx.Param("id")
	var connect models.Connect
	result := cc.DB.First(&connect, "id = ?", connectID)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "Not have this ID"})
		return
	}
	otp := utils.GenerateOTP(4)
	expiredIn := time.Now().Add(3 * time.Minute)

	updateConnect := models.Connect{
		OTP:       otp,
		ExpiresIn: expiredIn,
	}
	if err := cc.DB.Model(&connect).Updates(updateConnect).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "Can not refresh otp"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": "Refresh OTP", "data": gin.H{"id": connect.ID, "otp": otp}})

}

func (cc *ConnectController) SubmitOTP(ctx *gin.Context) {
	currentUser := ctx.MustGet("currentUser").(models.User)
	var patient models.Patient
	result := cc.DB.First(&patient, "id = ?", currentUser.ID)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "No post with that title exists"})
		return
	}
	var payload = struct {
		OTP string `json:"otp"`
	}{}
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	var connect models.Connect
	date := time.Now()
	result2 := cc.DB.First(&connect, "otp = ? AND expires_in >= ?", payload.OTP, date)
	if result2.Error != nil {
		// not fond this row / not match
		ctx.JSON(http.StatusNotFound, gin.H{"status": "success", "message": "OTP not match"})
		return

		// match
	} else {
		updatePatient := &models.Patient{
			HN:                 connect.HN,
			FirstName:          connect.FirstName,
			LastName:           connect.LastName,
			YearOfBirth:        connect.YearOfBirth,
			Gender:             connect.Gender,
			MainDoctorID:       connect.MainDoctorID,
			AssistanceDoctorID: connect.AssistanceDoctorID,
			DiseaseRisk:        connect.DiseaseRisk,
		}

		result := cc.DB.Model(&patient).Updates(updatePatient)
		if result.Error != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "Can not update profile patient"})
			return
		}
		cc.DB.Delete(&models.Connect{}, "otp = ?", payload.OTP)
		cc.DB.Where("expires_in < ?", date).Delete(&models.Connect{})
		ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": "OTP matching"})

	}

}
