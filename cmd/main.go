package main

import (
	"github.com/mcasperson/UltimateDockerLauncher/cmd/internal/executors"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		panic("Must supply the name of the executable to wrap")
	}

	var executor executors.Executor = executors.ExecuteAndWait{}
	executor.Execute(os.Args[1], getArguments())
}

func getArguments() []string {
	args := []string{}
	if len(os.Args) > 2 {
		args = os.Args[2:]
	}
	return args
}
