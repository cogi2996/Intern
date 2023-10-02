package models

import (
	"context"

	"github.com/ideal-forward/assistant-public-api/clients"
	"github.com/ideal-forward/assistant-public-api/entities"
	"gorm.io/gorm"
)

type ITimesheets interface {
	Create(ctx context.Context, data *entities.Timesheets) (int64, error)
	Read(ctx context.Context, id int64) (*entities.Timesheets, error)
	Update(ctx context.Context, data *entities.Timesheets) error
	Delete(ctx context.Context, id int64) error
	List(ctx context.Context, projectID int64) ([]*entities.Timesheets, error)
}

type Timesheets struct {
}

func (Timesheets) Create(ctx context.Context, data *entities.Timesheets) (int64, error) {
	err := clients.MySQLClient.WithContext(ctx).Create(data).Error
	return data.ID, err
}

func (Timesheets) Read(ctx context.Context, id int64) (*entities.Timesheets, error) {
	result := &entities.Timesheets{}
	err := clients.MySQLClient.WithContext(ctx).Preload("Creator").Preload("Project").Preload("Comments").Preload("Executors").Where("id = ?", id).First(result).Error
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (Timesheets) Update(ctx context.Context, data *entities.Timesheets) error {
	return clients.MySQLClient.Transaction(func(tx *gorm.DB) error {
		return tx.WithContext(ctx).Model(data).Session(&gorm.Session{FullSaveAssociations: true}).Where("id = ?", data.ID).Updates(data).Error
	})
}

func (Timesheets) Delete(ctx context.Context, id int64) error {
	return clients.MySQLClient.Transaction(func(tx *gorm.DB) error {
		return tx.WithContext(ctx).Model(&entities.Timesheets{}).Session(&gorm.Session{FullSaveAssociations: true}).Where("id = ?", id).Delete(&entities.Timesheets{}).Error
	})
}

func (Timesheets) List(ctx context.Context, projectID int64) ([]*entities.Timesheets, error) {
	result := []*entities.Timesheets{}
	err := clients.MySQLClient.WithContext(ctx).Model(&entities.Timesheets{}).Preload("Creator").Where("project_id = ?", projectID).Order("Date desc").Find(&result).Error
	if err != nil {
		return nil, err
	}

	return result, nil
}
