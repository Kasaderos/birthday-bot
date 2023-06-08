package config

import (
	"github.com/spf13/viper"
)

type ConfSt struct {
	Debug          bool   `mapstructure:"DEBUG"`
	LogLevel       string `mapstructure:"LOG_LEVEL"`
	PgDsn          string `mapstructure:"PG_DSN"`
	HttpListen     string `mapstructure:"HTTP_LISTEN"`
	HttpCors       bool   `mapstructure:"HTTP_CORS"`
	NotifyInterval string `mapstructure:"NOTIFY_INTERVAL"`
	CertPath       string `mapstructure:"CERT_PATH"`
	CertPsw        string `mapstructure:"CERT_PSW"`
	BotToken       string `mapstructure:"BOT_TOKEN"`
}

func Load() *ConfSt {
	result := &ConfSt{}

	viper.SetDefault("DEBUG", "false")
	viper.SetDefault("LOG_LEVEL", "info")
	viper.SetDefault("HTTP_LISTEN", ":80")
	viper.SetDefault("CERT_PATH", "")
	viper.SetDefault("CERT_PSW", "")

	viper.SetDefault("NOTIFY_INTERVAL", "10:00-18:00")

	viper.SetConfigFile("conf.yml")
	_ = viper.ReadInConfig()

	viper.AutomaticEnv()
	_ = viper.Unmarshal(&result)

	return result
}
