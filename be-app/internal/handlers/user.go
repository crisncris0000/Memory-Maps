package handlers

import (
	"fmt"
	"net/http"
	"net/smtp"
	"strconv"

	"github.com/crisncris0000/Memory-Maps/be-app/internal/models"
	"github.com/crisncris0000/Memory-Maps/be-app/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/jordan-wright/email"
)

type UserHandler struct {
	DB *models.UserModelImpl
}

type LoginForm struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SendEmail struct {
	Subject string `json:"subject"`
	Email   string `json:"email"`
	Body    string `json:"body"`
}

func NewUserHandler(uModelImpl *models.UserModelImpl) *UserHandler {
	return &UserHandler{DB: uModelImpl}
}

func (uHandler *UserHandler) GetUsers(context *gin.Context) {

	users, err := uHandler.DB.GetUsers()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error retrieving all users",
			"error":   err,
		})
		return
	}

	context.JSON(http.StatusOK, users)
}

func (uHandler *UserHandler) GetUserByID(context *gin.Context) {

	param := context.Param("id")

	id, err := strconv.Atoi(param)

	if err != nil {
		context.JSON(http.StatusNotAcceptable, gin.H{
			"message": "Cannot convert to int",
			"error":   err,
		})
		return
	}

	user, err := uHandler.DB.GetUserByID(id)

	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{
			"message": "Error querying database with users ID",
			"error":   err,
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Successfully retrieved user",
		"user":    user,
	})
}

func (uHandler *UserHandler) CreateUser(context *gin.Context) {
	var user models.User

	if err := context.ShouldBindJSON(&user); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error binding to JSON",
			"error":   err,
		})
		return
	}

	user.RoleID = 1

	err := uHandler.DB.CreateUser(user)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error creating user",
			"error":   err,
		})
		return
	}

	context.JSON(http.StatusCreated, gin.H{
		"message": "Successfully created user",
	})
}

func (uHandler *UserHandler) AuthenticateUser(context *gin.Context) {

	var loginForm LoginForm

	if err := context.ShouldBindJSON(&loginForm); err != nil {
		context.JSON(http.StatusNotAcceptable, gin.H{
			"message": "JSON Format not acceptable",
			"error":   err,
		})
		return
	}

	user, err := uHandler.DB.GetUserByEmail(loginForm.Email)

	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{
			"message": "Cannot find user's email",
			"error":   err,
		})
		return
	}

	matches := utils.ComparePasswords(user.Password, loginForm.Password)

	if !matches {
		context.JSON(http.StatusNotAcceptable, gin.H{
			"message": "Passwords do not match",
			"error":   err,
		})
		return
	}

	token, err := utils.GenerateJWTToken(user.FirstName, user.LastName, user.Email, user.ID, user.RoleID)

	if err != nil {
		context.JSON(http.StatusNotAcceptable, gin.H{
			"message": "Error retrieving JWT",
			"error":   err,
		})
		return
	}

	context.JSON(http.StatusAccepted, gin.H{
		"message": "Successfully login user",
		"token":   token,
	})
}

func (uHandler *UserHandler) SendEmail(context *gin.Context) {

	var emailMessage SendEmail

	if err := context.ShouldBindJSON(&emailMessage); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not bind to JSON",
			"error":   err,
		})
		return
	}

	email := &email.Email{
		To:      []string{"Christopherrivera384@gmail.com"},
		From:    emailMessage.Email,
		Subject: emailMessage.Subject,
		Text:    []byte(emailMessage.Body),
	}

	userEmail := utils.GetValueOfEnvKey("GMAIL_APP_USERNAME")

	password := utils.GetValueOfEnvKey("GMAIL_APP_PASSWORD")

	err := email.Send("smtp.gmail.com:587", smtp.PlainAuth("", userEmail, password, "smtp.gmail.com"))

	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not send email",
			"error":   err,
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Email sent successfully",
	})
}

func (uHandler *UserHandler) ResetPassword(context *gin.Context) {

}
