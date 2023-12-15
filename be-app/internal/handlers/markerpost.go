package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/crisncris0000/Memory-Maps/be-app/internal/models"
	"github.com/gin-gonic/gin"
)

type MarkerPostHandler struct {
	DB *models.MarkerPostImpl
}

type DateRange struct {
	StartDate string `json:"startDate"`
	EndDate   string `json:"endDate"`
}

type MarkerPostDTO struct {
	Lattitude    float32 `json:"lattitude"`
	Longitude    float32 `json:"longitude"`
	Description  string  `json:"description"`
	Image        []byte  `json:"image"`
	MimeType     string  `json:"mimeType"`
	Likes        int     `json:"likes"`
	VisibilityID int     `json:"visibilityID"`
	UserEmail    string  `json:"userEmail"`
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
	var markerPostDTO MarkerPostDTO

	if err := context.ShouldBindJSON(&markerPostDTO); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "error binding json of marker post DTO",
			"error":   err,
		})
		fmt.Println(err)
		return
	}

	markerID, err := mHandler.DB.CreateMarkerPost(models.MarkerPost{
		Lattitude:    markerPostDTO.Lattitude,
		Longitude:    markerPostDTO.Longitude,
		Description:  markerPostDTO.Description,
		Likes:        markerPostDTO.Likes,
		VisibilityID: 1,
		UserID:       1,
	})

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "error creating marker post with given data",
			"error":   err,
		})
		return
	}

	markerPostImage := models.MarkerPostImage{
		Image:    markerPostDTO.Image,
		MimeType: markerPostDTO.MimeType,
		MarkerID: int(markerID),
	}

	iModel := models.NewMarkerPostImageModel(mHandler.DB.DB)

	err = iModel.CreateSingleImage(markerPostImage)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "error creating marker post image with given data",
			"error":   err,
		})
	}

	context.JSON(http.StatusInternalServerError, gin.H{
		"message": "marker post created",
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
