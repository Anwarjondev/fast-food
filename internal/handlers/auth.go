package handlers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Anwarjondev/fast-food/config"
	"github.com/Anwarjondev/fast-food/internal/db"
	"github.com/Anwarjondev/fast-food/internal/repository"
	"github.com/Anwarjondev/fast-food/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var appConfig config.Config

func SetConfig(c config.Config) {
	appConfig = c
}

// RegisterRequest represents the request body for user registration
type RegisterRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// ConfirmRequest represents the request body for user confirmation
type ConfirmRequest struct {
	Code int `json:"code"`
}

// LoginRequest represents the request body for user login
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}


// ForgotPasswordRequest represents the request body for password reset request
type ForgotPasswordRequest struct {
	Email string `json:"email"`
}

// ResetPasswordRequest represents the request body for password reset
type ResetPasswordRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// ResendCodeRequest represents the request body for resending confirmation code
type ResendCodeRequest struct {
	Email string `json:"email"`
}

// Register godoc
// @Summary Register a new user
// @Description Register a new user with email and password
// @Tags auth
// @Accept json
// @Produce json
// @Param user body RegisterRequest true "User registration info"
// @Success 200 {object} map[string]interface{}
// @Router /register [post]
func Register(c *gin.Context) {
	var req RegisterRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userID, err := repository.CreateUser(req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	code := utils.GenerateCode()
	repository.SaveConfirmaion(userID, code)

	// Log the code for debugging
	fmt.Printf("Generated code for user %d: %d\n", userID, code)

	err = utils.SendEmailCode(req.Email, code, appConfig.SMPTHost, appConfig.SMTPPort, appConfig.EMAILSender, appConfig.EMAILPassword)
	if err != nil {
		fmt.Printf("Failed to send email: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to send confirmation email: %v", err)})
		return
	}

	go func() {
		time.Sleep(time.Minute)
		repository.MarkCodeAsPassed(userID)
	}()

	c.JSON(http.StatusOK, gin.H{"message": "Code sent to email"})
}

// Confirm godoc
// @Summary Confirm user registration
// @Description Confirm user registration with code
// @Tags auth
// @Accept json
// @Produce json
// @Param request body ConfirmRequest true "Confirmation request"
// @Success 200 {object} map[string]interface{}
// @Router /confirm [post]
func Confirm(c *gin.Context) {
	var req ConfirmRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Find user ID by code
	userID, err := repository.GetUserIDByCode(req.Code)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	ok, err := repository.CheckCode(userID, req.Code)
	if err != nil || !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	repository.MarkCodeAsPassed(userID)
	repository.ActiveUser(userID)
	c.JSON(http.StatusOK, gin.H{"message": "account verified"})
}

// Login godoc
// @Summary User login
// @Description Login with email and password
// @Tags auth
// @Accept json
// @Produce json
// @Param request body LoginRequest true "Login credentials"
// @Success 200 {object} map[string]interface{}
// @Router /login [post]
func Login(c *gin.Context) {
	var req LoginRequest
	token := uuid.New().String()

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	UserID, err := repository.LoginUser(req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	err = repository.SaveToken(UserID, token)
	c.JSON(http.StatusOK, gin.H{"message": "login successful", "Token": token})
}

// Logout godoc
// @Summary User logout
// @Description Logout user and invalidate token
// @Tags auth
// @Security BearerAuth
// @Success 200 {object} map[string]interface{}
// @Router /logout [post]
func Logout(c *gin.Context) {
	// Get user ID from the context (set by AuthMiddleware)
	userID := c.GetInt("user_id")
	if userID == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	err := repository.LogoutUser(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "logout successful"})
}

// ForgotPassword godoc
// @Summary Forgot password
// @Description Request password reset
// @Tags auth
// @Accept json
// @Produce json
// @Param request body ForgotPasswordRequest true "Forgot password request"
// @Success 200 {object} map[string]interface{}
// @Router /forgot-password [post]
func ForgotPassword(c *gin.Context) {
	var req ForgotPasswordRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var userID int
	err := db.DB.Get(&userID, `Select id from users where email = $1`, req.Email)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	code := utils.GenerateCode()
	err = utils.SendEmailCode(req.Email, code, appConfig.SMPTHost, appConfig.SMTPPort, appConfig.EMAILSender, appConfig.EMAILPassword)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	err = repository.SaveConfirmaion(userID, code)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "New code sent to email"})
}

// ResetPassword godoc
// @Summary Reset password
// @Description Reset password with code
// @Tags auth
// @Accept json
// @Produce json
// @Param request body ResetPasswordRequest true "Reset password request"
// @Success 200 {object} map[string]interface{}
// @Router /reset-password [post]
func ResetPassword(c *gin.Context) {
	var req ResetPasswordRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get user ID from email
	var userID int
	err := db.DB.Get(&userID, `Select id from users where email = $1`, req.Email)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// Update the password
	err = repository.UpdatePassword(userID, req.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update password"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Password reset successful"})
}

// ResendCode godoc
// @Summary Resend confirmation code
// @Description Resend confirmation code to user's email
// @Tags auth
// @Accept json
// @Produce json
// @Param request body ResendCodeRequest true "Resend code request"
// @Success 200 {object} map[string]interface{}
// @Router /resend-code [post]
func ResendCode(c *gin.Context) {
	var req ResendCodeRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get user ID from email
	var userID int
	err := db.DB.Get(&userID, `Select id from users where email = $1`, req.Email)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	code := utils.GenerateCode()
	err = repository.SaveConfirmaion(userID, code)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	err = utils.SendEmailCode(req.Email, code, appConfig.SMPTHost, appConfig.SMTPPort, appConfig.EMAILSender, appConfig.EMAILPassword)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "New code sent to email"})
}
