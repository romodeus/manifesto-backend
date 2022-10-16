package config

import (
	"os"
	"sync"
)

type AppConfig struct {
	REDIS_HOST     string
	REDIS_PASSWORD string
	SERVER_PORT    string
	JWT_SECRET     string
	BASE_URL       string
}

var mutex = sync.Mutex{}

func GetConfig() *AppConfig {
	mutex.Lock()
	defer mutex.Unlock()
	return loadConfig()
}

func loadConfig() *AppConfig {

	return &AppConfig{
		SERVER_PORT:    os.Getenv("SERVER_PORT"),
		BASE_URL:       os.Getenv("BASE_URL") + ":" + os.Getenv("SERVER_PORT"),
		REDIS_HOST:     os.Getenv("REDIS_HOST") + ":" + os.Getenv("REDIS_PORT"),
		REDIS_PASSWORD: os.Getenv("REDIS_PASSWORD"),
	}
}
