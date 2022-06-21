package config

import (
	"time"

	"github.com/sderohan/go-auth-server/pkg/utils"
)

type serverConfig struct {
	Address      string
	WriteTimeout time.Duration
	ReadTimeout  time.Duration
}

var srvConfig serverConfig

func initServerConfig() {
	srvConfig = serverConfig{
		Address:      utils.GetString("SERVER_ADDRESS"),
		WriteTimeout: utils.GetDuration("SERVER_WRITE_TIMEOUT") * time.Second,
		ReadTimeout:  utils.GetDuration("SERVER_READ_TIMEOUT") * time.Second,
	}
}

func GetServerConfig() serverConfig {
	return srvConfig
}
