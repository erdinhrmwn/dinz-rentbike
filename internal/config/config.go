package config

import (
	"github.com/spf13/viper"

	"dinz-rentbike/pkg/logger"
)

type Config struct {
	App      AppConfig
	Database DatabaseConfig
	Redis    RedisConfig
	Jwt      JwtConfig
	Xendit   XenditConfig
	Mailjet  MailjetConfig
}

type AppConfig struct {
	Host string
	Port int
	Name string
}

type DatabaseConfig struct {
	Host   string
	Port   int
	User   string
	Pass   string
	Name   string
	Schema string
}

type RedisConfig struct {
	Host     string
	Port     int
	Password string
	DB       int
}

type JwtConfig struct {
	Secret string
}

type XenditConfig struct {
	BaseURL      string
	PublicKey    string
	SecretKey    string
	WebhookToken string
}

type MailjetConfig struct {
	BaseURL   string
	ApiKey    string
	SecretKey string
}

func Load() (*Config, error) {
	v := viper.New()

	v.SetConfigFile(".env")
	v.SetConfigType("env")
	v.AutomaticEnv()

	if err := v.ReadInConfig(); err != nil {
		logger.Log.Warn().Msg("there is no .env file, using default values")
	}

	cfg := &Config{
		App: AppConfig{
			Host: v.GetString("APP_HOST"),
			Port: v.GetInt("APP_PORT"),
			Name: v.GetString("APP_NAME"),
		},
		Database: DatabaseConfig{
			Host:   v.GetString("DB_HOST"),
			Port:   v.GetInt("DB_PORT"),
			User:   v.GetString("DB_USER"),
			Pass:   v.GetString("DB_PASS"),
			Name:   v.GetString("DB_NAME"),
			Schema: v.GetString("DB_SCHEMA"),
		},
		Redis: RedisConfig{
			Host:     v.GetString("REDIS_HOST"),
			Port:     v.GetInt("REDIS_PORT"),
			Password: v.GetString("REDIS_PASSWORD"),
			DB:       v.GetInt("REDIS_DB"),
		},
		Jwt: JwtConfig{
			Secret: v.GetString("JWT_SECRET"),
		},
		Xendit: XenditConfig{
			BaseURL:      v.GetString("XENDIT_BASE_URL"),
			PublicKey:    v.GetString("XENDIT_PUBLIC_KEY"),
			SecretKey:    v.GetString("XENDIT_SECRET_KEY"),
			WebhookToken: v.GetString("XENDIT_WEBHOOK_TOKEN"),
		},
		Mailjet: MailjetConfig{
			BaseURL:   v.GetString("MAILJET_BASE_URL"),
			ApiKey:    v.GetString("MAILJET_API_KEY"),
			SecretKey: v.GetString("MAILJET_SECRET_KEY"),
		},
	}

	return cfg, nil
}
