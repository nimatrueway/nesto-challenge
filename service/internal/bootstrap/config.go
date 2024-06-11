package bootstrap

import (
	"time"

	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

type Config struct {
	Database struct {
		URL             string        `mapstructure:"url"`
		MaxConns        int           `mapstructure:"max_connections"`
		MaxConnIdleTime time.Duration `mapstructure:"max_connection_idle_time"`
	} `mapstructure:"database"`
	Log struct {
		Level string `mapstructure:"level"`
	} `mapstructure:"log"`
	Server struct {
		Host                     string        `mapstructure:"bind"`
		RequestReadHeaderTimeout time.Duration `mapstructure:"request_read_header_timeout"`
		CorsAllowedOrigins       []string      `mapstructure:"cors_allowed_origins"`
	} `mapstructure:"server"`
}

func LoadConfig() (*Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, errors.Wrapf(err, "failed to configuration file")
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, errors.Wrapf(err, "failed to unmarshal configuration")
	}

	return &config, nil
}
