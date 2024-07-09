package config

import (
	"fmt"
)

const (
	DefaultPassword = "123456"
	DefaultPort     = "7540"
)

type Config struct {
	AppPass   string
	SecretKey string // секретный ключ шифрования
	AppPort   string
}

// NewConfig конструктор объекта конфигурации приложения
func NewConfig(password string, encKey string, appPort string) (*Config, error) {
	// Если каких то переменных нет, то используем дефолтные значения
	if password == "" {
		password = DefaultPassword
	}
	if password == "" || encKey == "" {
		return nil, fmt.Errorf("invalid config")
	}

	if appPort == "" {
		appPort = DefaultPort
	}
	return &Config{AppPass: password, SecretKey: encKey, AppPort: appPort}, nil
}
