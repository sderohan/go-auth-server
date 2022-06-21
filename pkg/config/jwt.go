package config

import (
	"time"

	"github.com/sderohan/go-auth-server/pkg/utils"
)

type jwtConfig struct {
	SecretKey      string
	ExpiryDuration time.Duration
}

var jwtConf jwtConfig

func initJWTConfig() {
	jwtConf = jwtConfig{
		SecretKey:      utils.GetString("SECRET_KEY"),
		ExpiryDuration: utils.GetDuration("EXPIRY_DURATION"),
	}
}

func GetJWTConfig() jwtConfig {
	return jwtConf
}
