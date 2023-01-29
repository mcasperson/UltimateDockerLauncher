package envproviders

type EnvironmentProvider interface {
	GetEnvVar(name string) string
	GetAllEnvVars() []string
}
