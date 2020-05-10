package util

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestGetFirstEnvReturnsFirstIfSet(t *testing.T) {

	os.Setenv("A", "a")
	os.Setenv("B", "b")

	result := GetFirstEnv("A", "B")

	assert.Equal(t, "a", result)
}

func TestGetFirstEnvReturnsSecondIfFirstNotSet(t *testing.T) {

	os.Unsetenv("A")
	os.Setenv("B", "b")

	result := GetFirstEnv("A", "B")

	assert.Equal(t, "b", result)
}

func TestGetFirstEnvReturnsEmptyIfNoneSet(t *testing.T) {

	os.Unsetenv("A")
	os.Unsetenv("B")

	result := GetFirstEnv("A", "B")

	assert.Equal(t, "", result)
}
