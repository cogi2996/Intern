package services

import (
	"context"

	"github.com/google/uuid"
	"github.com/ideal-forward/assistant-public-api/entities"
	"github.com/ideal-forward/assistant-public-api/models"
)

type IPrivateTaskImage interface {
	Create(ctx context.Context, data *entities.PrivateTaskAttachFile) (int64, error)
	Delete(ctx context.Context, id int64) error
	List(ctx context.Context, taskID int64) ([]*entities.PrivateTaskAttachFile, error)
}

type PrivateTaskImage struct {
	Model models.IPrivateTaskAttachFile
}

func NewPrivateTaskImage() IPrivateTaskImage {
	return &PrivateTaskImage{
		Model: models.PrivateTaskAttachFile{},
	}
}

func (p *PrivateTaskImage) Create(ctx context.Context, data *entities.PrivateTaskAttachFile) (int64, error) {
	data.UUID = uuid.NewString()
	return p.Model.Create(ctx, data)
}

func (p *PrivateTaskImage) Delete(ctx context.Context, id int64) error {
	return p.Model.Delete(ctx, id)
}

func (p *PrivateTaskImage) List(ctx context.Context, taskID int64) ([]*entities.PrivateTaskAttachFile, error) {
	return p.Model.List(ctx, "task_id", taskID)
}
