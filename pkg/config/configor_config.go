package config

import (
	"github.com/falahlaz/boilerplate-golang/pkg/config/entity"
	"github.com/jinzhu/configor"
)

type ConfigorConfig struct {
	Data          *entity.ConfigData
	path          string
	pathEncrypted string
}

func NewConfigor(path string, pathEncrypted string) *ConfigorConfig {
	return &ConfigorConfig{
		path:          path,
		pathEncrypted: pathEncrypted,
	}
}

func (c *ConfigorConfig) Load(data *entity.ConfigData) error {
	c.Data = data
	err := configor.New(nil).Load(c.Data, c.path)
	if err != nil {
		return err
	}
	return nil
}
