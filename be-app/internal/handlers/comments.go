package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/crisncris0000/Memory-Maps/be-app/internal/models"
	"github.com/gin-gonic/gin"
)

type CommentsHandler struct {
	DB *models.CommentsModelImpl
}

func NewCommentsHandler(cModel *models.CommentsModelImpl) *CommentsHandler {
	return &CommentsHandler{DB: cModel}
}

func (cHandler *CommentsHandler) GetAllComments(context *gin.Context) {
	comments, err := cHandler.DB.GetAllComments()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err})
	}

	context.JSON(http.StatusOK, comments)
}

func (cHandler *CommentsHandler) GetAllCommentsByMarkerPostID(context *gin.Context) {
	param := context.Param("id")

	id, err := strconv.Atoi(param)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err})
	}

	comments, err := cHandler.DB.GetCommentsByMarkerID(id)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err})
	}

	context.JSON(http.StatusOK, comments)
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

	err := cHandler.DB.UpdateComment(comment.Comment)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "error updating comment " + err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"Message": "Comment has been updated"})
}

func (cHandler *CommentsHandler) DeleteComment(context *gin.Context) {
	param := context.Param("id")

	id, err := strconv.Atoi(param)

	if err != nil {
		context.JSON(http.StatusNotAcceptable, gin.H{"error": err})
		return
	}

	err = cHandler.DB.DeleteCommentByID(id)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	context.JSON(http.StatusOK, gin.H{"Message": "Successfully deleted comment"})
}
