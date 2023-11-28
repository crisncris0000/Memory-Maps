package handlers

import (
	"net/http"
	"strconv"

	"github.com/crisncris0000/Memory-Maps/be-app/internal/models"
	"github.com/gin-gonic/gin"
)

type FriendsWithHandler struct {
	DB *models.FriendsWithImpl
}

func NewFriendsWithHandler(db *models.FriendsWithImpl) *FriendsWithHandler {
	return &FriendsWithHandler{DB: db}
}

func (fHandler *FriendsWithHandler) GetUserFriends(context *gin.Context) {
	param := context.Param("id")

	id, err := strconv.Atoi(param)

	if err != nil {
		context.JSON(http.StatusNotAcceptable, gin.H{
			"message": "Error converting to integer",
			"error":   err,
		})
		return
	}

	friends, err := fHandler.DB.GetUserFriends(id)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error querying database to retrieve users friends",
			"error":   err,
		})
	}

	context.JSON(http.StatusOK, friends)
}

func (fHandler *FriendsWithHandler) DeleteFriend(context *gin.Context) {
	param := context.Param("id")

	id, err := strconv.Atoi(param)

	if err != nil {
		context.JSON(http.StatusNotAcceptable, gin.H{
			"message": "Error converting to integer",
			"error":   err,
		})
		return
	}

	err = fHandler.DB.DeleteFriend(id)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error querying database for deleting friend",
			"error":   err,
		})
	}

	context.JSON(http.StatusOK, gin.H{"message": "Deleted Friend"})
}
