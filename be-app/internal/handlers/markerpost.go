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
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	context.JSON(http.StatusOK, gin.H{"Success": posts})
}

func (mHandler *MarkerPostHandler) CreateMarkerPost(context *gin.Context) {
	latitudeStr := context.PostForm("latitude")
	longitudeStr := context.PostForm("longitude")
	imageFile, _ := context.FormFile("image")
	description := context.PostForm("latitude")
	likesStr := context.PostForm("latitude")
	visibilityIDStr := context.PostForm("latitude")
	userIDStr := context.PostForm("latitude")

	latitude, longitude, image, description, likes, visibilityID, userID, err :=
		utils.HandleMarkerPostConversion(latitudeStr, longitudeStr, userIDStr, description, likesStr, visibilityIDStr, imageFile)

	if err != nil {
		context.JSON(http.StatusNotAcceptable, gin.H{"error": err})
	}

	marker := models.MarkerPost{
		Lattitude:    latitude,
		Longitude:    longitude,
		Image:        image,
		Description:  description,
		Likes:        likes,
		VisibilityID: visibilityID,
		UserID:       userID,
	}

	err = mHandler.DB.CreateMarkerPost(marker)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"success": marker})
}

func (mHandler *MarkerPostHandler) FilterByDate(context *gin.Context) {

	var dateRange DateRange

	if err := context.BindJSON(&dateRange); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err})
	}

	startDate, err := time.Parse("2006-01-01", dateRange.StartDate)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err})
	}

	endDate, err := time.Parse("2006-01-01", dateRange.EndDate)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err})
	}

	posts, err := mHandler.DB.GetPostsByDate(startDate, endDate)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err})
	}

	context.JSON(http.StatusOK, gin.H{"success": posts})
}

func (mHandler *MarkerPostHandler) UpdatePost(context *gin.Context) {
	var marker models.MarkerPost

	if err := context.BindJSON(&marker); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err})
	}

	err := mHandler.DB.UpdatePost(marker)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err})
	}

	context.JSON(http.StatusOK, gin.H{"success": marker})
}

func (mHandler *MarkerPostHandler) DeletePost(context *gin.Context) {
	param := context.Param("id")

	id, err := strconv.Atoi(param)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err})
	}

	err = mHandler.DB.DeletePost(id)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err})
	}

	context.JSON(http.StatusOK, gin.H{"success": "Post deleted"})

}
