package private

import (
	"github.com/ideal-forward/assistant-public-api/services"
)

type Handler struct {
	Project      services.IProject
	Task         services.IPrivateTask
	TaskComment  services.IPrivateTaskComment
	TaskImage    services.IPrivateTaskImage
	User         services.IUser
	Register     services.IRegister
	FileUploader services.IFileUploader
	FileNamer    services.IFileNamer
	Excel        services.IExcel
}

func NewHandler() *Handler {
	return &Handler{
		Project:      services.NewProject(),
		Task:         services.NewPrivateTask(),
		TaskComment:  services.NewPrivateTaskComment(),
		TaskImage:    services.NewPrivateTaskImage(),
		User:         services.NewUser(),
		Register:     services.NewRegister(),
		FileUploader: services.NewFileUploader(),
		FileNamer:    services.NewFileNamer(),
		Excel:        services.NewExcel(),
	}
}
