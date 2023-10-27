package handlers

import (
	"github.com/crisncris0000/Memory-Maps/be-app/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserHandler struct {
	db *models.UserModelImpl
}

func NewUserHandler(uModelImpl *models.UserModelImpl) *UserHandler {
	return &UserHandler{db: uModelImpl}
}

func (uHandler *UserHandler) GetUsers(context *gin.Context) {

	users, err := uHandler.db.GetUsers()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, users)

}
