package models

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/ideal-forward/assistant-public-api/clients"
	"github.com/ideal-forward/assistant-public-api/entities"
	"gorm.io/gorm"
)

type IProject interface {
	Create(ctx context.Context, data *entities.Project) (int64, error)
	Read(ctx context.Context, id int64) (*entities.Project, error)
	Update(ctx context.Context, data *entities.Project) error
	Delete(ctx context.Context, id int64) error
	GetList(ctx context.Context) ([]*entities.Project, error)
	GetListByCondition(ctx context.Context, field string, value any) ([]*entities.Project, error)
	GetListByExecutor(ctx context.Context, executorID int64) ([]*entities.Project, error)

	AddMember(ctx context.Context, data *entities.ProjectMember) error
	DeleteMember(ctx context.Context, data *entities.ProjectMember) error
	GetListMember(ctx context.Context, projectID int64) ([]*entities.ProjectMember, error)

	AddExecutor(ctx context.Context, data *entities.ProjectExecutor) error
	DeleteExecutor(ctx context.Context, data *entities.ProjectExecutor) error
	GetListExecutor(ctx context.Context, projectID int64) ([]*entities.ProjectExecutor, error)

	CreatePhase(ctx context.Context, data *entities.ProjectPhase) error
	ReadPhase(ctx context.Context, id int64) (*entities.ProjectPhase, error)
	UpdatePhase(ctx context.Context, data *entities.ProjectPhase) error
	DeletePhase(ctx context.Context, data *entities.ProjectPhase) error
	GetListPhase(ctx context.Context, projectID int64) ([]*entities.ProjectPhase, error)

	CreateArea(ctx context.Context, data *entities.ProjectArea) error
	ReadArea(ctx context.Context, id int64) (*entities.ProjectArea, error)
	UpdateArea(ctx context.Context, data *entities.ProjectArea) error
	DeleteArea(ctx context.Context, data *entities.ProjectArea) error
	GetListArea(ctx context.Context, projectID int64) ([]*entities.ProjectArea, error)
}

type Project struct {
}

func (Project) Create(ctx context.Context, data *entities.Project) (int64, error) {
	err := clients.MySQLClient.Transaction(func(tx *gorm.DB) error {
		return tx.WithContext(ctx).Session(&gorm.Session{FullSaveAssociations: true}).Create(data).Error
	})

	return data.ID, err
}

