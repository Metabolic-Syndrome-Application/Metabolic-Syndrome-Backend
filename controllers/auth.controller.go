package controllers

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ployns/Metabolic-Syndrome-Backend/initializers"
	"github.com/ployns/Metabolic-Syndrome-Backend/models"
	"github.com/ployns/Metabolic-Syndrome-Backend/utils"
	"gorm.io/gorm"
)

type AuthController struct {
	DB *gorm.DB
}

func NewAuthController(DB *gorm.DB) AuthController {
	return AuthController{DB}
}

// [...] SignUp User
func (ac *AuthController) SignUpUser(ctx *gin.Context) {
	var payload *models.SignUpInput

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	var user models.User
	existingUser := ac.DB.First(&user, "Username = ?", strings.ToLower(payload.Username))
	if existingUser.Error == nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "Username is already in use"})
		return
	}
	// hash 13 number
	// if !strings.Contains(payload.Username, "@") {
	// 	hashedUsername, _ := utils.HashPassword(strings.ToLower(payload.Username))
	// 	existingUser := ac.DB.First(&user, "Username = ?", hashedUsername)
	// 	if existingUser.Error == nil {
	// 		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "Username is already in use"})
	// 		return
	// 	}
	// } else {
	// 	existingUser := ac.DB.First(&user, "Username = ?", strings.ToLower(payload.Username))
	// 	if existingUser.Error == nil {
	// 		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "Username is already in use"})
	// 		return
	// 	}

	if payload.Password != payload.PasswordConfirm {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "Passwords do not match"})
		return
	}

	// hashedPassword
	hashedPassword, err := utils.HashPassword(payload.Password)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": err.Error()})
		return
	}

	now := time.Now()
	newUser := models.User{
		Password:  hashedPassword,
		Role:      payload.Role,
		Verified:  true,
		CreatedAt: now,
		UpdatedAt: now,
	}

	if !strings.Contains(payload.Username, "@") {
		newUser.Type = "ID card number"
	} else {
		newUser.Type = "email"
	}

	newUser.Username = strings.ToLower(payload.Username)

	// สร้าง db User
	result := ac.DB.Create(&newUser)

	// สร้าง db patient
	if newUser.Role == "patient" {
		newPatient := &models.Patient{
			ID: newUser.ID,
		}
		a := ac.DB.Create(&newPatient)
		if a.Error != nil {
			fmt.Println("Error:", a.Error)
		} else {
			fmt.Println("User created successfully /patient")
		}
	} else if newUser.Role == "doctor" {
		newDoctor := &models.Doctor{
			ID:       newUser.ID,
			Username: newUser.Username,
		}
		a := ac.DB.Create(&newDoctor)
		if a.Error != nil {
			fmt.Println("Error:", a.Error)
		} else {
			fmt.Println("User created successfully /doctor")
		}
	} else if newUser.Role == "staff" {
		newStaff := &models.Staff{
			ID:       newUser.ID,
			Username: newUser.Username,
		}
		a := ac.DB.Create(&newStaff)
		if a.Error != nil {
			fmt.Println("Error:", a.Error)
		} else {
			fmt.Println("User created successfully /staff")
		}
	}

	if result.Error != nil && strings.Contains(result.Error.Error(), "duplicate key value violates unique") {
		ctx.JSON(http.StatusConflict, gin.H{"status": "fail", "message": "User with that Username already exists"})
		return
	} else if result.Error != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": "Something bad happened"})
		return
	}

	userResponse := &models.UserResponse{
		ID:        newUser.ID,
		Username:  newUser.Username,
		Role:      newUser.Role,
		CreatedAt: newUser.CreatedAt,
		UpdatedAt: newUser.UpdatedAt,
	}
	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": gin.H{"user": userResponse}})
}

