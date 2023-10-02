package models

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/ideal-forward/assistant-public-api/clients"
	"github.com/morkid/paginate"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Storage[T any] struct {
}

func (s Storage[T]) List(ctx context.Context, data any) (any, error) {
	pg := paginate.New()
	model := clients.MySQLClient.Model(new(T)).Preload(clause.Associations)
	result := pg.With(model).Request(ctx.(*gin.Context).Request).Response(new([]T))
	return result, nil
}

func (s Storage[T]) Create(ctx context.Context, data any) (any, error) {
	model := data.(T)
	err := clients.MySQLClient.Transaction(func(tx *gorm.DB) error {
		return tx.WithContext(ctx).Session(&gorm.Session{FullSaveAssociations: true}).Create(model).Error
	})

	return model, err
}

func (s Storage[T]) Read(ctx context.Context, data any) (any, error) {
	model := data.(T)
	result := clients.MySQLClient.Model(new(T)).Preload(clause.Associations).First(model)
	if err := result.Error; err != nil {
		return nil, err
	}

	return model, nil
}

func (s Storage[T]) Update(ctx context.Context, id int64, data any) error {
	model := data.(T)
	return clients.MySQLClient.Transaction(func(tx *gorm.DB) error {
		return tx.WithContext(ctx).Model(new(T)).Session(&gorm.Session{FullSaveAssociations: true}).Save(model).Where("id = ?", id).Error
	})
}

func (s Storage[T]) Delete(ctx context.Context, id int64) error {
	model := new(T)
	return clients.MySQLClient.Transaction(func(tx *gorm.DB) error {
		return tx.WithContext(ctx).Model(model).Session(&gorm.Session{FullSaveAssociations: true}).Delete(model).Where("id = ?", id).Error
	})
}
