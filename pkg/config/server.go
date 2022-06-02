package config

import (
	"time"

	"github.com/sderohan/go-auth-server/pkg/utils"
)

type ServerConfig struct {
	Address      string
	WriteTimeout time.Duration
	ReadTimeout  time.Duration
}

var serverConfig ServerConfig

func initServerConfig() {
	serverConfig = ServerConfig{
		Address:      utils.GetString("SERVER_ADDRESS"),
		WriteTimeout: utils.GetDuration("SERVER_WRITE_TIMEOUT") * time.Second,
		ReadTimeout:  utils.GetDuration("SERVER_READ_TIMEOUT") * time.Second,
	}
}

func GetServerConfig() ServerConfig {
	return serverConfig
}
