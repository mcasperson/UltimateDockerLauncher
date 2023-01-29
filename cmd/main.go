package main

import (
	"github.com/mcasperson/UltimateDockerLauncher/cmd/internal/argparsers"
	"github.com/mcasperson/UltimateDockerLauncher/cmd/internal/envproviders"
	"github.com/mcasperson/UltimateDockerLauncher/cmd/internal/envscanners"
	"github.com/mcasperson/UltimateDockerLauncher/cmd/internal/executors"
	"github.com/mcasperson/UltimateDockerLauncher/cmd/internal/manipulators"
	"github.com/mcasperson/UltimateDockerLauncher/cmd/internal/readers"
	"github.com/mcasperson/UltimateDockerLauncher/cmd/internal/writers"
	"github.com/rs/zerolog"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		panic("Must supply the name of the executable to wrap")
	}

	loggingLevel := os.Getenv("UDL_LOGGING_LEVEL")
	if loggingLevel != "" {
		level, err := zerolog.ParseLevel(loggingLevel)

		if err == nil {
			zerolog.SetGlobalLevel(level)
		}
	}

	var envprovider envproviders.EnvironmentProvider = envproviders.EnvVarProvider{}
	var argparser argparsers.ArgParser = argparsers.SimpleArgParser{}
	var executor executors.Executor = executors.ExecuteAndWait{}
	var writer writers.Writer = writers.FileWriter{}
	var reader readers.Reader = readers.FileReader{}
	var writeFileScanner envscanners.EnvScanner = envscanners.FileWriterEnvScanner{
		Writer: writer,
		Env:    envprovider,
	}
	var jsonManipluator envscanners.EnvScanner = envscanners.ManipulatorEnvScanner{
		Env: envprovider,
		Manipulator: manipulators.JsonManipulator{
			Writer: writer,
			Reader: reader,
		},
	}

	err := writeFileScanner.ProcessEnvVars()

	if err != nil {
		panic(err.Error())
	}

	err = jsonManipluator.ProcessEnvVars()

	if err != nil {
		panic(err.Error())
	}

	executor.Execute(argparser.GetExecutable(), argparser.GetArguments())
}
