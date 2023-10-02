package models

import (
	"context"
	"fmt"

	"github.com/ideal-forward/assistant-public-api/clients"
	"github.com/ideal-forward/assistant-public-api/entities"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ITaskAttachFile interface {
	Create(ctx context.Context, data *entities.TaskAttachFile) (int64, error)
	Delete(ctx context.Context, id int64) error
	List(ctx context.Context, field string, value any) ([]*entities.TaskAttachFile, error)
}

type TaskAttachFile struct {
}

func (TaskAttachFile) Create(ctx context.Context, data *entities.TaskAttachFile) (int64, error) {
	err := clients.MySQLClient.WithContext(ctx).Create(data).Error
	return data.ID, err
}

func (TaskAttachFile) Delete(ctx context.Context, id int64) error {
	return clients.MySQLClient.WithContext(ctx).Model(&entities.TaskAttachFile{}).Session(&gorm.Session{FullSaveAssociations: true}).Where("id = ?", id).Delete(&entities.TaskAttachFile{}).Error
}

func (TaskAttachFile) List(ctx context.Context, field string, value any) ([]*entities.TaskAttachFile, error) {
	result := []*entities.TaskAttachFile{}
	err := clients.MySQLClient.WithContext(ctx).Preload(clause.Associations).Where(fmt.Sprintf("%s = ?", field), value).Find(&result).Error
	if err != nil {
		return nil, err
	}

	return result, nil
}
