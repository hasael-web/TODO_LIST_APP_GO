package utils

import (
	"os"
	"regexp"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

// TODO: Envirotment logic

const projectDirName = "todo-project"

type Env interface {
	GetString(name string) string
}

type env struct {
	Env
}

type EnvGetter struct{}

func NewEnv() *env {
	return &env{Env: &EnvGetter{}}
}

func (e *env) Load(env string) {
	re := regexp.MustCompile(`^(.*` + projectDirName + `)`)
	cwd, _ := os.Getwd()
	rooPath := re.Find([]byte(cwd))

	var envFile string

	switch env {
	case "LOCAL":
		envFile = "local"
	case "DEV":
		envFile = "development"
	case "STAGING":
		envFile = "staging"
	case "PROD":
		envFile = "production"
	default:
		envFile = ""
	}

	err := godotenv.Load(string(rooPath) + `.env.` + envFile)

	if err != nil {
		logrus.WithFields(logrus.Fields{
			"cause": err,
			"cwd":   cwd,
		}).Fatal("Load .env file error")

		os.Exit(-1)
	}

}
func (r *EnvGetter) GetString(name string) string {
	return os.Getenv(name)
}

func (e *env) GetString(name string) string {
	if nil == e.Env {
		return ""
	}
	return e.Env.GetString(name)
}

func (e *env) GetInt(name string) int {
	s := e.GetString(name)
	i, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return i
}
