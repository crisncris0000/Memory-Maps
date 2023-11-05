package handlers

import "github.com/crisncris0000/Memory-Maps/be-app/internal/models"

type VisibilityHandler struct {
	VisibilityModel *models.MarkerPostImpl
}

func NewVisibilityHandler(vModel *models.MarkerPostImpl) *VisibilityHandler {
	return &VisibilityHandler{VisibilityModel: vModel}
}
