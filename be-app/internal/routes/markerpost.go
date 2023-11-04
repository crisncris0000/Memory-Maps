package routes

import (
	"github.com/crisncris0000/Memory-Maps/be-app/internal/handlers"
	"github.com/gin-gonic/gin"
)

type MarkerPostRouter struct {
	MarkerPostHandler *handlers.MarkerPostHandler
}

func NewMarkerPostRouter(mHandler *handlers.MarkerPostHandler) *MarkerPostRouter {
	return &MarkerPostRouter{MarkerPostHandler: mHandler}
}

func (mHandler *MarkerPostRouter) InitializeRouter(router *gin.Engine) {
	router.GET("/marker-posts", mHandler.MarkerPostHandler.GetAllMarkerPosts)
	router.GET("/marker-posts/filter", mHandler.MarkerPostHandler.FilterByDate)
	router.POST("/marker-posts/new", mHandler.MarkerPostHandler.CreateMarkerPost)
	router.DELETE("/marker-posts/delete/:id", mHandler.MarkerPostHandler.DeletePost)
}
