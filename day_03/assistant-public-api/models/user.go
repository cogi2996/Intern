package models

import (
	"context"

	"github.com/ideal-forward/assistant-public-api/clients"
	"github.com/ideal-forward/assistant-public-api/entities"
	"gorm.io/gorm"
)

type IUser interface {
	Create(ctx context.Context, data *entities.User) (int64, error)
	Read(ctx context.Context, id int64) (*entities.User, error)
	Update(ctx context.Context, data *entities.User) error
	Delete(ctx context.Context, id int64) error
	List(ctx context.Context) ([]*entities.User, error)
}

type User struct {
}

func (User) Create(ctx context.Context, data *entities.User) (int64, error) {
	err := clients.MySQLClient.WithContext(ctx).Create(data).Error
	return data.ID, err
}

func (User) Read(ctx context.Context, id int64) (*entities.User, error) {
	result := &entities.User{}
	err := clients.MySQLClient.WithContext(ctx).Preload("Manager").Where("id = ?", id).First(result).Error
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (User) Update(ctx context.Context, data *entities.User) error {
	return clients.MySQLClient.Transaction(func(tx *gorm.DB) error {
		return tx.WithContext(ctx).Model(data).Session(&gorm.Session{FullSaveAssociations: true}).Where("id = ?", data.ID).Updates(data).Error
	})
}

func (User) Delete(ctx context.Context, id int64) error {
	return clients.MySQLClient.Transaction(func(tx *gorm.DB) error {
		return tx.WithContext(ctx).Model(&entities.User{}).Session(&gorm.Session{FullSaveAssociations: true}).Where("id = ?", id).Delete(&entities.User{}).Error
	})
}

func (User) List(ctx context.Context) ([]*entities.User, error) {
	result := []*entities.User{}
	err := clients.MySQLClient.WithContext(ctx).Model(&entities.User{}).Preload("Manager").Preload("Executors").Find(&result).Error
	if err != nil {
		return nil, err
	}

	return result, nil
}
