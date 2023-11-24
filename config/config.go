package config

import (
	"log"
	"purchasing-management/internal/common/database"
	"purchasing-management/internal/common/http"

	validation "github.com/go-ozzo/ozzo-validation"
)

const defaultConfigPath = "./config/config.yaml"

type Config struct {
	App               App             `mapstructure:"app"`
	HTTP              http.Config     `mapstructure:"http"`
	PrimaryDataSource database.Config `mapstructure:"primary_data_source"`
}

type App struct {
	Environment string `mapstructure:"env"`
}

func (a App) verify() {
	err := validation.ValidateStruct(&a, validation.Field(&a.Environment, validation.Required))
	if err != nil {
		log.Fatalf("invalid app config %v", err)
	}
}

func LoadConfig() Config {
	configPath, err := getConfigPath()
	if err != nil {
		log.Panicf("Failed to get config path %v", err)
	}
	if configPath == "" {
		configPath = defaultConfigPath
	}

	var conf Config
	if err = loadConfig(configPath, &conf); err != nil {
		log.Panicf("Failed to load config %v", err)
	}
	verifyConfig(conf)
	printConfig(conf)

	return conf
}
