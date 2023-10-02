package clients

import (
	"fmt"

	"github.com/ideal-forward/assistant-public-api/entities"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	// "github.com/spf13/viper"
)

var MySQLClient *gorm.DB

func NewMySQLClient() (*gorm.DB, error) {
	username := viper.GetString("database.username")
	password := viper.GetString("database.password")
	addr := viper.GetString("database.address")
	dbname := viper.GetString("database.dbname")
	connStr := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&collation=utf8mb4_unicode_ci&parseTime=True",
		username, password, addr, dbname)

	var err error
	client, err := gorm.Open(mysql.Open(connStr), &gorm.Config{
		Logger:                                   logger.Default.LogMode(logger.Info),
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		return nil, err
	}

	client = client.Set("gorm:table_options", "ENGINE=InnoDB CHARSET=utf8mb4 auto_increment=1")

	// set data models many to many relationship
	/*
		err = client.SetupJoinTable(&Product{}, "Factories", &ProductFactory{})
		if err != nil {
			return nil, err
		}
	*/

	return client, nil
}

func AutoMigrate() error {

	err := MySQLClient.AutoMigrate(&entities.User{})
	if err != nil {
		return err
	}

	err = MySQLClient.AutoMigrate(&entities.Executor{})
	if err != nil {
		return err
	}

	err = MySQLClient.AutoMigrate(&entities.ProjectCategory{}, &entities.Region{}, &entities.Project{})
	if err != nil {
		return err
	}

	err = MySQLClient.AutoMigrate(&entities.ProjectMember{}, &entities.ProjectExecutor{}, &entities.ProjectPhase{}, &entities.ProjectArea{})
	if err != nil {
		return err
	}

	err = MySQLClient.AutoMigrate(&entities.Task{}, &entities.TaskComment{}, &entities.TaskAttachFile{}, &entities.TaskAssignHistory{})
	if err != nil {
		return err
	}

	err = MySQLClient.AutoMigrate(&entities.PrivateTask{}, &entities.PrivateTaskComment{}, &entities.PrivateTaskAttachFile{}, &entities.PrivateTaskAssignHistory{})
	if err != nil {
		return err
	}

	err = MySQLClient.AutoMigrate(&entities.AuditTask{}, &entities.AuditTaskComment{}, &entities.AuditTaskAttachFile{}, &entities.AuditTaskAssignHistory{})
	if err != nil {
		return err
	}

	err = MySQLClient.AutoMigrate(&entities.Notification{})
	if err != nil {
		return err
	}

	err = MySQLClient.AutoMigrate(&entities.Timesheets{}, &entities.TimesheetsExecutor{}, &entities.TimesheetsComment{})
	if err != nil {
		return err
	}

	return nil
}
