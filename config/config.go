package config

import (
	"fmt"
	"time"

	"github.com/gosimple/conf"
)

type Config struct {
	Username             string
	Password             string
	SentryFile           string
	EmailAddress         string
	EmailPassword        string
	EmailRetryCount      int
	EmailFetchRetryDelay time.Duration
}

func LoadConfig(filename string) (*Config, error) {
	botConfig := &Config{}
	config, err := conf.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("could not read config file '%s': %v", filename, err)
	}
	botConfig.Username, err = config.String("default", "username")
	if err != nil {
		return nil, fmt.Errorf("could not read username: %v", err)
	}
	botConfig.Password, err = config.String("default", "password")
	if err != nil {
		return nil, fmt.Errorf("could not read password: %v", err)
	}
	botConfig.SentryFile, err = config.String("default", "sentry-file")
	if err != nil {
		return nil, fmt.Errorf("could not read sentry-file: %v", err)
	}
	botConfig.EmailAddress, err = config.String("default", "email-address")
	if err != nil {
		return nil, fmt.Errorf("could not read email-address: %v", err)
	}
	botConfig.EmailPassword, err = config.String("default", "email-password")
	if err != nil {
		return nil, fmt.Errorf("could not read email-password: %v", err)
	}
	botConfig.EmailRetryCount, err = config.Int("default", "email-retry-count")
	if err != nil {
		return nil, fmt.Errorf("could not read email-retry-count: %v", err)
	}
	delayInSeconds, err := config.Int("default", "email-fetch-retry-delay")
	if err != nil {
		return nil, fmt.Errorf("could not read email-fetch-retry-delay: %v", err)
	}
	botConfig.EmailFetchRetryDelay = time.Duration(delayInSeconds) * time.Second
	return botConfig, nil
}
