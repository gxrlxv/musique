package config

import (
	"github.com/gxrlxv/musique/auth_service/pkg/logging"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/sirupsen/logrus"
	"time"

	"sync"
)

type Config struct {
	IsDebug *bool `yaml:"is_debug" env-required:"true"`
	Server  struct {
		Http struct {
			Addr string `yaml:"addr"`
		} `yaml:"http"`
		Grpc struct {
			Addr string `yaml:"addr"`
		} `yaml:"grpc"`
	} `yaml:"server"`
	Storage StorageConfig `yaml:"storage"`
	JWT     struct {
		AccessTokenTTL  time.Duration `yaml:"access_token_ttl"`
		RefreshTokenTTL time.Duration `yaml:"refresh_token_ttl"`
		SigningKey      string        `yaml:"signing_key"`
	} `yaml:"jwt"`
	Hasher struct {
		Salt string `yaml:"salt"`
	} `yaml:"hasher"`
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
