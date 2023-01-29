package executors

type Executor interface {
	Execute(executable string, args []string) error
}
