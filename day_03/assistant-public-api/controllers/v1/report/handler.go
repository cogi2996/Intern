package report

import "github.com/ideal-forward/assistant-public-api/services"

type Handler struct {
	Task services.ITask
}

func NewHandler() *Handler {
	return &Handler{
		Task: services.NewTask(),
	}
}
