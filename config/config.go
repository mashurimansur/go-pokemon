package config

import (
	"os"

	"github.com/joho/godotenv"
)

// Config : interface for config env
type Config interface {
	Get(key string) string
}

type configImpl struct {
}

// Get return key format string
func (config *configImpl) Get(key string) string {
	return os.Getenv(key)
}

// New : new env
func New(filenames ...string) Config {
	err := godotenv.Load(filenames...)
	if err != nil {
		panic(err)
	}
	return &configImpl{}
}
