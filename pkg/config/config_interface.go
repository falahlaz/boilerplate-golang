package config

import "github.com/falahlaz/boilerplate-golang/pkg/config/entity"

type ConfigInterface interface {
	Load(data *entity.ConfigData) error
}

var Config entity.ConfigData
