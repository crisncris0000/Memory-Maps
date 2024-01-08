package handlers

import (
	"net/http"
	"net/smtp"
	"strconv"

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

func (rt *ResetTokenHandler) GetResetToken(context *gin.Context) {

	var resetToken models.ResetToken

	if err := context.ShouldBindJSON(&resetToken); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Error binding JSON",
			"error":   err,
		})
		return
	}

	token, err := rt.DB.GetResetToken(resetToken.Token, resetToken.UserID)

	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{
			"message": "Reset Token not found",
			"error":   err,
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Token retrieved",
		"token":   token,
	})
}

func (rt *ResetTokenHandler) CreateResetToken(context *gin.Context) {
	var resetToken models.ResetToken

	if err := context.ShouldBindHeader(&resetToken); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Error binding to json",
			"error":   err,
		})
		return
	}

	err := rt.DB.CreateResetToken(resetToken)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error creating Reset Token",
			"error":   err,
		})
		return
	}

	userModel := models.NewUserModel(rt.DB.DB)

	user, err := userModel.GetUserByID(resetToken.UserID)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error retrieving user by ID",
			"error":   err,
		})
		return
	}

	email := &email.Email{
		To:      []string{user.Email},
		From:    "christopherrivera384@gmail.com",
		Subject: "Reset password request",
		Text:    []byte(""),
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

	context.JSON(http.StatusOK, gin.H{
		"message": "Please check your email for the code",
	})
}

func (rt *ResetTokenHandler) DeleteResetToken(context *gin.Context) {
	param := context.Param("id")

	id, err := strconv.Atoi(param)

	if err != nil {
		context.JSON(http.StatusNotAcceptable, gin.H{
			"message": "Cannot convert param to integer",
			"error":   err,
		})
	}

	err = rt.DB.DeleteResetToken(id)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error deleting Reset Token",
			"error":   err,
		})
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Successfully removed token",
	})
}
