package main

import (
	"fmt"
	"os"

	"github.com/falahlaz/boilerplate-golang/pkg/config"
	"github.com/falahlaz/boilerplate-golang/pkg/config/entity"
	"github.com/falahlaz/boilerplate-golang/pkg/constant"
	"github.com/falahlaz/boilerplate-golang/pkg/database"
	"github.com/labstack/echo/v4"
)

func init() {
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
	defaultConfig := entity.NewConfigData()

	// load configuration
	configConfigor := config.NewConfigor(path, pathPGP)
	configConfigor.Load(defaultConfig)
	config.Config = *defaultConfig
}

func main() {
	e := echo.New()

	database.Init()

	e.Logger.Fatal(e.Start(":" + os.Getenv(constant.APP_PORT)))
}
