package services

import (
	"context"

	"github.com/google/uuid"
	"github.com/ideal-forward/assistant-public-api/entities"
	"github.com/ideal-forward/assistant-public-api/models"
)

type IPrivateTask interface {
	Create(ctx context.Context, data *entities.PrivateTask) (int64, error)
	Read(ctx context.Context, id int64) (*entities.PrivateTask, error)
	Update(ctx context.Context, data *entities.PrivateTask) error
	UpdateStatus(ctx context.Context, userID, id int64, status entities.TaskStatus) error
	Delete(ctx context.Context, id int64) error
	List(ctx context.Context, filter *TaskFilter) ([]*entities.PrivateTask, error)
}

type PrivateTask struct {
	Model models.IPrivateTask
}

func NewPrivateTask() IPrivateTask {
	return &PrivateTask{
		Model: models.PrivateTask{},
	}
}

func (p *PrivateTask) Create(ctx context.Context, data *entities.PrivateTask) (int64, error) {
	data.UUID = uuid.NewString()
	data.Status = entities.StatusCreated.Value()
	return p.Model.Create(ctx, data)
}

func (p *PrivateTask) Read(ctx context.Context, id int64) (*entities.PrivateTask, error) {
	return p.Model.Read(ctx, id)
}

func (p *PrivateTask) Update(ctx context.Context, data *entities.PrivateTask) error {
	return p.Model.Update(ctx, data)
}

func (p *PrivateTask) UpdateStatus(ctx context.Context, userID, id int64, status entities.TaskStatus) error {
	return p.Model.UpdateField(ctx, userID, id, "status", status.Value())
}

func (p *PrivateTask) Delete(ctx context.Context, id int64) error {
	return p.Model.Delete(ctx, id)
}

func (p *PrivateTask) List(ctx context.Context, filter *TaskFilter) ([]*entities.PrivateTask, error) {
	return p.Model.List(ctx, filter.ToMap())
}
