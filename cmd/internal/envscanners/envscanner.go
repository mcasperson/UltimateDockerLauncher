package envscanners

type EnvScanner interface {
	ProcessEnvVars() error
}
