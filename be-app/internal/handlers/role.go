package handlers

import (
	"github.com/crisncris0000/Memory-Maps/be-app/internal/models"
)

type RoleHandler struct {
	RoleModel *models.RoleModelImpl
}

func NewRoleHandler(rModel *models.RoleModelImpl) *RoleHandler {
	return &RoleHandler{RoleModel: rModel}
}
