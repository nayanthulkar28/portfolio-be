package config

import (
	"github.com/spf13/viper"
)

type (
	Config struct {
		App    App    `mapstructure:",squash"`
		Http   Http   `mapstructure:",squash"`
		WebAPI WebAPI `mapstructure:",squash"`
	}

	App struct {
		Name    string `mapstructure:"APP_NAME"`
		Version string `mapstructure:"APP_VERSION"`
		APIKey  string `mapstructure:"API_KEY"`
	}

	Http struct {
		Port string `mapstructure:"HTTP_PORT"`
	}

	WebAPI struct {
		Email Email `mapstructure:",squash"`
	}

	Email struct {
		Name              string `mapstructure:"EMAIL_NAME"`
		FromEmailAddress  string `mapstructure:"FROM_EMAIL_ADDRESS"`
		FromEmailPassword string `mapstructure:"FROM_EMAIL_PASSWORD"`
		ToEmailAddress    string `mapstructure:"TO_EMAIL_ADDRESS"`
		SmtpAuthAddress   string `mapstructure:"SMTP_AUTH_ADDRESS"`
		SmtpServerAddress string `mapstructure:"SMTP_SERVER_ADDRESS"`
	}
)

func NewConfig(path string) (*Config, error) {
	var cfg *Config
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}
	err = viper.Unmarshal(&cfg)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}
