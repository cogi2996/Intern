package services

import (
	"context"

	"github.com/google/uuid"
	"github.com/ideal-forward/assistant-public-api/entities"
	"github.com/ideal-forward/assistant-public-api/models"
	"gorm.io/gorm"
)

type IExecutor interface {
	Create(ctx context.Context, data *entities.Executor) (int64, error)
	Read(ctx context.Context, id int64) (*entities.Executor, error)
	ReadByRepresenterID(ctx context.Context, representerID int64) (*entities.Executor, error)
	Update(ctx context.Context, data *entities.Executor) error
	Delete(ctx context.Context, id int64) error
	GetList(ctx context.Context) ([]*entities.Executor, error)
}

type Executor struct {
	Model models.IExecutor
}

func NewExecutor() IExecutor {
	return &Executor{
		Model: models.Executor{},
	}
}

func (p *Executor) Create(ctx context.Context, data *entities.Executor) (int64, error) {
	data.UUID = uuid.NewString()
	return p.Model.Create(ctx, data)
}

func (p *Executor) Read(ctx context.Context, id int64) (*entities.Executor, error) {
	return p.Model.Read(ctx, id)
}

func (p *Executor) ReadByRepresenterID(ctx context.Context, representerID int64) (*entities.Executor, error) {
	data, err := p.Model.ListWithCondition(ctx, "represented_by", representerID)
	if err != nil {
		return nil, err
	}

	if len(data) < 1 {
		return nil, gorm.ErrRecordNotFound
	}

	return data[0], nil
}

func (p *Executor) Update(ctx context.Context, data *entities.Executor) error {
	return p.Model.Update(ctx, data)
}

func (p *Executor) Delete(ctx context.Context, id int64) error {
	return p.Model.Delete(ctx, id)
}

func (p *Executor) GetList(ctx context.Context) ([]*entities.Executor, error) {
	return p.Model.List(ctx)
}
