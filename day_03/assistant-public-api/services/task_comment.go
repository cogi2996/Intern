package services

import (
	"context"

	"github.com/google/uuid"
	"github.com/ideal-forward/assistant-public-api/entities"
	"github.com/ideal-forward/assistant-public-api/models"
)

type ITaskComment interface {
	Create(ctx context.Context, data *entities.TaskComment) (int64, error)
	Delete(ctx context.Context, id int64) error
	List(ctx context.Context, taskID int64) ([]*entities.TaskComment, error)
}

type TaskComment struct {
	Model models.ITaskComment
}

func NewTaskComment() ITaskComment {
	return &TaskComment{
		Model: models.TaskComment{},
	}
}

func (p *TaskComment) Create(ctx context.Context, data *entities.TaskComment) (int64, error) {
	data.UUID = uuid.NewString()
	return p.Model.Create(ctx, data)
}

func (p *TaskComment) Delete(ctx context.Context, id int64) error {
	return p.Model.Delete(ctx, id)
}

func (p *TaskComment) List(ctx context.Context, taskID int64) ([]*entities.TaskComment, error) {
	return p.Model.List(ctx, "task_id", taskID)
}
