package routes

import (
	"github.com/crisncris0000/Memory-Maps/be-app/internal/handlers"
	"github.com/gin-gonic/gin"
)

type ResetTokenRouter struct {
	ResetTokenHandler *handlers.ResetTokenHandler
}

func NewResetTokenRouter(resetTokenHandler *handlers.ResetTokenHandler) *ResetTokenRouter {
	return &ResetTokenRouter{ResetTokenHandler: resetTokenHandler}
}

func (rtRouter *ResetTokenRouter) InitializeRouter(router *gin.Engine) {
	router.POST("/reset-token", rtRouter.ResetTokenHandler.ChangeUserPassword)
	router.POST("reset-token/new", rtRouter.ResetTokenHandler.CreateResetToken)
	router.DELETE("/reset-token/delete/:id", rtRouter.ResetTokenHandler.DeleteResetToken)
}
