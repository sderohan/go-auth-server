package utils

import (
	"fmt"
	"strings"
	"time"

	"github.com/sderohan/go-auth-server/pkg/constants"
	"github.com/spf13/viper"
)

func SetConfigFileName(filename string) {
	viper.SetConfigName(filename)
}

func SetConfigFileType(filetype string) {
	viper.SetConfigType(filetype)
}

func SetConfigFileSearchPath(path string) {
	viper.AddConfigPath(path)
}

func ReadConfig() {
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
}

func GetString(key string) string {
	value := viper.GetString(key)
	if strings.TrimSpace(value) == "" {
		panic(fmt.Sprintf(constants.CONFIG_KEY_NOT_SET, key))
	}
	return value
}

func GetDuration(key string) time.Duration {
	value := viper.GetDuration(key)
	if value == 0 {
		panic(fmt.Sprintf(constants.CONFIG_KEY_NOT_SET, key))
	}
	return value
}
