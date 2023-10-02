package models

import (
	"context"
	"fmt"

	"github.com/ideal-forward/assistant-public-api/clients"
	"github.com/ideal-forward/assistant-public-api/entities"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type IAuditTaskAttachFile interface {
	Create(ctx context.Context, data *entities.AuditTaskAttachFile) (int64, error)
	Delete(ctx context.Context, id int64) error
	List(ctx context.Context, field string, value any) ([]*entities.AuditTaskAttachFile, error)
}

type AuditTaskAttachFile struct {
}

func (AuditTaskAttachFile) Create(ctx context.Context, data *entities.AuditTaskAttachFile) (int64, error) {
	err := clients.MySQLClient.WithContext(ctx).Create(data).Error
	return data.ID, err
}

func (AuditTaskAttachFile) Delete(ctx context.Context, id int64) error {
	return clients.MySQLClient.WithContext(ctx).Model(&entities.AuditTaskAttachFile{}).Session(&gorm.Session{FullSaveAssociations: true}).Where("id = ?", id).Delete(&entities.AuditTaskAttachFile{}).Error
}

func (AuditTaskAttachFile) List(ctx context.Context, field string, value any) ([]*entities.AuditTaskAttachFile, error) {
	result := []*entities.AuditTaskAttachFile{}
	err := clients.MySQLClient.WithContext(ctx).Preload(clause.Associations).Where(fmt.Sprintf("%s = ?", field), value).Find(&result).Error
	if err != nil {
		return nil, err
	}

	return result, nil
}
