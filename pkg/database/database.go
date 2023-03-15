package database

import (
	"fmt"
	"time"

	"github.com/falahlaz/boilerplate-golang/pkg/config"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"gorm.io/plugin/dbresolver"
)

var connections map[string]*gorm.DB = make(map[string]*gorm.DB)
var DBResolver *gorm.DB

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
					MaxIdleConn: config.Config.Server.MaxIdleConn,
					MaxOpenConn: config.Config.Server.MaxOpenConn,
					MaxLifetime: config.Config.Server.MaxLifetime,
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
					MaxIdleConn: config.Config.Server.MaxIdleConn,
					MaxOpenConn: config.Config.Server.MaxOpenConn,
					MaxLifetime: config.Config.Server.MaxLifetime,
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
		connections[k] = db
		logrus.Info("successfully connected to database", k)
	}

	logrus.Info("successfully initialized database")
	return nil
}

func GetConnection(key string) (*gorm.DB, error) {
	if connections[key] == nil {
		return nil, fmt.Errorf("connection is undefined")
	}
	return connections[key], nil
}

func GetDB() (*gorm.DB, error) {
	if DBResolver == nil {
		return nil, fmt.Errorf("connection is undefined")
	}
	return DBResolver, nil
}

func Resolver() {
	dbConn := []*gorm.DB{}
	dbConfig := []gorm.Dialector{}

	for _, v := range config.Config.Multiconns {
		db, err := GetConnection(v)
		if err != nil {
			panic("Failed to make database resolver")
		}
		dbConn = append(dbConn, db)
		dbConfig = append(dbConfig, db.Dialector)
	}

	dbConn[0].Use(
		dbresolver.Register(dbresolver.Config{
			Replicas: dbConfig,
			Policy:   dbresolver.RandomPolicy{},
		}).
			SetConnMaxLifetime(time.Duration(config.Config.Server.MaxLifetime)).
			SetMaxIdleConns(config.Config.Server.MaxIdleConn).
			SetMaxOpenConns(config.Config.Server.MaxOpenConn))

	DBResolver = dbConn[0]
	logrus.Info("success create resolver")
}