func (Project) Read(ctx context.Context, id int64) (*entities.Project, error) {
	result := &entities.Project{}
	db := clients.MySQLClient.WithContext(ctx).
		Preload("Manager").
		Preload("Category").
		Preload("Region").
		Preload("Phases").
		Preload("Areas").
		Preload("Tasks.ParentTask").
		Preload("Tasks.Executor").
		Preload("Tasks.Acceptor").
		Preload("Tasks.AttachFiles").
		Preload("Tasks.Comments").
		Preload("Tasks.AssignHistories").
		Preload("Tasks.Creator")
	err := db.Where("id = ?", id).First(result).Error
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (Project) Update(ctx context.Context, data *entities.Project) error {
	return clients.MySQLClient.Transaction(func(tx *gorm.DB) error {
		return tx.WithContext(ctx).Model(data).Session(&gorm.Session{FullSaveAssociations: true}).Where("id = ?", data.ID).Updates(data).Error
	})
}

func (Project) Delete(ctx context.Context, id int64) error {
	return clients.MySQLClient.Transaction(func(tx *gorm.DB) error {
		return tx.WithContext(ctx).Model(&entities.Project{}).Session(&gorm.Session{FullSaveAssociations: true}).Where("id = ?", id).Delete(&entities.Project{}).Error
	})
}

func (Project) GetList(ctx context.Context) ([]*entities.Project, error) {
	result := []*entities.Project{}
	db := clients.MySQLClient.WithContext(ctx).
		Preload("Manager").
		Preload("Category").
		Preload("Region").
		Preload("Phases").
		Preload("Areas").
		Preload("Tasks.ParentTask").
		Preload("Tasks.Executor").
		Preload("Tasks.Acceptor").
		Preload("Tasks.AttachFiles").
		Preload("Tasks.Comments").
		Preload("Tasks.AssignHistories").
		Preload("Tasks.Creator")
	err := db.Model(&entities.Project{}).Find(&result).Error
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (Project) GetListByCondition(ctx context.Context, field string, value any) ([]*entities.Project, error) {
	result := []*entities.Project{}
	db := clients.MySQLClient.WithContext(ctx).
		Preload("Tasks.ParentTask").
		Preload("Tasks.Executor").
		Preload("Tasks.Acceptor").
		Preload("Tasks.AttachFiles").
		Preload("Tasks.Comments").
		Preload("Tasks.AssignHistories").
		Preload("Tasks.Creator")
	err := db.Model(&entities.Project{}).Where(fmt.Sprintf("%s = ?", field), value).Find(&result).Error
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (Project) GetListByExecutor(ctx context.Context, executorID int64) ([]*entities.Project, error) {
	result := []*entities.Project{}
	db := clients.MySQLClient.WithContext(ctx).
		Preload("Tasks.ParentTask").
		Preload("Tasks.Executor", "id = ?", executorID).
		Preload("Tasks.Acceptor").
		Preload("Tasks.AttachFiles").
		Preload("Tasks.Comments").
		Preload("Tasks.AssignHistories").
		Preload("Tasks.Creator")
	err := db.Model(&entities.Project{}).Find(&result).Error
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (Project) AddMember(ctx context.Context, data *entities.ProjectMember) error {
	return clients.MySQLClient.WithContext(ctx).Create(data).Error
}

func (Project) DeleteMember(ctx context.Context, data *entities.ProjectMember) error {
	return clients.MySQLClient.WithContext(ctx).Model(&entities.ProjectMember{}).Where("project_id = ? AND member_id = ?", data.ProjectID, data.MemberID).Delete(&entities.ProjectMember{}).Error
}

func (Project) GetListMember(ctx context.Context, projectID int64) ([]*entities.ProjectMember, error) {
	result := []*entities.ProjectMember{}
	db := clients.MySQLClient.WithContext(ctx).Preload("Project").Preload("Member")
	err := db.Model(&entities.ProjectMember{}).Where("project_id = ?", projectID).Find(&result).Error
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (Project) AddExecutor(ctx context.Context, data *entities.ProjectExecutor) error {
	return clients.MySQLClient.WithContext(ctx).Create(data).Error
}

func (Project) DeleteExecutor(ctx context.Context, data *entities.ProjectExecutor) error {
	return clients.MySQLClient.WithContext(ctx).Model(&entities.ProjectExecutor{}).Where("project_id = ? AND executor_id = ?", data.ProjectID, data.ExecutorID).Delete(&entities.ProjectExecutor{}).Error
}

func (Project) GetListExecutor(ctx context.Context, projectID int64) ([]*entities.ProjectExecutor, error) {
	result := []*entities.ProjectExecutor{}
	db := clients.MySQLClient.WithContext(ctx).Preload("Project").Preload("Executor")
	err := db.Model(&entities.ProjectExecutor{}).Where("project_id = ?", projectID).Find(&result).Error
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (Project) CreatePhase(ctx context.Context, data *entities.ProjectPhase) error {
	data.UUID = uuid.NewString()
	return clients.MySQLClient.WithContext(ctx).Create(data).Error
}

func (Project) ReadPhase(ctx context.Context, id int64) (*entities.ProjectPhase, error) {
	result := &entities.ProjectPhase{}
	err := clients.MySQLClient.WithContext(ctx).Preload("Project").Where("id = ?", id).First(result).Error
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (Project) UpdatePhase(ctx context.Context, data *entities.ProjectPhase) error {
	return clients.MySQLClient.WithContext(ctx).Model(data).Session(&gorm.Session{FullSaveAssociations: true}).Where("id = ?", data.ID).Updates(data).Error
}

func (Project) DeletePhase(ctx context.Context, data *entities.ProjectPhase) error {
	return clients.MySQLClient.WithContext(ctx).Model(&entities.ProjectPhase{}).Where("project_id = ? AND id = ?", data.ProjectID, data.ID).Delete(&entities.ProjectPhase{}).Error
}

func (Project) GetListPhase(ctx context.Context, projectID int64) ([]*entities.ProjectPhase, error) {
	result := []*entities.ProjectPhase{}
	db := clients.MySQLClient.WithContext(ctx).Preload("Project")
	err := db.Model(&entities.ProjectPhase{}).Where("project_id = ?", projectID).Find(&result).Error
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (Project) CreateArea(ctx context.Context, data *entities.ProjectArea) error {
	data.UUID = uuid.NewString()
	return clients.MySQLClient.WithContext(ctx).Create(data).Error
}

func (Project) ReadArea(ctx context.Context, id int64) (*entities.ProjectArea, error) {
	result := &entities.ProjectArea{}
	err := clients.MySQLClient.WithContext(ctx).Preload("Project").Where("id = ?", id).First(result).Error
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (Project) UpdateArea(ctx context.Context, data *entities.ProjectArea) error {
	return clients.MySQLClient.WithContext(ctx).Model(data).Session(&gorm.Session{FullSaveAssociations: true}).Where("id = ?", data.ID).Updates(data).Error
}

func (Project) DeleteArea(ctx context.Context, data *entities.ProjectArea) error {
	return clients.MySQLClient.WithContext(ctx).Model(&entities.ProjectArea{}).Where("project_id = ? AND id = ?", data.ProjectID, data.ID).Delete(&entities.ProjectArea{}).Error
}

func (Project) GetListArea(ctx context.Context, projectID int64) ([]*entities.ProjectArea, error) {
	result := []*entities.ProjectArea{}
	db := clients.MySQLClient.WithContext(ctx).Preload("Project")
	err := db.Model(&entities.ProjectArea{}).Where("project_id = ?", projectID).Find(&result).Error
	if err != nil {
		return nil, err
	}

	return result, nil
}
