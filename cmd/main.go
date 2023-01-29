package main

import (
	"github.com/mcasperson/UltimateDockerLauncher/cmd/internal/argparsers"
	"github.com/mcasperson/UltimateDockerLauncher/cmd/internal/envproviders"
	"github.com/mcasperson/UltimateDockerLauncher/cmd/internal/envscanners"
	"github.com/mcasperson/UltimateDockerLauncher/cmd/internal/executors"
	"github.com/mcasperson/UltimateDockerLauncher/cmd/internal/writers"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		panic("Must supply the name of the executable to wrap")
	}

	var envprovider envproviders.EnvironmentProvider = envproviders.EnvVarProvider{}
	var argparser argparsers.ArgParser = argparsers.SimpleArgParser{}
	var executor executors.Executor = executors.ExecuteAndWait{}
	var writer writers.Writer = writers.FileWriter{}
	var writeFileScanner envscanners.EnvScanner = envscanners.FileWriterEnvScanner{
		Writer: writer,
		Env:    envprovider,
	}

	err := writeFileScanner.ProcessEnvVars()

	if err != nil {
		panic(err.Error())
	}

	executor.Execute(argparser.GetExecutable(), argparser.GetArguments())
}
