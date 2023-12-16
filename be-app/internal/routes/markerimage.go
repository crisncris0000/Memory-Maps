package routes

import (
	"github.com/crisncris0000/Memory-Maps/be-app/internal/handlers"
	"github.com/gin-gonic/gin"
)

type MarkerPostImageRouter struct {
	MarkerImageHandler *handlers.MarkerPostIamgeHandler
}

func NewMarkerPostImageRouter(iHandler *handlers.MarkerPostIamgeHandler) *MarkerPostImageRouter {
	return &MarkerPostImageRouter{MarkerImageHandler: iHandler}
}

func (iRouter *MarkerPostImageRouter) InitializeRoutes(router *gin.Engine) {
	router.GET("/marker-post/images/:id", iRouter.MarkerImageHandler.GetMarkerPostImages)
}
