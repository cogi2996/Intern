package models

import (
	"context"

	"github.com/ideal-forward/assistant-public-api/clients"
	"github.com/ideal-forward/assistant-public-api/entities"
	"gorm.io/gorm"
)

type IProjectCategory interface {
	Create(ctx context.Context, data *entities.ProjectCategory) (int64, error)
	Read(ctx context.Context, id int64) (*entities.ProjectCategory, error)
	Update(ctx context.Context, data *entities.ProjectCategory) error
	Delete(ctx context.Context, id int64) error
	List(ctx context.Context) ([]*entities.ProjectCategory, error)
}

type ProjectCategory struct {
}

func (ProjectCategory) Create(ctx context.Context, data *entities.ProjectCategory) (int64, error) {
	err := clients.MySQLClient.WithContext(ctx).Create(data).Error
	return data.ID, err
}

func (ProjectCategory) Read(ctx context.Context, id int64) (*entities.ProjectCategory, error) {
	result := &entities.ProjectCategory{}
	err := clients.MySQLClient.WithContext(ctx).Where("id = ?", id).First(result).Error
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (ProjectCategory) Update(ctx context.Context, data *entities.ProjectCategory) error {
	return clients.MySQLClient.WithContext(ctx).Model(data).Session(&gorm.Session{FullSaveAssociations: true}).Where("id = ?", data.ID).Updates(data).Error
}

func (ProjectCategory) Delete(ctx context.Context, id int64) error {
	return clients.MySQLClient.WithContext(ctx).Model(&entities.ProjectCategory{}).Session(&gorm.Session{FullSaveAssociations: true}).Where("id = ?", id).Delete(&entities.ProjectCategory{}).Error
}

func (ProjectCategory) List(ctx context.Context) ([]*entities.ProjectCategory, error) {
	result := []*entities.ProjectCategory{}
	err := clients.MySQLClient.WithContext(ctx).Model(&entities.ProjectCategory{}).Find(&result).Error
	if err != nil {
		return nil, err
	}

	return result, nil
}
