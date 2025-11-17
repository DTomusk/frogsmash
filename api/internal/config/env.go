package config

import (
	"fmt"
	"os"
	"strconv"
)

func getEnv(key string) string {
	val := os.Getenv(key)
	if val == "" {
		panic(fmt.Sprintf("missing required env var: %s", key))
	}
	return val
}

func getInt(key string) int {
	v := getEnv(key)
	i, err := strconv.Atoi(v)
	if err != nil {
		panic(fmt.Sprintf("env var %s must be an integer", key))
	}
	return i
}

func getFloat(key string) float64 {
	v := getEnv(key)
	f, err := strconv.ParseFloat(v, 64)
	if err != nil {
		panic(fmt.Sprintf("env var %s must be a float", key))
	}
	return f
}

func getInt64(key string) int64 {
	v := getEnv(key)
	i, err := strconv.ParseInt(v, 10, 64)
	if err != nil {
		panic(fmt.Sprintf("env var %s must be int64", key))
	}
	return i
}
