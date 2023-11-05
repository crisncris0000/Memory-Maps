package handlers

import (
	"net/http"
	"strconv"

	"github.com/crisncris0000/Memory-Maps/be-app/internal/models"
	"github.com/gin-gonic/gin"
)

type VisibilityHandler struct {
	VisibilityModel *models.VisibilityImpl
}

func NewVisibilityHandler(vModel *models.VisibilityImpl) *VisibilityHandler {
	return &VisibilityHandler{VisibilityModel: vModel}
}

func (vHandler *VisibilityHandler) GetVisibilityName(context *gin.Context) {

	param := context.Param("id")

	id, err := strconv.Atoi(param)

	if err != nil {
		context.JSON(http.StatusNotAcceptable, gin.H{"error": err.Error()})
		return
	}

	visibilityName, err := vHandler.VisibilityModel.GetVisibilityByID(id)
	if err != nil {
		context.JSON(500, gin.H{"error": err.Error()})
		return
	}
	context.JSON(200, gin.H{"visibility": visibilityName})

}
