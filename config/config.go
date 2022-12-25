package config

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/oktaykcr/simple-go-rest-api/logger"
	"gopkg.in/yaml.v3"
	"os"
)

var Conf Config

type Config struct {
	Server struct {
		Port string `yaml:"port", envconfig:"SERVER_PORT"`
		Host string `yaml:"host", envconfig:"SERVER_HOST"`
	} `yaml:"server"`
	Database struct {
		Username string `yaml:"user", envconfig:"DB_USERNAME"`
		Password string `yaml:"pass", envconfig:"DB_PASSWORD"`
	} `yaml:"database"`
}

func Initialize() {
	readFile(&Conf)
	readEnv(&Conf)
}

func readFile(cfg *Config) {
	configFile, err := os.Open("config.yml")
	if err != nil {
		logger.Error(err.Error())
	}
	defer configFile.Close()

	decoder := yaml.NewDecoder(configFile)
	err = decoder.Decode(cfg)
	if err != nil {
		logger.Error(err.Error())
	}
}

func readEnv(cfg *Config) {
	err := envconfig.Process("", cfg)
	if err != nil {
		logger.Error(err.Error())
	}
}
