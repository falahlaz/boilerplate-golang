package database

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type db struct {
	Host string
	User string
	Pass string
	Port string
	Name string
}

type dbPostgreSQL struct {
	db
	SslMode string
	Tz      string
}

type dbMySQL struct {
	db
	Charset   string
	ParseTime string
	Loc       string
}

func (d *dbPostgreSQL) Init() (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s", d.Host, d.User, d.Pass, d.Name, d.Port, d.SslMode, d.Tz)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(setLogMode()),
	})
	if err != nil {
		return nil, err
	}
	return db, nil
}

func (d *dbMySQL) Init() (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=%s&loc=%s", d.User, d.Pass, d.Host, d.Port, d.Name, d.Charset, d.ParseTime, d.Loc)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(setLogMode()),
	})
	if err != nil {
		return nil, err
	}
	return db, nil
}

func setLogMode() logger.LogLevel {
	logLevel := logger.Info
	env := os.Getenv("APP_ENV")
	if env == "production" {
		logLevel = logger.Error
	}
	return logLevel
}
