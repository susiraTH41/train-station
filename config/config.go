package config

import (
	"strings"
	"sync"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

type (
	Config struct {
		Server   *Server   `mapstructure:"server" validate:"required"`
		Database *Database `mapstructure:"database" validate:"required"`
		LinkAPI *LinkAPI `mapstructure:"linkapi" validate:"required"`
		Auth *Auth `mapstructure:"auth" validate:"required"`
	}

	LinkAPI struct{
		GetStationUrl string `mapstructure:"trainstation" validate:"required"`
	}

	Auth struct{
		Secret string `mapstructure:"secret" validate:"required"`
	}

	Server struct {
		Port           int      `mapstructural:"port" validate:"required"`
		AllowOrigins   []string `mapstructural:"allowOrigins" validate:"required"`
		BodyLimit      string   `mapstructural:"bodyLimit" validate:"required"`
		TimeOut        time.Duration      `mapstructural:"timeout" validate:"required"`
	}


	Database struct {
		Host     string `mapstructure:"host" validate:"required"`
		Port     int    `mapstructure:"port" validate:"required"`
		User     string `mapstructure:"user" validate:"required"`
		Password string `mapstructure:"password" validate:"required"`
		DBName   string `mapstructure:"dbname" validate:"required"`
		SSLMode  string `mapstructure:"sslmode" validate:"required"`
		Schema   string `mapstructure:"schema" validate:"required"`
	
	}
)

var (
	once           sync.Once
	configInstance *Config
)

func ConfigGetting() *Config {
	once.Do(func() {
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
		viper.AddConfigPath(".")
		viper.AutomaticEnv()
		viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

		if err := viper.ReadInConfig(); err != nil {
			panic(err)
		}

		if err := viper.Unmarshal(&configInstance); err != nil {
			panic(err)
		}

		validate := validator.New()

		if err := validate.Struct(configInstance); err != nil {
			panic(err)
		}
	})

	return configInstance
}