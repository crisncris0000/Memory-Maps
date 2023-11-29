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
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error retrieving all comments",
			"error":   err,
		})
	}

	context.JSON(http.StatusOK, comments)
}

func (cHandler *CommentsHandler) GetAllCommentsByMarkerPostID(context *gin.Context) {
	param := context.Param("id")

	id, err := strconv.Atoi(param)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error converting to integer",
			"error":   err,
		})
	}

	comments, err := cHandler.DB.GetCommentsByMarkerID(id)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error querying database get comments by marker ID",
			"error":   err,
		})
	}

	context.JSON(http.StatusOK, gin.H{
		"message":  "Successfully retrieved comments for marker post",
		"comments": comments,
	})
}

func (cHandler *CommentsHandler) CreateComment(context *gin.Context) {
	var comment models.Comments

	if err := context.ShouldBindJSON(&comment); err != nil {
		context.JSON(http.StatusNotAcceptable, gin.H{
			"message": "Error binding JSON for Comment",
			"error":   err,
		})
		return
	}

	err := cHandler.DB.CreateComment(comment)

	if err != nil {
		fmt.Println("Error Creating Comment", err)
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Erorr querying database for creating comment",
			"error":   err,
		})
		return
	}

	context.JSON(http.StatusAccepted, gin.H{"message": "Comment has been posted"})
}

func (cHandler *CommentsHandler) UpdateComment(context *gin.Context) {
	var comment models.Comments

	if err := context.ShouldBindJSON(&comment); err != nil {
		context.JSON(http.StatusNotAcceptable, gin.H{
			"message": "Error binding json for comment",
			"error":   err,
		})
		return
	}

	err := cHandler.DB.UpdateComment(comment.Comment)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error updating comment",
			"error":   err,
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Comment has been updated",
	})
}

func (cHandler *CommentsHandler) DeleteComment(context *gin.Context) {
	param := context.Param("id")

	id, err := strconv.Atoi(param)

	if err != nil {
		context.JSON(http.StatusNotAcceptable, gin.H{
			"message": "Error converting id to integer",
			"error":   err,
		})
		return
	}

	err = cHandler.DB.DeleteCommentByID(id)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error deleting comment",
			"error":   err,
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Successfully deleted comment",
	})
}
