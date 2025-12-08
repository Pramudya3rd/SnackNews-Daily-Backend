package config

import (
	"github.com/spf13/viper"
)

type DatabaseConfig struct {
    Host     string `mapstructure:"host"`
    Port     int    `mapstructure:"port"`
    User     string `mapstructure:"user"`
    Password string `mapstructure:"password"`
    Name     string `mapstructure:"name"`
}

type ServerConfig struct {
    Port         int    `mapstructure:"port"`
    ReadTimeout  string `mapstructure:"read_timeout"`
    WriteTimeout string `mapstructure:"write_timeout"`
    JWTSecret    string `mapstructure:"jwt_secret"`
}

type Config struct {
    Database DatabaseConfig `mapstructure:"database"`
    Server   ServerConfig   `mapstructure:"server"`
}

// LoadConfig reads configuration from configs/config.yaml and returns Config
func LoadConfig() (Config, error) {
    var cfg Config

    viper.SetConfigName("config")
    viper.SetConfigType("yaml")
    viper.AddConfigPath("./configs")

    if err := viper.ReadInConfig(); err != nil {
        return cfg, err
    }

    if err := viper.Unmarshal(&cfg); err != nil {
        return cfg, err
    }

    return cfg, nil
}
