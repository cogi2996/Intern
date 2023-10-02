package services

import (
	"context"

	"github.com/google/uuid"
	"github.com/ideal-forward/assistant-public-api/entities"
	"github.com/ideal-forward/assistant-public-api/models"
)

type INotification interface {
	Create(ctx context.Context, data *entities.Notification) (int64, error)
	Read(ctx context.Context, id int64) (*entities.Notification, error)
	GetList(ctx context.Context) ([]*entities.Notification, error)
	GetListByReceiverID(ctx context.Context, receiverID int64) ([]*entities.Notification, error)
}

type Notification struct {
	Model models.INotification
}

func NewNotification() INotification {
	return &Notification{
		Model: models.Notification{},
	}
}

func (p *Notification) Create(ctx context.Context, data *entities.Notification) (int64, error) {
	data.UUID = uuid.NewString()
	return p.Model.Create(ctx, data)
}

func (p *Notification) Read(ctx context.Context, id int64) (*entities.Notification, error) {
	return p.Model.Read(ctx, id)
}

func (p *Notification) GetList(ctx context.Context) ([]*entities.Notification, error) {
	return p.Model.List(ctx)
}

func (p *Notification) GetListByReceiverID(ctx context.Context, receiverID int64) ([]*entities.Notification, error) {
	return p.Model.ListWithCondition(ctx, "received_by", receiverID)
}
