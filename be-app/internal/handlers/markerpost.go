package handlers

import (
	"fmt"
	"net/http"

	"github.com/crisncris0000/Memory-Maps/be-app/internal/models"
	"github.com/gin-gonic/gin"
)

type MarkerPostHandler struct {
	DB *models.MarkerPostImpl
}

func NewMarkerPostHandler(db *models.MarkerPostImpl) *MarkerPostHandler {
	return &MarkerPostHandler{DB: db}
}

func (mHandler *MarkerPostHandler) CreateMarkerPost(context *gin.Context) {

	var marker models.MarkerPost

	if err := context.BindJSON(&marker); err != nil {
		fmt.Println("Error binding json of marker post", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err := mHandler.DB.CreateMarkerPost(marker)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"success": marker})
}
