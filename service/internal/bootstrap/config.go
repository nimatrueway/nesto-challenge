package bootstrap

import (
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	Database struct {
		Dns             string        `mapstructure:"dsn"`
		MaxConns        int           `mapstructure:"max_connections"`
		MaxConnIdleTime time.Duration `mapstructure:"max_connection_idle_time"`
	} `mapstructure:"database"`
	Log struct {
		Level string `mapstructure:"level"`
	} `mapstructure:"log"`
	Server struct {
		Host               string   `mapstructure:"bind"`
		CorsAllowedOrigins []string `mapstructure:"cors_allowed_origins"`
	} `mapstructure:"server"`
}

func LoadConfig() (*Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	return &config, nil
}
