package models

import (
	"context"
	"fmt"

	"github.com/ideal-forward/assistant-public-api/clients"
	"github.com/ideal-forward/assistant-public-api/entities"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type IPrivateTaskComment interface {
	Create(ctx context.Context, data *entities.PrivateTaskComment) (int64, error)
	Read(ctx context.Context, id int64) (*entities.PrivateTaskComment, error)
	Update(ctx context.Context, data *entities.PrivateTaskComment) error
	Delete(ctx context.Context, id int64) error
	List(ctx context.Context, field string, value any) ([]*entities.PrivateTaskComment, error)
}

type PrivateTaskComment struct {
}

func (PrivateTaskComment) Create(ctx context.Context, data *entities.PrivateTaskComment) (int64, error) {
	err := clients.MySQLClient.WithContext(ctx).Create(data).Error
	return data.ID, err
}

func (PrivateTaskComment) Read(ctx context.Context, id int64) (*entities.PrivateTaskComment, error) {
	result := &entities.PrivateTaskComment{}
	err := clients.MySQLClient.WithContext(ctx).
		Preload("Creator").
		Preload("Task").
		Where("id = ?", id).First(result).Error
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (PrivateTaskComment) Update(ctx context.Context, data *entities.PrivateTaskComment) error {
	return clients.MySQLClient.WithContext(ctx).Model(data).Session(&gorm.Session{FullSaveAssociations: true}).Where("id = ?", data.ID).Updates(data).Error
}

func (PrivateTaskComment) Delete(ctx context.Context, id int64) error {
	return clients.MySQLClient.WithContext(ctx).Model(&entities.PrivateTaskComment{}).Session(&gorm.Session{FullSaveAssociations: true}).Where("id = ?", id).Delete(&entities.PrivateTaskComment{}).Error
}

func (PrivateTaskComment) List(ctx context.Context, field string, value any) ([]*entities.PrivateTaskComment, error) {
	result := []*entities.PrivateTaskComment{}
	err := clients.MySQLClient.WithContext(ctx).Preload(clause.Associations).Where(fmt.Sprintf("%s = ?", field), value).Find(&result).Error
	if err != nil {
		return nil, err
	}

	return result, nil
}
