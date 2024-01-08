package handler

import (
	"YTStreamGoApi/config"
	"YTStreamGoApi/user"
	"YTStreamGoApi/validator"
)

type Handler struct {
	userStore user.Store
	validator *validator.Validator
	config    *config.Config
}

func NewHandler(userStore user.Store, config *config.Config) *Handler {
	return &Handler{
		userStore: userStore,
		validator: validator.NewValidator(),
		config:    config,
	}
}
