package routes

import (
	"github.com/crisncris0000/Memory-Maps/be-app/internal/handlers"
	"github.com/gin-gonic/gin"
)

type RoleRouter struct {
	*handlers.RoleHandler
}

func NewRoleRouter(rHandler *handlers.RoleHandler) *RoleRouter {
	return &RoleRouter{RoleHandler: rHandler}
}

func (rRouter *RoleRouter) InitializeRouter(router *gin.Engine) {

}
