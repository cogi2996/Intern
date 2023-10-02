package services

import (
	"context"

	"github.com/google/uuid"
	"github.com/ideal-forward/assistant-public-api/entities"
	"github.com/ideal-forward/assistant-public-api/models"
)

type IProjectCategory interface {
	Create(ctx context.Context, data *entities.ProjectCategory) (int64, error)
	Read(ctx context.Context, id int64) (*entities.ProjectCategory, error)
	Update(ctx context.Context, data *entities.ProjectCategory) error
	Delete(ctx context.Context, id int64) error
	List(ctx context.Context) ([]*entities.ProjectCategory, error)
}

type ProjectCategory struct {
	Model models.IProjectCategory
}

func NewProjectCategory() IProjectCategory {
	return &ProjectCategory{
		Model: models.ProjectCategory{},
	}
}

func (p *ProjectCategory) Create(ctx context.Context, data *entities.ProjectCategory) (int64, error) {
	data.UUID = uuid.NewString()
	return p.Model.Create(ctx, data)
}

func (p *ProjectCategory) Read(ctx context.Context, id int64) (*entities.ProjectCategory, error) {
	return p.Model.Read(ctx, id)
}

func (p *ProjectCategory) Update(ctx context.Context, data *entities.ProjectCategory) error {
	return p.Model.Update(ctx, data)
}

func (p *ProjectCategory) Delete(ctx context.Context, id int64) error {
	return p.Model.Delete(ctx, id)
}

func (p *ProjectCategory) List(ctx context.Context) ([]*entities.ProjectCategory, error) {
	return p.Model.List(ctx)
}
