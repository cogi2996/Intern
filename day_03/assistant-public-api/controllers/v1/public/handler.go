package public

import (
	"github.com/ideal-forward/assistant-public-api/services"
)

type Handler struct {
	Project      services.IProject
	Task         services.ITask
	User         services.IUser
	Register     services.IRegister
	Executor     services.IExecutor
	TaskComment  services.ITaskComment
	TaskImage    services.ITaskImage
	FileUploader services.IFileUploader
	FileNamer    services.IFileNamer
	Excel        services.IExcel
	Notification services.INotification
}

func NewHandler() *Handler {
	return &Handler{
		Project:      services.NewProject(),
		Task:         services.NewTask(),
		User:         services.NewUser(),
		Register:     services.NewRegister(),
		Executor:     services.NewExecutor(),
		TaskComment:  services.NewTaskComment(),
		TaskImage:    services.NewTaskImage(),
		FileUploader: services.NewFileUploader(),
		FileNamer:    services.NewFileNamer(),
		Excel:        services.NewExcel(),
		Notification: services.NewNotification(),
	}
}
