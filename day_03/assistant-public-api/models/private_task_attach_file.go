package models

import (
	"context"
	"fmt"

	"github.com/ideal-forward/assistant-public-api/clients"
	"github.com/ideal-forward/assistant-public-api/entities"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type IPrivateTaskAttachFile interface {
	Create(ctx context.Context, data *entities.PrivateTaskAttachFile) (int64, error)
	Delete(ctx context.Context, id int64) error
	List(ctx context.Context, field string, value any) ([]*entities.PrivateTaskAttachFile, error)
}

type PrivateTaskAttachFile struct {
}

func (PrivateTaskAttachFile) Create(ctx context.Context, data *entities.PrivateTaskAttachFile) (int64, error) {
	err := clients.MySQLClient.WithContext(ctx).Create(data).Error
	return data.ID, err
}

func (PrivateTaskAttachFile) Delete(ctx context.Context, id int64) error {
	return clients.MySQLClient.WithContext(ctx).Model(&entities.PrivateTaskAttachFile{}).Session(&gorm.Session{FullSaveAssociations: true}).Where("id = ?", id).Delete(&entities.PrivateTaskAttachFile{}).Error
}

func (PrivateTaskAttachFile) List(ctx context.Context, field string, value any) ([]*entities.PrivateTaskAttachFile, error) {
	result := []*entities.PrivateTaskAttachFile{}
	err := clients.MySQLClient.WithContext(ctx).Preload(clause.Associations).Where(fmt.Sprintf("%s = ?", field), value).Find(&result).Error
	if err != nil {
		return nil, err
	}

	return result, nil
}
