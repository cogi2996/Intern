package services

import (
	"context"

	"github.com/google/uuid"
	"github.com/ideal-forward/assistant-public-api/entities"
	"github.com/ideal-forward/assistant-public-api/models"
	"gorm.io/gorm"
)

type ITask interface {
	Create(ctx context.Context, data *entities.Task) (int64, error)
	Read(ctx context.Context, id int64) (*entities.Task, error)
	Update(ctx context.Context, data *entities.Task) error
	UpdateStatus(ctx context.Context, id int64, status entities.TaskStatus, star int) error
	Delete(ctx context.Context, id int64) error
	List(ctx context.Context, filter *TaskFilter, withRelativeInfo bool) ([]*entities.Task, error)
	ListByCons(ctx context.Context, filters *TaskFilters, withRelativeInfo bool) ([]*entities.Task, error)
	ListConstructorTask(ctx context.Context, userID int64) ([]*entities.Task, error)
}

type Task struct {
	Model    models.ITask
	Executor models.IExecutor
}

func NewTask() ITask {
	return &Task{
		Model:    models.Task{},
		Executor: models.Executor{},
	}
}

func (p *Task) Create(ctx context.Context, data *entities.Task) (int64, error) {
	data.UUID = uuid.NewString()
	data.Status = entities.StatusCreated.Value()
	return p.Model.Create(ctx, data)
}

func (p *Task) Read(ctx context.Context, id int64) (*entities.Task, error) {
	return p.Model.Read(ctx, id)
}

func (p *Task) Update(ctx context.Context, data *entities.Task) error {
	return p.Model.Update(ctx, data)
}

func (p *Task) UpdateStatus(ctx context.Context, id int64, status entities.TaskStatus, star int) error {
	return p.Model.UpdateFields(ctx, id, map[string]interface{}{"status": status.Value(), "star": star})
}

func (p *Task) Delete(ctx context.Context, id int64) error {
	return p.Model.Delete(ctx, id)
}

func (p *Task) List(ctx context.Context, filter *TaskFilter, withRelativeInfo bool) ([]*entities.Task, error) {
	return p.Model.List(ctx, filter.ToMap(), withRelativeInfo)
}

func (p *Task) ListByCons(ctx context.Context, filters *TaskFilters, withRelativeInfo bool) ([]*entities.Task, error) {
	return p.Model.ListByCons(ctx, filters.ToMap(), withRelativeInfo)
}

func (p *Task) ListConstructorTask(ctx context.Context, userID int64) ([]*entities.Task, error) {
	executor, err := p.Executor.ListWithCondition(ctx, "represented_by", userID)
	if err != nil {
		return nil, err
	}

	if len(executor) < 1 {
		return nil, gorm.ErrRecordNotFound
	}

	filters := make(map[string]any)
	filters["executed_by"] = executor[0].ID
	return p.Model.List(ctx, filters, true)
}
