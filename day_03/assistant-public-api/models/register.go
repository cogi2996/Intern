package models

import (
	"context"

	"github.com/ideal-forward/assistant-public-api/clients"
	"github.com/ideal-forward/assistant-public-api/entities"
	"gorm.io/gorm"
)

type IRegister interface {
	Create(ctx context.Context, data *entities.Register) (int64, error)
	Read(ctx context.Context, id int64) (*entities.Register, error)
	Update(ctx context.Context, data *entities.Register) error
	Delete(ctx context.Context, id int64) error

	ReadByUsername(ctx context.Context, username string) (*entities.Register, error)
}

type Register struct {
}

func (Register) Create(ctx context.Context, data *entities.Register) (int64, error) {
	err := clients.MySQLClient.Transaction(func(tx *gorm.DB) error {
		return tx.WithContext(ctx).Session(&gorm.Session{FullSaveAssociations: true}).Create(data).Error
	})

	return data.ID, err
}

func (Register) Read(ctx context.Context, id int64) (*entities.Register, error) {
	result := &entities.Register{}
	err := clients.MySQLClient.WithContext(ctx).Where("id = ?", id).First(result).Error
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (Register) Update(ctx context.Context, data *entities.Register) error {
	return clients.MySQLClient.Transaction(func(tx *gorm.DB) error {
		return tx.WithContext(ctx).Model(data).Session(&gorm.Session{FullSaveAssociations: true}).Save(data).Where("id = ?", data.ID).Error
	})
}

func (Register) Delete(ctx context.Context, id int64) error {
	return clients.MySQLClient.Transaction(func(tx *gorm.DB) error {
		return tx.WithContext(ctx).Model(&entities.Register{}).Session(&gorm.Session{FullSaveAssociations: true}).Delete(&entities.Register{}).Where("id = ?", id).Error
	})
}

func (Register) ReadByUsername(ctx context.Context, username string) (*entities.Register, error) {
	result := &entities.Register{}
	err := clients.MySQLClient.WithContext(ctx).Where("username = ?", username).First(result).Error
	if err != nil {
		return nil, err
	}

	return result, nil
}
