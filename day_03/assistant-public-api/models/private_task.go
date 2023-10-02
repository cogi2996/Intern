package models

import (
	"context"
	"fmt"

	"github.com/ideal-forward/assistant-public-api/clients"
	"github.com/ideal-forward/assistant-public-api/entities"
	"github.com/ideal-forward/assistant-public-api/pkg/util"
	"gorm.io/gorm"
)

type IPrivateTask interface {
	Create(ctx context.Context, data *entities.PrivateTask) (int64, error)
	Read(ctx context.Context, id int64) (*entities.PrivateTask, error)
	Update(ctx context.Context, data *entities.PrivateTask) error
	UpdateField(ctx context.Context, userID, id int64, field string, value any) error
	Delete(ctx context.Context, id int64) error
	List(ctx context.Context, filter map[string]any) ([]*entities.PrivateTask, error)
}

type PrivateTask struct {
}

func (PrivateTask) Create(ctx context.Context, data *entities.PrivateTask) (int64, error) {
	err := clients.MySQLClient.Transaction(func(tx *gorm.DB) error {
		txErr := tx.WithContext(ctx).Session(&gorm.Session{FullSaveAssociations: true}).Create(data).Error
		if txErr != nil {
			return txErr
		}

		data.Code = util.MakeCodeByYear(data.ID)
		return tx.WithContext(ctx).Model(&data).Where("id = ?", data.ID).Update("code", data.Code).Error
	})

	return data.ID, err
}

func (PrivateTask) UpdateField(ctx context.Context, userID, id int64, field string, value any) error {
	return clients.MySQLClient.WithContext(ctx).Model(&entities.PrivateTask{}).Session(&gorm.Session{FullSaveAssociations: true}).Where("id = ?", id).Update(field, value).Error
}

func (PrivateTask) Read(ctx context.Context, id int64) (*entities.PrivateTask, error) {
	result := &entities.PrivateTask{}
	db := clients.MySQLClient.WithContext(ctx).
		Preload("Project").
		Preload("Area").
		Preload("Phase").
		Preload("ParentTask").
		Preload("Creator").
		Preload("Executor").
		Preload("Acceptor").
		Preload("AttachFiles").
		Preload("Comments").
		Preload("Comments.Creator").
		Preload("AssignHistories").
		Preload("StatusHistories").
		Preload("Creator")
	err := db.Where("id = ?", id).First(result).Error
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (PrivateTask) Update(ctx context.Context, data *entities.PrivateTask) error {
	return clients.MySQLClient.Transaction(func(tx *gorm.DB) error {
		return tx.WithContext(ctx).Model(data).Session(&gorm.Session{FullSaveAssociations: true}).Where("id = ?", data.ID).Updates(data).Error
	})
}

func (PrivateTask) Delete(ctx context.Context, id int64) error {
	return clients.MySQLClient.Transaction(func(tx *gorm.DB) error {
		return tx.WithContext(ctx).Model(&entities.PrivateTask{}).Session(&gorm.Session{FullSaveAssociations: true}).Where("id = ?", id).Delete(&entities.PrivateTask{}).Error
	})
}

func (PrivateTask) List(ctx context.Context, filters map[string]any) ([]*entities.PrivateTask, error) {
	result := []*entities.PrivateTask{}
	db := clients.MySQLClient.WithContext(ctx).
		Preload("Project").
		Preload("Area").
		Preload("Phase").
		Preload("ParentTask").
		Preload("ChildTasks").
		Preload("Creator").
		Preload("Executor").
		Preload("Acceptor").
		Preload("AttachFiles").
		Preload("Comments").
		Preload("AssignHistories").
		Preload("StatusHistories").
		Preload("Creator")

	for field, value := range filters {
		if field == "start_time" {
			db = db.Where("start_time >= ?", value)
		} else if field == "end_time" {
			db = db.Where("end_time <= ?", value)
		} else {
			db = db.Where(fmt.Sprintf("%s = ?", field), value)
		}
	}

	err := db.Find(&result).Error
	if err != nil {
		return nil, err
	}

	return result, nil
}
