package handlers

import (
	"net/http"
	"strconv"

	"github.com/crisncris0000/Memory-Maps/be-app/internal/models"
	"github.com/gin-gonic/gin"
)

type VisibilityHandler struct {
	DB *models.VisibilityImpl
}

func NewVisibilityHandler(vModel *models.VisibilityImpl) *VisibilityHandler {
	return &VisibilityHandler{DB: vModel}
}

func (vHandler *VisibilityHandler) GetVisibilityByID(context *gin.Context) {

	param := context.Param("id")

	id, err := strconv.Atoi(param)

	if err != nil {
		context.JSON(http.StatusNotAcceptable, gin.H{
			"messsage": "Error converting to integer",
			"error":    err,
		})
		return
	}

	visibilityName, err := vHandler.DB.GetVisibilityByID(id)
	if err != nil {
		context.JSON(500, gin.H{
			"message": "Error querying database for visibility ID",
			"error":   err,
		})
		return
	}
	context.JSON(200, gin.H{"visibility": visibilityName})

}
