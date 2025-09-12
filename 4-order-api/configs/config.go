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
	DB   DBConfig
	Auth AuthConfig
}

type DBConfig struct {
	Dsn string
}

type AuthConfig struct {
	Secret string
}

func LoadConfig(configName string) *Config {
	err := godotenv.Load(filepath.Join(RootDir(), configName))
	if err != nil {
		log.Printf("Error loading %s file, using default config\n", configName)
	}

	return &Config{
		DB: DBConfig{
			Dsn: os.Getenv("DSN"),
		},
		Auth: AuthConfig{
			Secret: os.Getenv("TOKEN"),
		},
	}
}

func RootDir() string {
	_, b, _, _ := runtime.Caller(0)
	return filepath.Join(filepath.Dir(b), "..")
}
