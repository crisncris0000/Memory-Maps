package routes

import (
	"github.com/crisncris0000/Memory-Maps/be-app/internal/handlers"
	"github.com/gin-gonic/gin"
)

type CommentsRouter struct {
	CommentsHandler *handlers.CommentsHandler
}

func NewCommentsRouter(cHandler *handlers.CommentsHandler) *CommentsRouter {
	return &CommentsRouter{CommentsHandler: cHandler}
}

func (cRouter *CommentsRouter) InitializeRouter(router *gin.Engine) {
	router.GET("/comments", cRouter.CommentsHandler.GetAllComments)
	router.GET("/comments/:id", cRouter.CommentsHandler.GetAllCommentsByMarkerPostID)
	router.POST("/comments/new", cRouter.CommentsHandler.CreateComment)
	router.PUT("/comments/edit", cRouter.CommentsHandler.UpdateComment)
	router.DELETE("/comments/delete/:id", cRouter.CommentsHandler.DeleteComment)
}
