package services

import (
	"context"

	"github.com/google/uuid"
	"github.com/ideal-forward/assistant-public-api/entities"
	"github.com/ideal-forward/assistant-public-api/models"
)

type IPrivateTaskComment interface {
	Create(ctx context.Context, data *entities.PrivateTaskComment) (int64, error)
	Delete(ctx context.Context, id int64) error
	List(ctx context.Context, taskID int64) ([]*entities.PrivateTaskComment, error)
}

type PrivateTaskComment struct {
	Model models.IPrivateTaskComment
}

func NewPrivateTaskComment() IPrivateTaskComment {
	return &PrivateTaskComment{
		Model: models.PrivateTaskComment{},
	}
}

func (p *PrivateTaskComment) Create(ctx context.Context, data *entities.PrivateTaskComment) (int64, error) {
	data.UUID = uuid.NewString()
	return p.Model.Create(ctx, data)
}

func (p *PrivateTaskComment) Delete(ctx context.Context, id int64) error {
	return p.Model.Delete(ctx, id)
}

func (p *PrivateTaskComment) List(ctx context.Context, taskID int64) ([]*entities.PrivateTaskComment, error) {
	return p.Model.List(ctx, "task_id", taskID)
}
