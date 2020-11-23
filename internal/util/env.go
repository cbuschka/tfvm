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

// IsEnvVarTrue tests if a env var is set to 1, true, yes, y or t.
func IsEnvVarTrue(name string) bool {
	value := os.Getenv(name)
	return value == "1" || value == "true" || value == "yes" || value == "y" || value == "t"
}
