package argparsers

import "os"

type SimpleArgParser struct {
}

func (a SimpleArgParser) GetExecutable() string {
	return os.Args[1]
}

func (a SimpleArgParser) GetArguments() []string {
	args := []string{}
	if len(os.Args) > 2 {
		args = os.Args[2:]
	}
	return args
}
