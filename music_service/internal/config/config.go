package config

import (
	"github.com/gxrlxv/musique/music_service/pkg/logging"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/sirupsen/logrus"
	"sync"
)

type Config struct {
	Server struct {
		Http struct {
			Addr string `yaml:"addr"`
		} `yaml:"http"`
		Grpc struct {
			Addr string `yaml:"addr"`
		} `yaml:"grpc"`
	} `yaml:"server"`
	Storage StorageConfig `yaml:"storage"`
}

type StorageConfig struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	Database string `json:"database"`
	Username string `json:"username"`
	Password string `json:"password"`
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
