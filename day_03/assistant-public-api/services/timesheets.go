package services

import (
	"context"

	"github.com/google/uuid"
	"github.com/ideal-forward/assistant-public-api/entities"
	"github.com/ideal-forward/assistant-public-api/models"
)

type ITimesheets interface {
	Create(ctx context.Context, data *entities.Timesheets) (int64, error)
	Read(ctx context.Context, id int64) (*entities.Timesheets, error)
	Update(ctx context.Context, data *entities.Timesheets) error
	Delete(ctx context.Context, id int64) error
	List(ctx context.Context, projectID int64) ([]*entities.Timesheets, error)

	CreateExecutor(ctx context.Context, data *entities.TimesheetsExecutor) (int64, error)
	ReadExecutor(ctx context.Context, id int64) (*entities.TimesheetsExecutor, error)
	UpdateExecutor(ctx context.Context, data *entities.TimesheetsExecutor) error
	DeleteExecutor(ctx context.Context, id int64) error
	ListExecutor(ctx context.Context, timesheetsID int64) ([]*entities.TimesheetsExecutor, error)

	CreateComment(ctx context.Context, data *entities.TimesheetsComment) (int64, error)
	ListComment(ctx context.Context, timesheetsID int64) ([]*entities.TimesheetsComment, error)
}

type Timesheets struct {
	Model    models.ITimesheets
	Executor models.ITimesheetsExecutor
	Comment  models.ITimesheetsComment
}

func NewTimesheets() ITimesheets {
	return &Timesheets{
		Model:    models.Timesheets{},
		Executor: models.TimesheetsExecutor{},
		Comment:  models.TimesheetsComment{},
	}
}

func (t *Timesheets) Create(ctx context.Context, data *entities.Timesheets) (int64, error) {
	data.UUID = uuid.NewString()
	return t.Model.Create(ctx, data)
}

func (t *Timesheets) Read(ctx context.Context, id int64) (*entities.Timesheets, error) {
	return t.Model.Read(ctx, id)
}

func (t *Timesheets) Update(ctx context.Context, data *entities.Timesheets) error {
	return t.Model.Update(ctx, data)
}

func (t *Timesheets) Delete(ctx context.Context, id int64) error {
	return t.Model.Delete(ctx, id)
}

func (t *Timesheets) List(ctx context.Context, projectID int64) ([]*entities.Timesheets, error) {
	return t.Model.List(ctx, projectID)
}

func (t *Timesheets) CreateExecutor(ctx context.Context, data *entities.TimesheetsExecutor) (int64, error) {
	data.UUID = uuid.NewString()
	return t.Executor.Create(ctx, data)
}

func (t *Timesheets) ReadExecutor(ctx context.Context, id int64) (*entities.TimesheetsExecutor, error) {
	return t.Executor.Read(ctx, id)
}

func (t *Timesheets) UpdateExecutor(ctx context.Context, data *entities.TimesheetsExecutor) error {
	return t.Executor.Update(ctx, data)
}

func (t *Timesheets) DeleteExecutor(ctx context.Context, id int64) error {
	return t.Executor.Delete(ctx, id)
}

func (t *Timesheets) ListExecutor(ctx context.Context, timesheetsID int64) ([]*entities.TimesheetsExecutor, error) {
	return t.Executor.List(ctx, timesheetsID)
}

func (t *Timesheets) CreateComment(ctx context.Context, data *entities.TimesheetsComment) (int64, error) {
	data.UUID = uuid.NewString()
	return t.Comment.Create(ctx, data)
}

func (t *Timesheets) ListComment(ctx context.Context, timesheetsID int64) ([]*entities.TimesheetsComment, error) {
	return t.Comment.List(ctx, timesheetsID)
}
