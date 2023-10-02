package models

import (
	"context"

	"github.com/ideal-forward/assistant-public-api/clients"
	"github.com/ideal-forward/assistant-public-api/entities"
	"gorm.io/gorm"
)

type ITimesheetsComment interface {
	Create(ctx context.Context, data *entities.TimesheetsComment) (int64, error)
	Read(ctx context.Context, id int64) (*entities.TimesheetsComment, error)
	Update(ctx context.Context, data *entities.TimesheetsComment) error
	Delete(ctx context.Context, id int64) error
	List(ctx context.Context, timesheetsID int64) ([]*entities.TimesheetsComment, error)
}

type TimesheetsComment struct {
}

func (TimesheetsComment) Create(ctx context.Context, data *entities.TimesheetsComment) (int64, error) {
	err := clients.MySQLClient.WithContext(ctx).Create(data).Error
	return data.ID, err
}

func (TimesheetsComment) Read(ctx context.Context, id int64) (*entities.TimesheetsComment, error) {
	result := &entities.TimesheetsComment{}
	err := clients.MySQLClient.WithContext(ctx).Where("id = ?", id).First(result).Error
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (TimesheetsComment) Update(ctx context.Context, data *entities.TimesheetsComment) error {
	return clients.MySQLClient.Transaction(func(tx *gorm.DB) error {
		return tx.WithContext(ctx).Model(data).Session(&gorm.Session{FullSaveAssociations: true}).Where("id = ?", data.ID).Updates(data).Error
	})
}

func (TimesheetsComment) Delete(ctx context.Context, id int64) error {
	return clients.MySQLClient.Transaction(func(tx *gorm.DB) error {
		return tx.WithContext(ctx).Model(&entities.TimesheetsComment{}).Session(&gorm.Session{FullSaveAssociations: true}).Where("id = ?", id).Delete(&entities.TimesheetsComment{}).Error
	})
}

func (TimesheetsComment) List(ctx context.Context, timesheetsID int64) ([]*entities.TimesheetsComment, error) {
	result := []*entities.TimesheetsComment{}
	err := clients.MySQLClient.WithContext(ctx).Model(&entities.TimesheetsComment{}).Preload("Creator").Where("timesheets_id = ?", timesheetsID).Find(&result).Error
	if err != nil {
		return nil, err
	}

	return result, nil
}
