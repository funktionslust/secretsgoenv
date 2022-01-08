# secretsGOenv

go library to set docker secrets as environment variables
## Usage

```go
package main

import "github.com/funktionslust/secretsgoenv"

func main() {
    // available options
    useDockerDefaultPath := "" // in case of linux /run/secrets
    overwriteExistingEnvVars := true
    envVarPrefix := "DOCKER_"

    // must be executed before accessing affected env vars
    err := secretsgoenv.Load(useDockerDefaultPath, overwriteExistingEnvVars, envVarPrefix)
    if err != nil {
        panic(err)
    }

    // e.g. if the secrets dir contains a file "password" with the content "secret"
    // the env var DOCKER_PASSWORD is now set to secret, no matter what DOCKER_PASSWORD was before
}
```
