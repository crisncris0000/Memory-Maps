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
	router.GET("/users/:id", uRouter.UserHandler.GetUserByID)
	router.POST("/users/new", uRouter.UserHandler.CreateUser)
	router.POST("/users/login", uRouter.UserHandler.AuthenticateUser)
	router.POST("/users/send-email", uRouter.UserHandler.SendEmail)
}
