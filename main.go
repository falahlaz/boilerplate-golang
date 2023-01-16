package main

import (
	"os"

	"github.com/falahlaz/boilerplate-golang/pkg/constant"
	"github.com/falahlaz/boilerplate-golang/pkg/database"
	"github.com/falahlaz/boilerplate-golang/pkg/env"
	"github.com/labstack/echo/v4"
)

func Init() {
	e := env.NewEnv()
	e.Load()
}

func main() {
	e := echo.New()

	// Init()
	database.Init()

	e.Logger.Fatal(e.Start(":" + os.Getenv(constant.APP_PORT)))
}
