package config

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/qiangxue/go-env"
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

const (
	envVariablesPrefix = "APP_"

	defaultBindAddr = ":8080"
)

// Config содержит конфигурационные данные для работы сервиса.
type Config struct {
	BindAddr string `yaml:"bind_addr" env:"BIND_ADDR"` // адрес, на котором запущен сервер (по умолчанию, ":8080")
	DSN      string `yaml:"dsn" env:"DSN,secret"`      // строка подключения к базе данных
}

// Validate проверяет, достаточно ли информации в конфиге для запуска сервиса.
func (c Config) Validate() error {
	return validation.ValidateStruct(&c, validation.Field(&c.DSN, validation.Required))
}

// Load подгружает конфигурационные данные в конфиг из файла по указанному пути и из переменных окружения.
func Load(yamlConfigPath string) (*Config, error) {
	c := Config{
		BindAddr: defaultBindAddr,
	}

	bytes, err := ioutil.ReadFile(yamlConfigPath)
	if err != nil {
		return nil, err
	}
	if err = yaml.Unmarshal(bytes, &c); err != nil {
		return nil, err
	}

	// загрузка переменных окружения с указанным префиксом
	if err = env.New(envVariablesPrefix, nil).Load(&c); err != nil {
		return nil, err
	}

	if err = c.Validate(); err != nil {
		return nil, err
	}

	return &c, nil
}