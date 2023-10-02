package services

import (
	"context"

	"github.com/google/uuid"
	"github.com/ideal-forward/assistant-public-api/entities"
	"github.com/ideal-forward/assistant-public-api/models"
)

type IRegion interface {
	Create(ctx context.Context, data *entities.Region) (int64, error)
	Read(ctx context.Context, id int64) (*entities.Region, error)
	Update(ctx context.Context, data *entities.Region) error
	Delete(ctx context.Context, id int64) error
	List(ctx context.Context) ([]*entities.Region, error)
}

type Region struct {
	Model models.IRegion
}

func NewRegion() IRegion {
	return &Region{
		Model: models.Region{},
	}
}

func (p *Region) Create(ctx context.Context, data *entities.Region) (int64, error) {
	data.UUID = uuid.NewString()
	return p.Model.Create(ctx, data)
}

func (p *Region) Read(ctx context.Context, id int64) (*entities.Region, error) {
	return p.Model.Read(ctx, id)
}

func (p *Region) Update(ctx context.Context, data *entities.Region) error {
	return p.Model.Update(ctx, data)
}

func (p *Region) Delete(ctx context.Context, id int64) error {
	return p.Model.Delete(ctx, id)
}

func (p *Region) List(ctx context.Context) ([]*entities.Region, error) {
	return p.Model.List(ctx)
}
