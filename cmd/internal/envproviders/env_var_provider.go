package envproviders

import "os"

type EnvVarProvider struct {
}

func (e EnvVarProvider) GetEnvVar(name string) string {
	return os.Getenv(name)
}

func (e EnvVarProvider) GetAllEnvVars() []string {
	return os.Environ()
}
