package routes

import (
	"github.com/crisncris0000/Memory-Maps/be-app/internal/handlers"
	"github.com/gin-gonic/gin"
)

type FriendsWithRouter struct {
	FriendsWithHandler *handlers.FriendsWithHandler
}

func NewFriendsWithRouter(fHandler *handlers.FriendsWithHandler) *FriendsWithRouter {
	return &FriendsWithRouter{FriendsWithHandler: fHandler}
}

func (fRouter *FriendsWithRouter) InitializeRouter(router *gin.Engine) {
	router.GET("/friends-with/:id", fRouter.FriendsWithHandler.GetUserFriends)
	router.DELETE("/friends-with/remove/:id", fRouter.FriendsWithHandler.DeleteFriend)
}
