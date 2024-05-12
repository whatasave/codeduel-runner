package utils

import (
	"fmt"
	"os"
)

func GetEnv(key string) (string, error) {
	value, exists := os.LookupEnv(key)
	if !exists {
		return "", fmt.Errorf("%s not defined in .env file", key)
	}

	return value, nil
}
