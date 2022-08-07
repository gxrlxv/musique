package config

import (
	"github.com/gxrlxv/musique/gateway/pkg/logging"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/sirupsen/logrus"
	"sync"
)

type Config struct {
	Rest struct {
		Addr string `yaml:"addr"`
	} `yaml:"rest"`
	Auth struct {
		Addr string `yaml:"addr"`
	} `yaml:"auth"`
	Artist struct {
		Addr string `yaml:"addr"`
	} `yaml:"artist"`
}

var instance *Config
var once sync.Once

func GetConfig() *Config {
	once.Do(func() {
		logger := logging.GetLogger()
		logger.Info("read application configuration")
		instance = &Config{}
		if err := cleanenv.ReadConfig("config.yml", instance); err != nil {
			help, _ := cleanenv.GetDescription(instance, nil)
			logrus.Info(help)
			logrus.Fatal(err)
		}
	})
	return instance
}
