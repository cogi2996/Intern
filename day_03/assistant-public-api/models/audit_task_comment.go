package models

import (
	"context"
	"fmt"

	"github.com/ideal-forward/assistant-public-api/clients"
	"github.com/ideal-forward/assistant-public-api/entities"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type IAuditTaskComment interface {
	Create(ctx context.Context, data *entities.AuditTaskComment) (int64, error)
	Read(ctx context.Context, id int64) (*entities.AuditTaskComment, error)
	Update(ctx context.Context, data *entities.AuditTaskComment) error
	Delete(ctx context.Context, id int64) error
	List(ctx context.Context, field string, value any) ([]*entities.AuditTaskComment, error)
}

type AuditTaskComment struct {
}

func (AuditTaskComment) Create(ctx context.Context, data *entities.AuditTaskComment) (int64, error) {
	err := clients.MySQLClient.WithContext(ctx).Create(data).Error
	return data.ID, err
}

func (AuditTaskComment) Read(ctx context.Context, id int64) (*entities.AuditTaskComment, error) {
	result := &entities.AuditTaskComment{}
	err := clients.MySQLClient.WithContext(ctx).
		Preload("Creator").
		Preload("Task").
		Where("id = ?", id).First(result).Error
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (AuditTaskComment) Update(ctx context.Context, data *entities.AuditTaskComment) error {
	return clients.MySQLClient.WithContext(ctx).Model(data).Session(&gorm.Session{FullSaveAssociations: true}).Where("id = ?", data.ID).Updates(data).Error
}

func (AuditTaskComment) Delete(ctx context.Context, id int64) error {
	return clients.MySQLClient.WithContext(ctx).Model(&entities.AuditTaskComment{}).Session(&gorm.Session{FullSaveAssociations: true}).Where("id = ?", id).Delete(&entities.AuditTaskComment{}).Error
}

func (AuditTaskComment) List(ctx context.Context, field string, value any) ([]*entities.AuditTaskComment, error) {
	result := []*entities.AuditTaskComment{}
	err := clients.MySQLClient.WithContext(ctx).Preload(clause.Associations).Where(fmt.Sprintf("%s = ?", field), value).Find(&result).Error
	if err != nil {
		return nil, err
	}

	return result, nil
}
