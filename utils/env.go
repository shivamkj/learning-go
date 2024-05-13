package utils

import (
	"github.com/joho/godotenv"
)

func LoadEnvFile(source string) error {
	if source == "" {
		source = ".env"
	}
	err := godotenv.Load(source)
	return err
}
