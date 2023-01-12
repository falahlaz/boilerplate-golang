package env

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

type Env interface {
	GetString(key string) string
	GetBool(key string) bool
	GetInt(key string) int
	GetFloat(key string) float64
}

type env struct{}

func NewEnv() *env {
	return &env{}
}

func (env *env) Load() {
	if err := godotenv.Load(".env"); err != nil {
		logrus.WithFields(logrus.Fields{
			"cause": err,
		}).Fatal("Load .env file error")
	}
	logrus.Info("This program is running in " + env.GetString("APP_ENV") + " environment")
}

// GetBool implements Env
func (*env) GetBool(key string) bool {
	s := os.Getenv(key)
	i, err := strconv.ParseBool(s)
	if err != nil {
		return false
	}
	return i
}

// GetFloat implements Env
func (*env) GetFloat(key string) float64 {
	s := os.Getenv(key)
	i, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0
	}
	return i
}

// GetInt implements Env
func (*env) GetInt(key string) int {
	s := os.Getenv(key)
	i, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return i
}

// GetString implements Env
func (*env) GetString(key string) string {
	return os.Getenv(key)
}
