package models

import (
	"context"
	"fmt"

	"github.com/ideal-forward/assistant-public-api/clients"
	"github.com/ideal-forward/assistant-public-api/entities"
	"github.com/ideal-forward/assistant-public-api/pkg/util"
	"gorm.io/gorm"
)

type ITask interface {
	Create(ctx context.Context, data *entities.Task) (int64, error)
	Read(ctx context.Context, id int64) (*entities.Task, error)
	Update(ctx context.Context, data *entities.Task) error
	UpdateFields(ctx context.Context, id int64, data map[string]interface{}) error
	Delete(ctx context.Context, id int64) error
	List(ctx context.Context, filter map[string]any, getRefInfo bool) ([]*entities.Task, error)
	ListByCons(ctx context.Context, filters map[string]any, getRefInfo bool) ([]*entities.Task, error)
}

type Task struct {
}

func (Task) Create(ctx context.Context, data *entities.Task) (int64, error) {
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

func (Task) Read(ctx context.Context, id int64) (*entities.Task, error) {
	result := &entities.Task{}
	db := clients.MySQLClient.WithContext(ctx).
		Preload("Project").
		Preload("Area").
		Preload("Phase").
		Preload("ParentTask").
		Preload("ChildTasks").
		Preload("Creator").
		Preload("Executor").
		Preload("Reporter").
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

func (Task) Update(ctx context.Context, data *entities.Task) error {
	return clients.MySQLClient.Transaction(func(tx *gorm.DB) error {
		return tx.WithContext(ctx).Model(data).Session(&gorm.Session{FullSaveAssociations: true}).Where("id = ?", data.ID).Updates(data).Error
	})
}

func (Task) UpdateFields(ctx context.Context, id int64, data map[string]interface{}) error {
	return clients.MySQLClient.WithContext(ctx).Model(&entities.Task{}).Session(&gorm.Session{FullSaveAssociations: true}).Where("id = ?", id).Updates(data).Error
}

func (Task) Delete(ctx context.Context, id int64) error {
	return clients.MySQLClient.Transaction(func(tx *gorm.DB) error {
		return tx.WithContext(ctx).Model(&entities.Task{}).Session(&gorm.Session{FullSaveAssociations: true}).Where("id = ?", id).Delete(&entities.Task{}).Error
	})
}

func (Task) List(ctx context.Context, filters map[string]any, getRefInfo bool) ([]*entities.Task, error) {
	result := []*entities.Task{}
	db := clients.MySQLClient.WithContext(ctx)
	if getRefInfo {
		db = db.Preload("Project").
			Preload("Area").
			Preload("Phase").
			Preload("ParentTask").
			Preload("Creator").
			Preload("Executor").
			Preload("Reporter").
			Preload("Acceptor").
			Preload("AttachFiles").
			Preload("Comments").
			Preload("AssignHistories").
			Preload("StatusHistories").
			Preload("Creator")
	}

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

func (Task) ListByCons(ctx context.Context, filters map[string]any, withRelativeInfo bool) ([]*entities.Task, error) {
	result := []*entities.Task{}
	db := clients.MySQLClient.WithContext(ctx)
	if withRelativeInfo {
		db = db.Preload("Project").
			Preload("Area").
			Preload("Phase").
			Preload("ParentTask").
			Preload("Creator").
			Preload("Executor").
			Preload("Acceptor").
			Preload("AttachFiles").
			Preload("Comments").
			Preload("AssignHistories").
			Preload("StatusHistories").
			Preload("Creator")
	}

	for field, values := range filters {
		if field == "start_time" {
			db = db.Where("start_time >= ?", values)
		} else if field == "end_time" {
			db = db.Where("end_time <= ?", values)
		} else {
			db = db.Where(fmt.Sprintf("%s IN ?", field), values)
		}
	}

	err := db.Find(&result).Error
	if err != nil {
		return nil, err
	}

	return result, nil
}
