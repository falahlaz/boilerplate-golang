package database

import (
	"fmt"
	"os"

	"github.com/falahlaz/boilerplate-golang/pkg/constant"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var (
	conn *gorm.DB
	err  error
)

func Init() error {
	dbConfiguration := &dbPostgreSQL{
		db: db{
			Host: os.Getenv(constant.DB_HOST),
			User: os.Getenv(constant.DB_USER),
			Pass: os.Getenv(constant.DB_PASS),
			Port: os.Getenv(constant.DB_PORT),
			Name: os.Getenv(constant.DB_NAME),
		},
		SslMode: os.Getenv(constant.DB_SSLMODE),
		Tz:      os.Getenv(constant.DB_TZ),
	}

	conn, err = dbConfiguration.Init()
	if err != nil {
		panic(fmt.Sprintf("failed to connect to database %s", dbConfiguration.db.Name))
	}

	logrus.Info("successfully initialized database")
	return nil
}

func GetConnection() (*gorm.DB, error) {
	if conn == nil {
		return nil, fmt.Errorf("connection is undefined")
	}
	return conn, nil
}
