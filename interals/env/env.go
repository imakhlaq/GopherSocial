package env

import (
	"os"
)

func GetEnv(key, fallback string) string {
	val, ok := os.LookupEnv(key)

	//checking if value exits
	if !ok {
		return fallback
	}
	return val
}
