package config

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

type config struct {
	AppHost string `mapstructure:"APP_HOST"`
	AppPort string `mapstructure:"App_PORT"`
}

const (
	defaultEnvFile = ".env"
)

func LoadEnv(path string) (cfg config, err error) {
	viper.SetConfigFile(getConfigFullPath(path))
	viper.AutomaticEnv()

	if err = viper.ReadInConfig(); err != nil {
		return
	}

	err = viper.Unmarshal(&cfg)
	return
}

func getConfigFullPath(path string) string {
	envFile := defaultEnvFile
	if os.Getenv("APP_ENV") != "" {
		envFile = fmt.Sprintf("%s.%s", envFile, os.Getenv("APP_ENV"))
	}

	return filepath.Join(path, envFile)
}

func SetAppEnv(appEnv string) {
	err := os.Setenv("APP_ENV", appEnv)
	if err != nil {
		log.Fatal("unable to set app environment:", err)
	}
}
