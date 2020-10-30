package util

import "os"

// GetFirstEnv get the value of the first defined environment variable.
func GetFirstEnv(names ...string) string {
	for _, name := range names {
		value := os.Getenv(name)
		if value != "" {
			return value
		}
	}

	return ""
}
