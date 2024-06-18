package configs

import (
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

var c *config

type config struct {
	Port    string
	WorkDir string
}

func Init() {
	err := godotenv.Load()
	if err != nil {
		logrus.Fatal("Error loading .env file")
	}

	c = &config{
		Port:    getString("PORT", "8080"),
		WorkDir: getWorkDir("WORKDIR"),
	}
}

func Get() *config {
	return c
}

func getString(key string, defaultValue ...string) string {
	env := os.Getenv(key)
	if env == "" && len(defaultValue) > 0 {
		env = defaultValue[0]
	}

	return env
}

func getWorkDir(key string) string {
	workdir := getString(key)
	if workdir == "" {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			workdir = "."
		}

		workdir = filepath.Join(homeDir, "sheds")
	}

	return workdir
}
