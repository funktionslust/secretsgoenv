package secretsgoenv

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

// Load sets docker secrets as environment variables
// empty dir defaults to docker default secret locations
// overwrite gives docker secrets precedence over already set environment variables
// prefix extends the environment variable name
// environment variable names get converted to upper case
// docs: https://docs.docker.com/engine/swarm/secrets/#how-docker-manages-secrets
func Load(dir string, overwrite bool, prefix string) error {
	if dir == "" {
		if runtime.GOOS == "windows" {
			dir = filepath.FromSlash("C:\\ProgramData\\Docker\\secrets")
		} else {
			dir = filepath.FromSlash("/run/secrets")
		}
	}
	fileInfo, err := os.Stat(dir)
	if err != nil {
		return err
	}
	if !fileInfo.IsDir() {
		return fmt.Errorf("%s is no directory", dir)
	}
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return err
	}
	for _, file := range files {
		buf, err := ioutil.ReadFile(filepath.FromSlash(fmt.Sprintf("%s/%s", dir, file.Name())))
		if err != nil {
			return err
		}
		name := fmt.Sprintf("%s%s", prefix, strings.ToUpper(file.Name()))
		secretValue := strings.TrimSpace(string(buf))
		_, exists := os.LookupEnv(name)
		if !exists || (exists && overwrite) {
			err := os.Setenv(name, secretValue)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
