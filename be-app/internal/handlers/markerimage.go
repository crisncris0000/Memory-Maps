package handlers

import (
	"net/http"
	"strconv"

	"github.com/crisncris0000/Memory-Maps/be-app/internal/models"
	"github.com/gin-gonic/gin"
)

type MarkerPostIamgeHandler struct {
	DB *models.MarkerPostImageImpl
}

func NewMarkerPostImageHandler(db *models.MarkerPostImageImpl) *MarkerPostIamgeHandler {
	return &MarkerPostIamgeHandler{DB: db}
}

func (iHandler *MarkerPostIamgeHandler) GetMarkerPostImages(context *gin.Context) {
	param := context.Param("id")

	id, err := strconv.Atoi(param)

	if err != nil {
		context.JSON(http.StatusNotAcceptable, gin.H{
			"message": "Cannot convert param to integer",
			"error":   err,
		})
		return
	}

	markerPostImages, err := iHandler.DB.GetMarkerPostImages(id)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Database error retrieving Images",
			"error":   err,
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Successfully retrieved images",
		"images":  markerPostImages,
	})
}
