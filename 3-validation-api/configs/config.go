// Package configs for load configuration from os env
package configs

import (
	"log"
	"os"
	"path/filepath"
	"runtime"

	"github.com/joho/godotenv"
)

type Config struct {
	Email         string
	Password      string
	Address       string
	Port          string
	StoreFilename string
}

func LoadConfig(configName string) *Config {
	err := godotenv.Load(filepath.Join(RootDir(), configName))
	if err != nil {
		log.Printf("Error loading %s file, using default config\n", configName)
	}

	return &Config{
		Email:         os.Getenv("EMAIL"),
		Password:      os.Getenv("PASSWORD"),
		Address:       os.Getenv("SMTP_ADDRESS"),
		Port:          os.Getenv("SMTP_PORT"),
		StoreFilename: os.Getenv("STORE_FILENAME"),
	}
}

func RootDir() string {
	_, b, _, _ := runtime.Caller(0)
	return filepath.Join(filepath.Dir(b), "..")
}
