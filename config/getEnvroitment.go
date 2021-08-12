package config

import "os"

func getEnv(key, default_val string) string {
	value, exist := os.LookupEnv(key)
	if !exist {
		return default_val
	}
	return value
}
