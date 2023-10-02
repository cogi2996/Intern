package services

import (
	"fmt"

	"github.com/ideal-forward/assistant-public-api/pkg/jwt"
	"github.com/spf13/viper"
)

var (
	TokenMaker *jwt.JWTMaker
	Config     *ServerConfig
)

func InitGlobalServices() {
	TokenMaker = &jwt.JWTMaker{
		SecretKey: "golanginuse-secret-key",
		Lifetime:  4320,
		Issuer:    "assistant-public-api",
	}

	Config = &ServerConfig{
		Domain: viper.GetString("service.domain"),
		Format: &FormatConfig{
			DateFormat: "2006-01-02",
			TimeFormat: "2006-01-02 15:04:05",
		},
		FileStorage: &FileStorageConfig{
			URLPath:        "/api/v1/images",
			Folder:         "./public/images/",
			ThumbnailWidth: 200,
		},
	}

	fmt.Printf("INFO: Config.Domain %s \n", Config.Domain)
}
