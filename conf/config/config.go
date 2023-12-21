package config

import (
	"os"
	"reflect"
	"strings"

	"gopkg.in/yaml.v2"
)

type Config struct {
	BaseURL              string `yaml:"base_url" env:"BASE_URL"`
	OpenAIKey            string `yaml:"openai_key" env:"OPENAI_API_KEY"`
	LineBotChannelSecret string `yaml:"line_bot_channel_secret" env:"LINE_BOT_CHANNEL_SECRET"`
	LineBotChannelToken  string `yaml:"line_bot_channel_token" env:"LINE_BOT_CHANNEL_TOKEN"`
	TGBotToken           string `yaml:"tg_bot_token" env:"TG_BOT_TOKEN"`
}

var Conf Config

func updateConfigFromEnv(cfg interface{}) {
	v := reflect.ValueOf(cfg).Elem()
	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		fieldType := t.Field(i)

		envVar := fieldType.Tag.Get("env")
		if envVar == "" {
			envVar = strings.ToUpper(fieldType.Name)
		}

		if value, exists := os.LookupEnv(envVar); exists {
			if field.CanSet() {
				field.SetString(value)
			}
		}
	}
}

func (c *Config) LoadConfig(configPath string) error {
	configFile, err := os.ReadFile(configPath)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(configFile, c)
	if err != nil {
		return err
	}

	updateConfigFromEnv(c)

	return nil
}
