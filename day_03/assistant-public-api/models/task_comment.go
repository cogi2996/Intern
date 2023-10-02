package models

import (
	"context"
	"fmt"

	"github.com/ideal-forward/assistant-public-api/clients"
	"github.com/ideal-forward/assistant-public-api/entities"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ITaskComment interface {
	Create(ctx context.Context, data *entities.TaskComment) (int64, error)
	Read(ctx context.Context, id int64) (*entities.TaskComment, error)
	Update(ctx context.Context, data *entities.TaskComment) error
	Delete(ctx context.Context, id int64) error
	List(ctx context.Context, field string, value any) ([]*entities.TaskComment, error)
}

type TaskComment struct {
}

func (TaskComment) Create(ctx context.Context, data *entities.TaskComment) (int64, error) {
	err := clients.MySQLClient.WithContext(ctx).Create(data).Error
	return data.ID, err
}

func (TaskComment) Read(ctx context.Context, id int64) (*entities.TaskComment, error) {
	result := &entities.TaskComment{}
	err := clients.MySQLClient.WithContext(ctx).
		Preload("Creator").
		Preload("Task").
		Where("id = ?", id).First(result).Error
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (TaskComment) Update(ctx context.Context, data *entities.TaskComment) error {
	return clients.MySQLClient.WithContext(ctx).Model(data).Session(&gorm.Session{FullSaveAssociations: true}).Where("id = ?", data.ID).Updates(data).Error
}

func (TaskComment) Delete(ctx context.Context, id int64) error {
	return clients.MySQLClient.WithContext(ctx).Model(&entities.TaskComment{}).Session(&gorm.Session{FullSaveAssociations: true}).Where("id = ?", id).Delete(&entities.TaskComment{}).Error
}

func (TaskComment) List(ctx context.Context, field string, value any) ([]*entities.TaskComment, error) {
	result := []*entities.TaskComment{}
	err := clients.MySQLClient.WithContext(ctx).Preload(clause.Associations).Where(fmt.Sprintf("%s = ?", field), value).Find(&result).Error
	if err != nil {
		return nil, err
	}

	return result, nil
}
