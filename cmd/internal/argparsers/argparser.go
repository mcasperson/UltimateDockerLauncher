package argparsers

type ArgParser interface {
	GetExecutable() string
	GetArguments() []string
}
