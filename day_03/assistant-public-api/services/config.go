package services

import (
	"fmt"
	"os"

	"github.com/fsnotify/fsnotify"
	"github.com/ideal-forward/assistant-public-api/pkg/util"
	"github.com/spf13/viper"
)

type FormatConfig struct {
	DateFormat string
	TimeFormat string
}

type FileStorageConfig struct {
	URLPath        string
	Folder         string
	ThumbnailWidth int
}

type ServerConfig struct {
	Domain      string
	Format      *FormatConfig
	FileStorage *FileStorageConfig
}

func InitViper() {

	cfgFile := "./config.toml"
	fmt.Printf("Read config file from env: [%s] \n", cfgFile)

	folder, fileName, ext, err := util.ExtractFilePath(cfgFile)
	if err != nil {
		fmt.Printf("Extract config file failed %s err: %s \n", viper.ConfigFileUsed(), err.Error())
		os.Exit(-1)
	}
	fmt.Printf("Extract config file success folder[%s] fileName[%s] ext[%s] \n", folder, fileName, ext)

	// Setting
	viper.AddConfigPath(folder)
	viper.SetConfigName(fileName)
	viper.AutomaticEnv()
	viper.SetConfigType(ext)

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("FATAL: Viper using config file failed %s err: %s \n", viper.ConfigFileUsed(), err.Error())
		os.Exit(-1)
	}

	fmt.Printf("Service using config file: %s \n", viper.ConfigFileUsed())
	//watch on config change
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Printf("Config file changed: %s", e.Name)
	})

	fmt.Println("Start initialize config success.")
}
