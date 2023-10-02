package services

import (
	"context"

	"github.com/google/uuid"
	"github.com/ideal-forward/assistant-public-api/entities"
	"github.com/ideal-forward/assistant-public-api/models"
)

type IAuditTask interface {
	Create(ctx context.Context, data *entities.AuditTask) (int64, error)
	Read(ctx context.Context, id int64) (*entities.AuditTask, error)
	Update(ctx context.Context, data *entities.AuditTask) error
	UpdateStatus(ctx context.Context, userID, id int64, status entities.TaskStatus) error
	Delete(ctx context.Context, id int64) error
	List(ctx context.Context, filter *TaskFilter) ([]*entities.AuditTask, error)
}

type AuditTask struct {
	Model models.IAuditTask
}

func NewAuditTask() IAuditTask {
	return &AuditTask{
		Model: models.AuditTask{},
	}
}

func (p *AuditTask) Create(ctx context.Context, data *entities.AuditTask) (int64, error) {
	data.UUID = uuid.NewString()
	data.Status = entities.StatusCreated.Value()
	return p.Model.Create(ctx, data)
}

func (p *AuditTask) Read(ctx context.Context, id int64) (*entities.AuditTask, error) {
	return p.Model.Read(ctx, id)
}

func (p *AuditTask) Update(ctx context.Context, data *entities.AuditTask) error {
	return p.Model.Update(ctx, data)
}

func (p *AuditTask) UpdateStatus(ctx context.Context, userID, id int64, status entities.TaskStatus) error {
	return p.Model.UpdateField(ctx, userID, id, "status", status.Value())
}

func (p *AuditTask) Delete(ctx context.Context, id int64) error {
	return p.Model.Delete(ctx, id)
}

func (p *AuditTask) List(ctx context.Context, filter *TaskFilter) ([]*entities.AuditTask, error) {
	return p.Model.List(ctx, filter.ToMap())
}
