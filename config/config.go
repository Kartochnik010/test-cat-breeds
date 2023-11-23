package config

import (
	"errors"
	"os"
	"strconv"

	_ "github.com/joho/godotenv/autoload"
)

type Config struct {
	BaseURL         string
	RequestInterval int
	OutFileName     string
}

func InitConfig() (*Config, error) {
	baseURL := os.Getenv("URL_BASE")
	if baseURL == "" {
		return &Config{}, errors.New("var URL_BASE is unset")
	}

	requestIntervalString := os.Getenv("var REQUEST_INTERVAL")
	rI, err := strconv.Atoi(requestIntervalString)
	if err != nil {
		return &Config{}, err
	}
	outFileName := os.Getenv("OUT_FILE_NAME")
	if baseURL == "" {
		return &Config{}, errors.New("var OUT_FILE_NAME is unset")
	}

	return &Config{
		BaseURL:         baseURL,
		RequestInterval: rI,
		OutFileName:     outFileName,
	}, nil
}
