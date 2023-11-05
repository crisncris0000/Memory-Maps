package routes

import (
	"github.com/crisncris0000/Memory-Maps/be-app/internal/handlers"
	"github.com/gin-gonic/gin"
)

type VisibilityRouter struct {
	vHandler *handlers.VisibilityHandler
}

func NewVisibilityRouter(vHandler *handlers.VisibilityHandler) *VisibilityRouter {
	return &VisibilityRouter{vHandler: vHandler}
}

func (vRouter *VisibilityRouter) InitializeRouter(router *gin.Engine) {

}
