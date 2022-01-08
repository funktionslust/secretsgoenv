package secretsgoenv_test

import (
	"os"
	"testing"

	"github.com/funktionslust/secretsgoenv"
)

const (
	dir           = "test"
	prefix        = "DOCKER_"
	key           = "PASSWORD"
	secretValue   = "secret"
	originalValue = "terces"
)

func TestLoad(t *testing.T) {
	// Setup test environment variable PASSWORD=terces
	err := os.Setenv(key, originalValue)
	if err != nil {
		t.Errorf("os.Setenv: %v", err)
	}
	if os.Getenv(key) != originalValue {
		t.Errorf("originalValue not set: %v", originalValue)
	}
	// Verify that overwrite == false keeps original value "terces"
	overwrite := false
	err = secretsgoenv.Load(dir, overwrite, "")
	if err != nil {
		t.Errorf("secretsgoenv.Load: %v", err)
	}
	if os.Getenv(key) != originalValue {
		t.Error("secretsgoenv.Load must not be overwritten")
	}
	// Verify that overwrite == true changes original value to "secret"
	overwrite = true
	err = secretsgoenv.Load(dir, overwrite, "")
	if err != nil {
		t.Errorf("secretsgoenv.Load: %v", err)
	}
	if os.Getenv(key) != secretValue {
		t.Errorf("secretsgoenv.Load should have been overwritten to `%v` but is `%v`", secretValue, os.Getenv(key))
	}
	os.Setenv(key, originalValue) // reset originalValue
	// Verify prefix works, and original value is preserved
	err = secretsgoenv.Load(dir, overwrite, prefix)
	if err != nil {
		t.Errorf("secretsgoenv.Load: %v", err)
	}
	if os.Getenv(prefix+key) != secretValue {
		t.Errorf("secretsgoenv.Load with prefix should have been overwritten to `%v` but is `%v`", secretValue, os.Getenv(prefix+key))
	}
	if os.Getenv(key) != originalValue {
		t.Errorf("secretsgoenv.Load with prefix must not change the original value `%v` but is `%v`", originalValue, os.Getenv(key))
	}
}
