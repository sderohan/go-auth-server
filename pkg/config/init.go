package config

import (
	"github.com/sderohan/go-auth-server/pkg/utils"
)

func loadConfigFile() {
	var configFile string = "application"
	utils.SetConfigFileName(configFile)
	utils.SetConfigFileType("env")
	utils.SetConfigFileSearchPath(".")
	utils.ReadConfig()
}

func InitConfigs() {
	loadConfigFile()
	initServerConfig()
}
