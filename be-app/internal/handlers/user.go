package handlers

import (
	"fmt"
	"net/http"

	"github.com/crisncris0000/Memory-Maps/be-app/internal/models"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	UserModel *models.UserModelImpl
}

func NewUserHandler(uModelImpl *models.UserModelImpl) *UserHandler {
	return &UserHandler{UserModel: uModelImpl}
}

func (uHandler *UserHandler) GetUsers(context *gin.Context) {

	users, err := uHandler.UserModel.GetUsers()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, users)
}

func (uHandler *UserHandler) CreateUser(context *gin.Context) {
	var user models.User

	if err := context.ShouldBindJSON(&user); err != nil {
		fmt.Println("Error binding the JSON of user", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err := uHandler.UserModel.CreateUser(user)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"success": user})
}
