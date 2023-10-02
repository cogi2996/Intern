package audit

import (
	"github.com/ideal-forward/assistant-public-api/services"
)

type Handler struct {
	Project      services.IProject
	Task         services.IAuditTask
	TaskComment  services.IAuditTaskComment
	TaskImage    services.IAuditTaskImage
	User         services.IUser
	Register     services.IRegister
	FileUploader services.IFileUploader
	FileNamer    services.IFileNamer
	Excel        services.IExcel
}

func NewHandler() *Handler {
	return &Handler{
		Project:      services.NewProject(),
		Task:         services.NewAuditTask(),
		TaskComment:  services.NewAuditTaskComment(),
		TaskImage:    services.NewAuditTaskImage(),
		User:         services.NewUser(),
		Register:     services.NewRegister(),
		FileUploader: services.NewFileUploader(),
		FileNamer:    services.NewFileNamer(),
		Excel:        services.NewExcel(),
	}
}
