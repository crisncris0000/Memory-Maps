package handlers

import (
	"net/http"
	"strconv"

	"github.com/crisncris0000/Memory-Maps/be-app/internal/models"
	"github.com/gin-gonic/gin"
)

type RoleHandler struct {
	DB *models.RoleModelImpl
}

func NewRoleHandler(rModel *models.RoleModelImpl) *RoleHandler {
	return &RoleHandler{DB: rModel}
}

func (rHandler *RoleHandler) GetRole(context *gin.Context) {
	id, err := strconv.Atoi(context.Param("id"))

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	role, err := rHandler.DB.GetRole(id)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	context.JSON(http.StatusOK, role)
}
