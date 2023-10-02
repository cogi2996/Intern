package services

import (
	"context"

	"github.com/google/uuid"
	"github.com/ideal-forward/assistant-public-api/entities"
	"github.com/ideal-forward/assistant-public-api/models"
)

type IAuditTaskComment interface {
	Create(ctx context.Context, data *entities.AuditTaskComment) (int64, error)
	Delete(ctx context.Context, id int64) error
	List(ctx context.Context, taskID int64) ([]*entities.AuditTaskComment, error)
}

type AuditTaskComment struct {
	Model models.IAuditTaskComment
}

func NewAuditTaskComment() IAuditTaskComment {
	return &AuditTaskComment{
		Model: models.AuditTaskComment{},
	}
}

func (p *AuditTaskComment) Create(ctx context.Context, data *entities.AuditTaskComment) (int64, error) {
	data.UUID = uuid.NewString()
	return p.Model.Create(ctx, data)
}

func (p *AuditTaskComment) Delete(ctx context.Context, id int64) error {
	return p.Model.Delete(ctx, id)
}

func (p *AuditTaskComment) List(ctx context.Context, taskID int64) ([]*entities.AuditTaskComment, error) {
	return p.Model.List(ctx, "task_id", taskID)
}
