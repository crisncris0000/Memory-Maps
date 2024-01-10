package handlers

import (
	"math/rand"
	"net/http"
	"net/smtp"
	"strconv"
	"time"

	"github.com/crisncris0000/Memory-Maps/be-app/internal/models"
	"github.com/crisncris0000/Memory-Maps/be-app/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/jordan-wright/email"
)

type ResetTokenHandler struct {
	DB models.ResetTokenImpl
}

func NewResetTokenHandler(db *models.ResetTokenImpl) *ResetTokenHandler {
	return &ResetTokenHandler{DB: *db}
}

type ResetPasswordDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Token    string `json:"token"`
}

func (rt *ResetTokenHandler) ChangeUserPassword(context *gin.Context) {

	var resetPasswordDTO ResetPasswordDTO

	if err := context.ShouldBindJSON(&resetPasswordDTO); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Error binding JSON",
			"error":   err,
		})
		return
	}

	token, err := rt.DB.GetResetToken(resetPasswordDTO.Token, resetPasswordDTO.Email)

	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{
			"message": "Reset Token not found",
			"error":   err,
		})
		return
	}

	uModel := models.NewUserModel(rt.DB.DB)

	hashedPassword, err := utils.HashPassword(resetPasswordDTO.Password)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error hashing password",
			"error":   err,
		})
		return
	}

	err = uModel.UpdateUser(resetPasswordDTO.Email, hashedPassword)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error updating user",
			"error":   err,
		})
		return
	}

	err = rt.DB.DeleteResetToken(token)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error deleting token",
			"error":   err,
		})
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Token retrieved",
		"token":   token,
	})
}

func (rt *ResetTokenHandler) CreateResetToken(context *gin.Context) {
	var resetToken models.ResetToken

	if err := context.ShouldBindJSON(&resetToken); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Error binding to json",
			"error":   err,
		})
		return
	}

	uModel := models.NewUserModel(rt.DB.DB)

	exists, err := uModel.UserExists(resetToken.Email)

	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{
			"message": "User does not exist",
			"error":   err,
		})
		return
	}

	if !exists {
		context.JSON(http.StatusNotFound, gin.H{
			"message": "User does not exist",
		})
		return
	}

	seed := time.Now().UnixNano()
	rng := rand.New(rand.NewSource(seed))
	code := rng.Intn(1000000)

	email := &email.Email{
		To:      []string{resetToken.Email},
		From:    "christopherrivera384@gmail.com",
		Subject: "Reset password request",
		Text:    []byte(strconv.Itoa(code)),
	}

	userEmail := utils.GetValueOfEnvKey("GMAIL_APP_USERNAME")

	password := utils.GetValueOfEnvKey("GMAIL_APP_PASSWORD")

	err = email.Send("smtp.gmail.com:587", smtp.PlainAuth("", userEmail, password, "smtp.gmail.com"))

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error sending the email",
			"error":   err,
		})
		return
	}

	resetToken.Token = strconv.Itoa(code)

	err = rt.DB.CreateResetToken(resetToken)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error creating Reset Token",
			"error":   err,
		})
		return
	}

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error retrieving user by ID",
			"error":   err,
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Please check your email for the code",
	})
}
