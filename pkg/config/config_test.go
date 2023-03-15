package config

import (
	"fmt"
	"os"
	"testing"

	"github.com/falahlaz/boilerplate-golang/pkg/config/entity"
)

func TestConfigConfigor(t *testing.T) {
	env := os.Getenv("APP_ENV")
	conf := os.Getenv("CONFIG")
	confPGP := os.Getenv("CONFIGPGP")

	path := ""
	pathPGP := ""

	switch env {
	case "development":
		path = fmt.Sprintf("/run/secrets/%s", conf)
		pathPGP = fmt.Sprintf("/run/secrets/%s", confPGP)
	case "staging":
		path = fmt.Sprintf("/run/secrets/%s", conf)
		pathPGP = fmt.Sprintf("/run/secrets/%s", confPGP)
	case "production":
		path = fmt.Sprintf("/run/secrets/%s", conf)
		pathPGP = fmt.Sprintf("/run/secrets/%s", confPGP)
	default:
		conf = "config.yml"
		confPGP = ""
		path = conf
		pathPGP = confPGP
	}

	// init configuration
	// defaultConfig := entity.NewConfigData()

	// load configuration
	configConfigor := NewConfigor(path, pathPGP)
	configConfigor.Load(&entity.ConfigData{})
}
