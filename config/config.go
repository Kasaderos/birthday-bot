package config

import (
	"github.com/spf13/viper"
)

type ConfSt struct {
	Debug              bool   `mapstructure:"DEBUG"`
	LogLevel           string `mapstructure:"LOG_LEVEL"`
	PgDsn              string `mapstructure:"PG_DSN"`
	HttpListen         string `mapstructure:"HTTP_LISTEN"`
	HttpCors           bool   `mapstructure:"HTTP_CORS"`
	NotifyTimeInterval string `mapstructure:"NOTIFY_TIME_INTERVAL"`
	CertPath           string `mapstructure:"CERT_PATH"`
	CertPsw            string `mapstructure:"CERT_PSW"`
	BotToken           string `mapstructure:"BOT_TOKEN"`
}

func Load() *ConfSt {
	result := &ConfSt{}

	viper.SetDefault("DEBUG", "false")
	viper.SetDefault("LOG_LEVEL", "info")
	viper.SetDefault("HTTP_LISTEN", ":8888")
	viper.SetDefault("CERT_PATH", "")
	viper.SetDefault("CERT_PSW", "")

	viper.SetDefault("NOTIFY_TIME_INTERVAL", "10:00-18:00")

	viper.AutomaticEnv()
	result.BotToken = viper.GetString("BOT_TOKEN")
	result.PgDsn = viper.GetString("PG_DSN")
	result.HttpListen = viper.GetString("HTTP_LISTEN")
	result.LogLevel = viper.GetString("LOG_LEVEL")

	return result
}
