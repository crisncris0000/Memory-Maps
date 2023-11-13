package routes

import (
	"github.com/crisncris0000/Memory-Maps/be-app/internal/handlers"
	"github.com/gin-gonic/gin"
)

type PendingRequestRouter struct {
	PendingHandler *handlers.PendingRequestHandler
}

func NewPendingRequestRouter(pHandler *handlers.PendingRequestHandler) *PendingRequestRouter {
	return &PendingRequestRouter{PendingHandler: pHandler}
}

func (pRouter *PendingRequestRouter) InitializeRouter(router *gin.Engine) {
	router.GET("/pending-request/:id", pRouter.PendingHandler.GetUserPendingRequests)
	router.POST("/pending-request/send", pRouter.PendingHandler.SendFriendRequest)
	router.DELETE("/pending-request/decline", pRouter.PendingHandler.DeclineFriendRequest)
}
