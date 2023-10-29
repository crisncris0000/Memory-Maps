package routes

import (
	"github.com/crisncris0000/Memory-Maps/be-app/internal/handlers"
	"github.com/gin-gonic/gin"
)

type UserRouter struct {
	UserHandler *handlers.UserHandler
}

func NewUserRouter(uHandler *handlers.UserHandler) *UserRouter {
	return &UserRouter{UserHandler: uHandler}
}

func (uRouter *UserRouter) InitializeUserRouter(router *gin.Engine) {
	router.GET("/users", uRouter.UserHandler.GetUsers)
	router.POST("/users/new", uRouter.UserHandler.CreateUser)
}
