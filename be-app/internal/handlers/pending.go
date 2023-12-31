package handlers

import (
	"net/http"
	"strconv"

	"github.com/crisncris0000/Memory-Maps/be-app/internal/models"
	"github.com/gin-gonic/gin"
)

type PendingRequestHandler struct {
	DB *models.PendingRequestModelImpl
}

func NewPendingRequestHandler(pModel *models.PendingRequestModelImpl) *PendingRequestHandler {
	return &PendingRequestHandler{DB: pModel}
}

func (pModel *PendingRequestHandler) GetUserPendingRequests(context *gin.Context) {
	param := context.Param("id")

	id, err := strconv.Atoi(param)

	if err != nil {
		context.JSON(http.StatusNotAcceptable, gin.H{
			"message": "Error converting to integer for Pending Requests",
			"error":   err,
		})
		return
	}

	requests, err := pModel.DB.GetUserPendingRequests(id)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error querying database for Pending Requests",
			"error":   err,
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message":  "successfully retrieved pending requests",
		"requests": requests,
	})
}

func (pModel *PendingRequestHandler) SendFriendRequest(context *gin.Context) {
	var pendingRequest models.PendingRequest

	if err := context.ShouldBindJSON(&pendingRequest); err != nil {
		context.JSON(http.StatusNotAcceptable, gin.H{
			"message": "Error binding to JSON for Friend Request",
			"error":   err,
		})
		return
	}

	err := pModel.DB.SendFriendRequest(pendingRequest)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Friend request sent!",
	})
}

func (pModel *PendingRequestHandler) DeclineFriendRequest(context *gin.Context) {
	param := context.Param("id")

	id, err := strconv.Atoi(param)

	if err != nil {
		context.JSON(http.StatusNotAcceptable, gin.H{"error": err})
		return
	}

	err = pModel.DB.DeclineFriendRequest(id)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error declining friend request",
			"error":   err,
		})
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Successfully declined request",
	})
}
