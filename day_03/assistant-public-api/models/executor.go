package models

import (
	"context"
	"fmt"

	"github.com/ideal-forward/assistant-public-api/clients"
	"github.com/ideal-forward/assistant-public-api/entities"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type IExecutor interface {
	Create(ctx context.Context, data *entities.Executor) (int64, error)
	Read(ctx context.Context, id int64) (*entities.Executor, error)
	Update(ctx context.Context, data *entities.Executor) error
	Delete(ctx context.Context, id int64) error
	List(ctx context.Context) ([]*entities.Executor, error)
	ListWithCondition(ctx context.Context, field string, value any) ([]*entities.Executor, error)
}

type Executor struct {
}

func (Executor) Create(ctx context.Context, data *entities.Executor) (int64, error) {
	err := clients.MySQLClient.Transaction(func(tx *gorm.DB) error {
		return tx.WithContext(ctx).Session(&gorm.Session{FullSaveAssociations: true}).Create(data).Error
	})

	return data.ID, err
}

func (Executor) Read(ctx context.Context, id int64) (*entities.Executor, error) {
	result := &entities.Executor{}
	err := clients.MySQLClient.WithContext(ctx).Preload("Representer.Members").Where("id = ?", id).First(result).Error
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (Executor) Update(ctx context.Context, data *entities.Executor) error {
	return clients.MySQLClient.Transaction(func(tx *gorm.DB) error {
		return tx.WithContext(ctx).Model(data).Session(&gorm.Session{FullSaveAssociations: true}).Where("id = ?", data.ID).Updates(data).Error
	})
}

func (Executor) Delete(ctx context.Context, id int64) error {
	return clients.MySQLClient.Transaction(func(tx *gorm.DB) error {
		return tx.WithContext(ctx).Model(&entities.Executor{}).Session(&gorm.Session{FullSaveAssociations: true}).Where("id = ?", id).Delete(&entities.Executor{}).Error
	})
}

func (Executor) List(ctx context.Context) ([]*entities.Executor, error) {
	result := []*entities.Executor{}
	err := clients.MySQLClient.WithContext(ctx).Preload("Representer.Members").Find(&result).Error
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (Executor) ListWithCondition(ctx context.Context, field string, value any) ([]*entities.Executor, error) {
	result := []*entities.Executor{}
	err := clients.MySQLClient.WithContext(ctx).Preload(clause.Associations).Where(fmt.Sprintf("%s = ?", field), value).Find(&result).Error
	if err != nil {
		return nil, err
	}

	return result, nil
}
