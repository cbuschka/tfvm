package util

import "os"

// IsAlphaFeatureEnabled tests if alpha features are enabled via env var TFVM_ALPHA_ENABLED.
func IsAlphaFeatureEnabled() bool {
	value := os.Getenv("TFVM_ALPHA_ENABLED")
	return value == "1" || value == "true" || value == "yes" || value == "y" || value == "t"
}
