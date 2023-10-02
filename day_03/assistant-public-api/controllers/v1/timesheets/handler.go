package timesheets

import "github.com/ideal-forward/assistant-public-api/services"

type Handler struct {
	Timesheets services.ITimesheets
	User       services.IUser
}

func NewHandler() *Handler {
	return &Handler{
		Timesheets: services.NewTimesheets(),
		User:       services.NewUser(),
	}
}
