package services

import (
	"context"

	"github.com/google/uuid"
	"github.com/ideal-forward/assistant-public-api/entities"
	"github.com/ideal-forward/assistant-public-api/models"
)

type ITaskImage interface {
	Create(ctx context.Context, data *entities.TaskAttachFile) (int64, error)
	Delete(ctx context.Context, id int64) error
	List(ctx context.Context, taskID int64) ([]*entities.TaskAttachFile, error)
}

type TaskImage struct {
	Model models.ITaskAttachFile
}

func NewTaskImage() ITaskImage {
	return &TaskImage{
		Model: models.TaskAttachFile{},
	}
}

func (p *TaskImage) Create(ctx context.Context, data *entities.TaskAttachFile) (int64, error) {
	data.UUID = uuid.NewString()
	return p.Model.Create(ctx, data)
}

func (p *TaskImage) Delete(ctx context.Context, id int64) error {
	return p.Model.Delete(ctx, id)
}

func (p *TaskImage) List(ctx context.Context, taskID int64) ([]*entities.TaskAttachFile, error) {
	return p.Model.List(ctx, "task_id", taskID)
}
