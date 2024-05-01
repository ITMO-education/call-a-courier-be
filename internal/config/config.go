package config

import (
	"github.com/godverv/matreshka"
)

const (
	devConfigPath  = "./config/dev.yaml"
	prodConfigPath = "./config/config.yaml"
)

var defaultConfig *config

type config struct {
	matreshka.AppConfig
}

func GetConfig() matreshka.Config {
	return defaultConfig
}

func (c *config) AppInfo() matreshka.AppInfo {
	return c.AppConfig.AppInfo
}

func (c *config) Api() matreshka.API {
	return &c.AppConfig.Servers
}

func (c *config) Resources() matreshka.Resource {
	return &c.AppConfig.Resources
}

func (c *config) GetMatreshka() *matreshka.AppConfig {
	return &c.AppConfig
}
