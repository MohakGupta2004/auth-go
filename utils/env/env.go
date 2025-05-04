package env

import (
	"os"
	"strconv"
)

func GetString(required_pkg string, fallback string) string {
	pkg, ok := os.LookupEnv(required_pkg)
	if !ok {
		return fallback
	}

	return pkg
}

func GetInt(required_pkg string, fallback int) int {
	pkg, ok := os.LookupEnv(required_pkg)
	if !ok {
		return fallback
	}
	valAsInt, err := strconv.Atoi(pkg)
	if err != nil {
		return fallback
	}
	return valAsInt
}
