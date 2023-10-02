package home

import (
	"github.com/ideal-forward/assistant-public-api/services"
)

type Handler struct {
	PrivateTask services.IPrivateTask
	PublicTask  services.ITask
	User        services.IUser
}

func NewHandler() *Handler {
	return &Handler{
		PrivateTask: services.NewPrivateTask(),
		PublicTask:  services.NewTask(),
		User:        services.NewUser(),
	}
}
