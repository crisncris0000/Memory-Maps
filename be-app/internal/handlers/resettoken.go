package handlers

import (
	"net/http"

	"github.com/crisncris0000/Memory-Maps/be-app/internal/models"
	"github.com/gin-gonic/gin"
)

type ResetTokenHandler struct {
	DB models.ResetTokenImpl
}

func NewResetTokenHandler(db *models.ResetTokenImpl) *ResetTokenHandler {
	return &ResetTokenHandler{DB: *db}
}

func (rt *ResetTokenHandler) GetResetToken(context *gin.Context) {

	var resetToken models.ResetToken

	if err := context.ShouldBindJSON(&resetToken); err != nil {
		context.JSON(http.StatusNotFound, gin.H{
			"message": "Error binding JSON",
			"error":   err,
		})
		return
	}

	token, err := rt.DB.GetResetToken(resetToken.Token, resetToken.UserID)

	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{
			"message": "Reset Token not found",
			"error":   err,
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Token retrieved",
		"token":   token,
	})
}
