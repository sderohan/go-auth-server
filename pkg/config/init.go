package config

import (
	"github.com/sderohan/go-auth-server/pkg/utils"
)

func LoadConfigFile() {
	var configFile string = "application"
	utils.SetConfigFileName(configFile)
	utils.SetConfigFileType("env")
	utils.SetConfigFileSearchPath(".")
	utils.ReadConfig()
}

func InitConfigs() {
	LoadConfigFile()
	initServerConfig()
}
