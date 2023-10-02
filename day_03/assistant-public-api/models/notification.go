package models

import (
	"context"
	"fmt"

	"github.com/ideal-forward/assistant-public-api/clients"
	"github.com/ideal-forward/assistant-public-api/entities"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type INotification interface {
	Create(ctx context.Context, data *entities.Notification) (int64, error)
	Read(ctx context.Context, id int64) (*entities.Notification, error)
	Update(ctx context.Context, data *entities.Notification) error
	Delete(ctx context.Context, id int64) error
	List(ctx context.Context) ([]*entities.Notification, error)
	ListWithCondition(ctx context.Context, field string, value any) ([]*entities.Notification, error)
}

type Notification struct {
}

func (Notification) Create(ctx context.Context, data *entities.Notification) (int64, error) {
	err := clients.MySQLClient.Transaction(func(tx *gorm.DB) error {
		return tx.WithContext(ctx).Session(&gorm.Session{FullSaveAssociations: true}).Create(data).Error
	})

	return data.ID, err
}

func (Notification) Read(ctx context.Context, id int64) (*entities.Notification, error) {
	result := &entities.Notification{}
	err := clients.MySQLClient.WithContext(ctx).Preload("Creator").Preload("PublicTask").Where("id = ?", id).First(result).Error
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (Notification) Update(ctx context.Context, data *entities.Notification) error {
	return clients.MySQLClient.Transaction(func(tx *gorm.DB) error {
		return tx.WithContext(ctx).Model(data).Session(&gorm.Session{FullSaveAssociations: true}).Where("id = ?", data.ID).Updates(data).Error
	})
}

func (Notification) Delete(ctx context.Context, id int64) error {
	return clients.MySQLClient.Transaction(func(tx *gorm.DB) error {
		return tx.WithContext(ctx).Model(&entities.Notification{}).Session(&gorm.Session{FullSaveAssociations: true}).Where("id = ?", id).Delete(&entities.Notification{}).Error
	})
}

func (Notification) List(ctx context.Context) ([]*entities.Notification, error) {
	result := []*entities.Notification{}
	err := clients.MySQLClient.WithContext(ctx).Preload("Creator").Preload("PublicTask").Order("created_at DESC").Find(&result).Error
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (Notification) ListWithCondition(ctx context.Context, field string, value any) ([]*entities.Notification, error) {
	result := []*entities.Notification{}
	err := clients.MySQLClient.WithContext(ctx).Preload(clause.Associations).Where(fmt.Sprintf("%s = ?", field), value).Find(&result).Error
	if err != nil {
		return nil, err
	}

	return result, nil
}
