package models

import (
	"context"

	"github.com/ideal-forward/assistant-public-api/clients"
	"github.com/ideal-forward/assistant-public-api/entities"
	"gorm.io/gorm"
)

type IRegion interface {
	Create(ctx context.Context, data *entities.Region) (int64, error)
	Read(ctx context.Context, id int64) (*entities.Region, error)
	Update(ctx context.Context, data *entities.Region) error
	Delete(ctx context.Context, id int64) error
	List(ctx context.Context) ([]*entities.Region, error)
}

type Region struct {
}

func (Region) Create(ctx context.Context, data *entities.Region) (int64, error) {
	err := clients.MySQLClient.WithContext(ctx).Create(data).Error
	return data.ID, err
}

func (Region) Read(ctx context.Context, id int64) (*entities.Region, error) {
	result := &entities.Region{}
	err := clients.MySQLClient.WithContext(ctx).Where("id = ?", id).First(result).Error
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (Region) Update(ctx context.Context, data *entities.Region) error {
	return clients.MySQLClient.WithContext(ctx).Model(data).Session(&gorm.Session{FullSaveAssociations: true}).Where("id = ?", data.ID).Updates(data).Error
}

func (Region) Delete(ctx context.Context, id int64) error {
	return clients.MySQLClient.WithContext(ctx).Model(&entities.Region{}).Session(&gorm.Session{FullSaveAssociations: true}).Where("id = ?", id).Delete(&entities.Region{}).Error
}

func (Region) List(ctx context.Context) ([]*entities.Region, error) {
	result := []*entities.Region{}
	err := clients.MySQLClient.WithContext(ctx).Model(&entities.Region{}).Find(&result).Error
	if err != nil {
		return nil, err
	}

	return result, nil
}
