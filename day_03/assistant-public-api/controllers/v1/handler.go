package v1

import (
	"github.com/ideal-forward/assistant-public-api/services"
)

type Handler struct {
	ProjectService         services.IProject
	TaskService            services.ITask
	UserService            services.IUser
	RegisterService        services.IRegister
	ExecutorService        services.IExecutor
	TaskCommentService     services.ITaskComment
	TaskImageService       services.ITaskImage
	FileUploaderService    services.IFileUploader
	FileNamerService       services.IFileNamer
	RegionService          services.IRegion
	ProjectCategoryService services.IProjectCategory
	Notification           services.INotification
}

func NewHandler() *Handler {
	return &Handler{
		ProjectService:         services.NewProject(),
		TaskService:            services.NewTask(),
		UserService:            services.NewUser(),
		RegisterService:        services.NewRegister(),
		ExecutorService:        services.NewExecutor(),
		TaskCommentService:     services.NewTaskComment(),
		TaskImageService:       services.NewTaskImage(),
		FileUploaderService:    services.NewFileUploader(),
		FileNamerService:       services.NewFileNamer(),
		RegionService:          services.NewRegion(),
		ProjectCategoryService: services.NewProjectCategory(),
		Notification:           services.NewNotification(),
	}
}
