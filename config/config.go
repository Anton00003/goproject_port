package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Host   string
	Level  string
	NumIn  int
	NumOut int
}

func New() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	return &Config{
		Host:   getEnv("HOST"),
		Level:  getEnv("LEVEL"),
		NumIn:  getEnvInt("NUMIN"),
		NumOut: getEnvInt("NUMOUT"),
	}, nil
}

func getEnv(key string) string {
	val, exists := os.LookupEnv(key)
	if exists {
		return val
	}
	return ""
}

func getEnvInt(key string) int {
	valS := getEnv(key)
	valI, err := strconv.Atoi(valS)
	if err != nil {
		return 0
	}
	return valI
}
