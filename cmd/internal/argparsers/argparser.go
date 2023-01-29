package argparsers

type ArgParser interface {
	HasExecutable() bool
	GetExecutable() string
	GetArguments() []string
}
