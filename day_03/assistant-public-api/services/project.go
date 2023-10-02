package services

import (
	"context"

	"github.com/google/uuid"
	"github.com/ideal-forward/assistant-public-api/entities"
	"github.com/ideal-forward/assistant-public-api/models"
)

type IProject interface {
	Create(ctx context.Context, data *entities.Project) (int64, error)
	Read(ctx context.Context, id int64) (*entities.Project, error)
	Update(ctx context.Context, data *entities.Project) error
	Delete(ctx context.Context, id int64) error
	GetList(ctx context.Context) ([]*entities.Project, error)
	GetListByExecutor(ctx context.Context, executorID int64) ([]*entities.Project, error)

	AddMember(ctx context.Context, data *entities.ProjectMember) error
	DeleteMember(ctx context.Context, data *entities.ProjectMember) error
	ListMember(ctx context.Context, projectID int64) ([]*entities.ProjectMember, error)

	AddExecutor(ctx context.Context, data *entities.ProjectExecutor) error
	DeleteExecutor(ctx context.Context, data *entities.ProjectExecutor) error
	ListExecutor(ctx context.Context, projectID int64) ([]*entities.ProjectExecutor, error)

	CreatePhase(ctx context.Context, data *entities.ProjectPhase) error
	ReadPhase(ctx context.Context, id int64) (*entities.ProjectPhase, error)
	UpdatePhase(ctx context.Context, data *entities.ProjectPhase) error
	DeletePhase(ctx context.Context, data *entities.ProjectPhase) error
	ListPhase(ctx context.Context, projectID int64) ([]*entities.ProjectPhase, error)

	CreateArea(ctx context.Context, data *entities.ProjectArea) error
	ReadArea(ctx context.Context, id int64) (*entities.ProjectArea, error)
	UpdateArea(ctx context.Context, data *entities.ProjectArea) error
	DeleteArea(ctx context.Context, data *entities.ProjectArea) error
	ListArea(ctx context.Context, projectID int64) ([]*entities.ProjectArea, error)
}

type Project struct {
	Model models.IProject
}

func NewProject() IProject {
	return &Project{
		Model: models.Project{},
	}
}

func (p *Project) Create(ctx context.Context, data *entities.Project) (int64, error) {
	data.UUID = uuid.NewString()
	return p.Model.Create(ctx, data)
}

func (p *Project) Read(ctx context.Context, id int64) (*entities.Project, error) {
	return p.Model.Read(ctx, id)
}

func (p *Project) Update(ctx context.Context, data *entities.Project) error {
	return p.Model.Update(ctx, data)
}

func (p *Project) Delete(ctx context.Context, id int64) error {
	return p.Model.Delete(ctx, id)
}

func (p *Project) GetList(ctx context.Context) ([]*entities.Project, error) {
	return p.Model.GetList(ctx)
}

func (p *Project) GetListByExecutor(ctx context.Context, executorID int64) ([]*entities.Project, error) {
	return p.Model.GetListByExecutor(ctx, executorID)
}

func (p *Project) AddMember(ctx context.Context, data *entities.ProjectMember) error {
	return p.Model.AddMember(ctx, data)
}

func (p *Project) DeleteMember(ctx context.Context, data *entities.ProjectMember) error {
	return p.Model.DeleteMember(ctx, data)
}

func (p *Project) ListMember(ctx context.Context, projectID int64) ([]*entities.ProjectMember, error) {
	return p.Model.GetListMember(ctx, projectID)
}

func (p *Project) AddExecutor(ctx context.Context, data *entities.ProjectExecutor) error {
	return p.Model.AddExecutor(ctx, data)
}

func (p *Project) DeleteExecutor(ctx context.Context, data *entities.ProjectExecutor) error {
	return p.Model.DeleteExecutor(ctx, data)
}

func (p *Project) ListExecutor(ctx context.Context, projectID int64) ([]*entities.ProjectExecutor, error) {
	return p.Model.GetListExecutor(ctx, projectID)
}

func (p *Project) CreatePhase(ctx context.Context, data *entities.ProjectPhase) error {
	return p.Model.CreatePhase(ctx, data)
}

func (p *Project) ReadPhase(ctx context.Context, id int64) (*entities.ProjectPhase, error) {
	return p.Model.ReadPhase(ctx, id)
}

func (p *Project) UpdatePhase(ctx context.Context, data *entities.ProjectPhase) error {
	return p.Model.UpdatePhase(ctx, data)
}

func (p *Project) DeletePhase(ctx context.Context, data *entities.ProjectPhase) error {
	return p.Model.DeletePhase(ctx, data)
}

func (p *Project) ListPhase(ctx context.Context, projectID int64) ([]*entities.ProjectPhase, error) {
	return p.Model.GetListPhase(ctx, projectID)
}

func (p *Project) CreateArea(ctx context.Context, data *entities.ProjectArea) error {
	return p.Model.CreateArea(ctx, data)
}

func (p *Project) ReadArea(ctx context.Context, id int64) (*entities.ProjectArea, error) {
	return p.Model.ReadArea(ctx, id)
}

func (p *Project) UpdateArea(ctx context.Context, data *entities.ProjectArea) error {
	return p.Model.UpdateArea(ctx, data)
}

func (p *Project) DeleteArea(ctx context.Context, data *entities.ProjectArea) error {
	return p.Model.DeleteArea(ctx, data)
}

func (p *Project) ListArea(ctx context.Context, projectID int64) ([]*entities.ProjectArea, error) {
	return p.Model.GetListArea(ctx, projectID)
}
