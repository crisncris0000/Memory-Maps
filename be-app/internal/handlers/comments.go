package handlers

import (
	"fmt"
	"net/http"

	"github.com/crisncris0000/Memory-Maps/be-app/internal/models"
	"github.com/gin-gonic/gin"
)

type CommentsHandler struct {
	DB *models.CommentsModelImpl
}

func NewCommentsModel(cModel *models.CommentsModelImpl) *CommentsHandler {
	return &CommentsHandler{DB: cModel}
}

func (cHandler *CommentsHandler) CreateComment(context *gin.Context) {
	var comment models.Comments

	if err := context.ShouldBindJSON(&comment); err != nil {
		fmt.Println("Error binding JSON", err)
		context.JSON(http.StatusNotAcceptable, gin.H{"error": err})
		return
	}

	err := cHandler.DB.CreateComment(comment)

	if err != nil {
		fmt.Println("Error Creating Comment", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	context.JSON(http.StatusAccepted, gin.H{"Message": "Comment has been posted"})
}

func (cHandler *CommentsHandler) UpdateComment(context *gin.Context) {
	var comment models.Comments

	if err := context.ShouldBindJSON(&comment); err != nil {
		context.JSON(http.StatusNotAcceptable, gin.H{"error": "error binding json " + err.Error()})
		return
	}

}
