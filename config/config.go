package config

import (
	"os"
	"sync"
)

type AppConfig struct {
	DB_NAME     string
	DB_HOST     string
	DB_PORT     int
	DB_USER     string
	DB_PASS     string
	SERVER_PORT string
	JWT_SECRET  string
	BASE_URL    string
}

var mutex = sync.Mutex{}

func GetConfig() *AppConfig {
	mutex.Lock()
	defer mutex.Unlock()
	return loadConfig()
}

func loadConfig() *AppConfig {
	return &AppConfig{
		SERVER_PORT: os.Getenv("SERVER_PORT"),
		BASE_URL:    os.Getenv("BASE_URL"),
	}
}
