package handlers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/crisncris0000/Memory-Maps/be-app/internal/models"
	"github.com/crisncris0000/Memory-Maps/be-app/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type UserHandler struct {
	DB *models.UserModelImpl
}

type LoginForm struct {
	Email    string `json:"email"`
	Password string `json:"password"`
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
			"message": "passwords do not match",
			"error":   err,
		})
		return
	}

	token := jwt.New(jwt.SigningMethodEdDSA)

	claims := token.Claims.(jwt.MapClaims)

	claims["email"] = user.Email
	claims["role"] = user.RoleID
	claims["exp"] = time.Now().Add(time.Hour * 24)

	context.JSON(http.StatusAccepted, gin.H{
		"message": "Successfully login user",
		"token":   token,
	})
}

func (uHandler *UserHandler) GetJWTToken(context *gin.Context) {
	email := context.Param("email")

	token := jwt.New(jwt.SigningMethodEdDSA)

	claims := token.Claims.(jwt.MapClaims)

	claims["email"] = email
	claims["role"] = 1
	claims["exp"] = time.Now().Add(time.Hour * 24)

	context.JSON(http.StatusAccepted, gin.H{
		"message": "Successfully login user",
		"token":   token,
	})
}
