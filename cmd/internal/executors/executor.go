package executors

type Executor interface {
	Execute(executable string, args []string) error
	ExitCode(err error) int
}
