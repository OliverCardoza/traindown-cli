package internal

import (
	"fmt"
	"os"
)

// EnvFallbackFlag is a custom flag type which retrieves the value from either
// an input flag or env var. If both are set then the flag takes priority.
// This type implements the pflag.Value interface.
type EnvFallbackFlag struct {
	flagName string
	envVar   string
	value    string
}

// NewEnvFallbackFlag constructs an EnvFallbackFlag with the provided env var
// as a fallback location to read values.
func NewEnvFallbackFlag(flagName, envVar string) *EnvFallbackFlag {
	return &EnvFallbackFlag{
		flagName: flagName,
		envVar:   envVar,
	}
}

// String returns a default value as a string. This is only used within the
// pflag library. Callers should prefer to use Get().
func (e *EnvFallbackFlag) String() string {
	return ""
}

// Set is used within the pflag library to configure value based on a provided
// flag value.
func (e *EnvFallbackFlag) Set(input string) error {
	e.value = input
	return nil
}

// Get returns the value preferring the flag input over the env var value.
// If neither are set then an error is returned.
func (e *EnvFallbackFlag) Get() (string, error) {
	if e.value != "" {
		return e.value, nil
	}
	osValue, ok := os.LookupEnv(e.envVar)
	if !ok {
		return "", fmt.Errorf("Error: required flag(s) \"%s\" not set", e.flagName)
	}
	return osValue, nil
}

func (e *EnvFallbackFlag) Type() string {
	return "EnvFallbackFlag"
}
