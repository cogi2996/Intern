package models

import (
	"context"

	"github.com/ideal-forward/assistant-public-api/clients"
	"github.com/ideal-forward/assistant-public-api/entities"
	"gorm.io/gorm"
)

type ITimesheetsExecutor interface {
	Create(ctx context.Context, data *entities.TimesheetsExecutor) (int64, error)
	Read(ctx context.Context, id int64) (*entities.TimesheetsExecutor, error)
	Update(ctx context.Context, data *entities.TimesheetsExecutor) error
	Delete(ctx context.Context, id int64) error
	List(ctx context.Context, timesheetsID int64) ([]*entities.TimesheetsExecutor, error)
}

type TimesheetsExecutor struct {
}

func (TimesheetsExecutor) Create(ctx context.Context, data *entities.TimesheetsExecutor) (int64, error) {
	err := clients.MySQLClient.WithContext(ctx).Create(data).Error
	return data.ID, err
}

func (TimesheetsExecutor) Read(ctx context.Context, id int64) (*entities.TimesheetsExecutor, error) {
	result := &entities.TimesheetsExecutor{}
	err := clients.MySQLClient.WithContext(ctx).Where("id = ?", id).First(result).Error
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (TimesheetsExecutor) Update(ctx context.Context, data *entities.TimesheetsExecutor) error {
	return clients.MySQLClient.Transaction(func(tx *gorm.DB) error {
		return tx.WithContext(ctx).Model(data).Session(&gorm.Session{FullSaveAssociations: true}).Where("id = ?", data.ID).Updates(data).Error
	})
}

func (TimesheetsExecutor) Delete(ctx context.Context, id int64) error {
	return clients.MySQLClient.Transaction(func(tx *gorm.DB) error {
		return tx.WithContext(ctx).Model(&entities.TimesheetsExecutor{}).Session(&gorm.Session{FullSaveAssociations: true}).Where("id = ?", id).Delete(&entities.TimesheetsExecutor{}).Error
	})
}

func (TimesheetsExecutor) List(ctx context.Context, timesheetsID int64) ([]*entities.TimesheetsExecutor, error) {
	result := []*entities.TimesheetsExecutor{}
	err := clients.MySQLClient.WithContext(ctx).Model(&entities.TimesheetsExecutor{}).
		Preload("Creator").
		Preload("Updater").
		Preload("Executor").
		Where("timesheets_id = ?", timesheetsID).Find(&result).Error
	if err != nil {
		return nil, err
	}

	return result, nil
}
