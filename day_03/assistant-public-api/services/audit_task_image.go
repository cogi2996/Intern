package services

import (
	"context"

	"github.com/google/uuid"
	"github.com/ideal-forward/assistant-public-api/entities"
	"github.com/ideal-forward/assistant-public-api/models"
)

type IAuditTaskImage interface {
	Create(ctx context.Context, data *entities.AuditTaskAttachFile) (int64, error)
	Delete(ctx context.Context, id int64) error
	List(ctx context.Context, taskID int64) ([]*entities.AuditTaskAttachFile, error)
}

type AuditTaskImage struct {
	Model models.IAuditTaskAttachFile
}

func NewAuditTaskImage() IAuditTaskImage {
	return &AuditTaskImage{
		Model: models.AuditTaskAttachFile{},
	}
}

func (p *AuditTaskImage) Create(ctx context.Context, data *entities.AuditTaskAttachFile) (int64, error) {
	data.UUID = uuid.NewString()
	return p.Model.Create(ctx, data)
}

func (p *AuditTaskImage) Delete(ctx context.Context, id int64) error {
	return p.Model.Delete(ctx, id)
}

func (p *AuditTaskImage) List(ctx context.Context, taskID int64) ([]*entities.AuditTaskAttachFile, error) {
	return p.Model.List(ctx, "task_id", taskID)
}
