package handlers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/crisncris0000/Memory-Maps/be-app/internal/models"
	"github.com/crisncris0000/Memory-Maps/be-app/internal/utils"
	"github.com/gin-gonic/gin"
)

type MarkerPostHandler struct {
	DB *models.MarkerPostImpl
}

type DateRange struct {
	StartDate string `json:"startDate"`
	EndDate   string `json:"endDate"`
}

func NewMarkerPostHandler(mModel *models.MarkerPostImpl) *MarkerPostHandler {
	return &MarkerPostHandler{DB: mModel}
}

func (mHandler *MarkerPostHandler) GetAllMarkerPosts(context *gin.Context) {
	posts, err := mHandler.DB.GetMarkerPosts()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "error retrieving marker posts",
			"error":   err,
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{"markerposts": posts})
}

func (mHandler *MarkerPostHandler) CreateMarkerPost(context *gin.Context) {
	latitudeStr := context.PostForm("latitude")
	longitudeStr := context.PostForm("longitude")
	imageFile, _ := context.FormFile("image")
	description := context.PostForm("description")
	visibilityIDStr := context.PostForm("visibilityID")
	userIDStr := context.PostForm("userID")

	latitude, longitude, image, description, visibilityID, userID, err :=
		utils.HandleMarkerPostConversion(latitudeStr, longitudeStr, userIDStr, description, visibilityIDStr, imageFile)

	if err != nil {
		context.JSON(http.StatusNotAcceptable, gin.H{
			"message": "Error converting when creating marker post",
			"error":   err,
		})
		return
	}

	marker := models.MarkerPost{
		Lattitude:    latitude,
		Longitude:    longitude,
		Image:        image,
		Description:  description,
		Likes:        0,
		VisibilityID: visibilityID,
		UserID:       userID,
	}

	err = mHandler.DB.CreateMarkerPost(marker)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error querying database for creating marker post",
			"error":   err,
		})
		return
	}

	context.JSON(http.StatusCreated, gin.H{
		"message": "Successfully created marker post",
	})
}

func (mHandler *MarkerPostHandler) FilterByDate(context *gin.Context) {

	var dateRange DateRange

	if err := context.BindJSON(&dateRange); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "error binding to JSON for date range",
			"error":   err,
		})
		return
	}

	startDate, err := time.Parse("2006-01-01", dateRange.StartDate)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error parsing start date time",
			"error":   err,
		})
		return
	}

	endDate, err := time.Parse("2006-01-01", dateRange.EndDate)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error parsing end date time",
			"error":   err,
		})
		return
	}

	posts, err := mHandler.DB.GetPostsByDate(startDate, endDate)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "error querying database to filter posts by date",
			"error":   err,
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "successfully retrieved posts",
		"posts":   posts,
	})
}

func (mHandler *MarkerPostHandler) UpdatePost(context *gin.Context) {
	var marker models.MarkerPost

	if err := context.BindJSON(&marker); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error binding JSON for updating marker",
			"error":   err,
		})
		return
	}

	err := mHandler.DB.UpdatePost(marker)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error querying database to update marker",
			"error":   err,
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{"success": marker})
}

func (mHandler *MarkerPostHandler) DeletePost(context *gin.Context) {
	param := context.Param("id")

	id, err := strconv.Atoi(param)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error converting to integer for post ID",
			"error":   err,
		})
		return
	}

	err = mHandler.DB.DeletePost(id)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error deleting post by ID",
			"error":   err,
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Post deleted",
	})

}
