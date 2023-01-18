package database

import (
	"fmt"

	"github.com/falahlaz/boilerplate-golang/pkg/config"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var dbConnections map[string]*gorm.DB

func Init() error {
	dbConfigurations := make(map[string]Db)

	for _, v := range config.Config.Databases {
		switch v.Connection {
		case "pgsql":
			dbConfigurations[v.Alias] = &dbPostgreSQL{
				db: db{
					Host:        v.Host,
					User:        v.User,
					Pass:        v.Pass,
					Port:        v.Port,
					Name:        v.Name,
					MaxIdleConn: v.MaxIdleConn,
					MaxOpenConn: v.MaxOpenConn,
					MaxLifetime: v.MaxLifetime,
					AutoMigrate: v.Migration,
					Seeder:      v.Seeder,
				},
				SslMode: v.Ssl,
				Tz:      v.Tz,
			}
		case "mysql":
			dbConfigurations[v.Alias] = &dbMySQL{
				db: db{
					Host:        v.Host,
					User:        v.User,
					Pass:        v.Pass,
					Port:        v.Port,
					Name:        v.Name,
					MaxIdleConn: v.MaxIdleConn,
					MaxOpenConn: v.MaxOpenConn,
					MaxLifetime: v.MaxLifetime,
					AutoMigrate: v.Migration,
					Seeder:      v.Seeder,
				},
				Charset:   v.Charset,
				ParseTime: v.ParseTime,
				Loc:       v.Location,
			}
		}
	}

	for k, v := range dbConfigurations {
		db, err := v.Init()
		if err != nil {
			panic(fmt.Sprintf("failed to connect to %s database", k))
		}
		dbConnections[k] = db
	}

	logrus.Info("successfully initialized database")
	return nil
}

func GetConnection(key string) (*gorm.DB, error) {
	if dbConnections[key] == nil {
		return nil, fmt.Errorf("connection is undefined")
	}
	return dbConnections[key], nil
}
