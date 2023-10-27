package handlers

import "github.com/crisncris0000/Memory-Maps/be-app/internal/models"

type UserHandler struct {
	db *model.UserModelImpl
}

func newUserHandler(uModelImpl *model.UserModelImpl) *UserHandler {
	return &UserHandler{db: uModelImpl}
}
