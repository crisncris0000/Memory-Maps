package handlers

import (
	"net/http"
	"strconv"

	"github.com/crisncris0000/Memory-Maps/be-app/internal/models"
	"github.com/gin-gonic/gin"
)

type TagsHandler struct {
	DB *models.TagsModelImpl
}

func NewTagsHandler(db *models.TagsModelImpl) *TagsHandler {
	return &TagsHandler{DB: db}
}

func (tHandler *TagsHandler) CreateTag(context *gin.Context) {
	tagName := context.PostForm("tagName")

	err := tHandler.DB.CreateTag(tagName)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failure to create Tag within database",
			"error":   err,
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Successfully created tag",
	})
}

func (tHandler *TagsHandler) DeleteTag(context *gin.Context) {
	param := context.Param("id")

	id, err := strconv.Atoi(param)

	if err != nil {
		context.JSON(http.StatusNotAcceptable, gin.H{
			"message": "Failure to convert tag ID to integer",
			"error":   err,
		})
		return
	}

	err = tHandler.DB.DeleteTag(id)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error querying database to delete Tag",
			"error":   err,
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Successfully deleted tag",
	})
}