// [...] SignIn User
func (ac *AuthController) SignInUser(ctx *gin.Context) {
	var payload *models.SignInInput

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	var user models.User
	result := ac.DB.First(&user, "Username = ?", strings.ToLower(payload.Username))
	if result.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "Invalid Username or Password"})
		return
	}

	//hash 13 number
	// if !strings.Contains(payload.Username, "@") {
	// 	fmt.Println("not have @")
	// 	hashedUsername, err := utils.HashPassword(strings.ToLower(payload.Username))
	// 	if err != nil {
	// 		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message1": err.Error()})
	// 		return
	// 	}
	// 	result := ac.DB.First(&user, "Username = ?", hashedUsername)
	// 	fmt.Println("hashedUsername", hashedUsername)
	// 	if result.Error != nil {
	// 		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message2": "Invalid Username or Password"})
	// 		return
	// 	}
	// } else {
	// 	result := ac.DB.First(&user, "Username = ?", strings.ToLower(payload.Username))
	// 	fmt.Println("have @")
	// 	if result.Error != nil {
	// 		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message3": "Invalid Username or Password"})
	// 		return
	// 	}
	// }

	if err := utils.VerifyPassword(user.Password, payload.Password); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "Invalid Username or Password"})
		return
	}

	config, _ := initializers.LoadConfig(".")

	// Generate Tokens
	access_token, err := utils.CreateToken(config.AccessTokenExpiresIn, user.ID, config.AccessTokenPrivateKey)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	refresh_token, err := utils.CreateToken(config.RefreshTokenExpiresIn, user.ID, config.RefreshTokenPrivateKey)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	ctx.SetCookie("access_token", access_token, config.AccessTokenMaxAge*60, "/", config.HOST, false, true)
	ctx.SetCookie("refresh_token", refresh_token, config.RefreshTokenMaxAge*60, "/", config.HOST, false, true)
	ctx.SetCookie("logged_in", "true", config.AccessTokenMaxAge*60, "/", config.HOST, false, false)

	if user.Role == "patient" {

		ctx.JSON(http.StatusOK, gin.H{"status": "success", "user": gin.H{"role": user.Role}, "access_token": access_token})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"status": "success", "user": gin.H{"username": user.Username, "role": user.Role}, "access_token": access_token})

	}
}

// [...] Refresh Access Token
func (ac *AuthController) RefreshAccessToken(ctx *gin.Context) {
	message := "could not refresh access token"

	cookie, err := ctx.Cookie("refresh_token")

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"status": "fail", "message": message})
		return
	}

	config, _ := initializers.LoadConfig(".")

	sub, err := utils.ValidateToken(cookie, config.RefreshTokenPublicKey)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	var user models.User
	result := ac.DB.First(&user, "id = ?", fmt.Sprint(sub))
	if result.Error != nil {
		ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"status": "fail", "message": "the user belonging to this token no logger exists"})
		return
	}

	access_token, err := utils.CreateToken(config.AccessTokenExpiresIn, user.ID, config.AccessTokenPrivateKey)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	ctx.SetCookie("access_token", access_token, config.AccessTokenMaxAge*60, "/", config.HOST, false, true)
	ctx.SetCookie("logged_in", "true", config.AccessTokenMaxAge*60, "/", config.HOST, false, false)

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "access_token": access_token})
}

// [...] Logout user
func (ac *AuthController) LogoutUser(ctx *gin.Context) {
	config, _ := initializers.LoadConfig(".")
	ctx.SetCookie("access_token", "", -1, "/", config.HOST, false, true)
	ctx.SetCookie("refresh_token", "", -1, "/", config.HOST, false, true)
	ctx.SetCookie("logged_in", "", -1, "/", config.HOST, false, false)

	ctx.JSON(http.StatusOK, gin.H{"status": "success"})
}

// Change password
func (ac *AuthController) ChangePassword(ctx *gin.Context) {
	currentUser := ctx.MustGet("currentUser").(models.User)
	var payload = struct {
		CurrentPassword string `json:"currentPassword"`
		NewPassword     string `json:"newPassword"`
		ConfirmPassword string `json:"confirmPassword"`
	}{}
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	if err := utils.VerifyPassword(currentUser.Password, payload.CurrentPassword); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "Wrong current password"})
		return
	}

	if payload.CurrentPassword == payload.NewPassword {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "New password cannot be the same as the current password"})
		return
	}

	if payload.NewPassword != payload.ConfirmPassword {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "New password and Confirm password do not match"})
		return
	}

	hashedPassword, err := utils.HashPassword(payload.NewPassword)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": err.Error()})
		return
	}

	updatePassword := &models.User{
		Password: hashedPassword,
	}
	result1 := ac.DB.Model(&currentUser).Updates(updatePassword)
	if result1.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "Can not update password"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": "Change password success"})

}
