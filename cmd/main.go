package main

import (
	"errors"
	"github.com/mcasperson/UltimateDockerLauncher/cmd/internal/argparsers"
	"github.com/mcasperson/UltimateDockerLauncher/cmd/internal/customerror"
	"github.com/mcasperson/UltimateDockerLauncher/cmd/internal/envproviders"
	"github.com/mcasperson/UltimateDockerLauncher/cmd/internal/envscanners"
	"github.com/mcasperson/UltimateDockerLauncher/cmd/internal/executors"
	"github.com/mcasperson/UltimateDockerLauncher/cmd/internal/manipulators"
	inimanipulators "github.com/mcasperson/UltimateDockerLauncher/cmd/internal/manipulators/inimanipulator"
	"github.com/mcasperson/UltimateDockerLauncher/cmd/internal/manipulators/jsonmanipulators"
	"github.com/mcasperson/UltimateDockerLauncher/cmd/internal/manipulators/tomlmanipulators"
	"github.com/mcasperson/UltimateDockerLauncher/cmd/internal/manipulators/yamlmanipulators"
	"github.com/mcasperson/UltimateDockerLauncher/cmd/internal/readers"
	"github.com/mcasperson/UltimateDockerLauncher/cmd/internal/writers"
	"github.com/rs/zerolog"
	"io/fs"
	"os"
)

func setLogging() {
	loggingLevel := os.Getenv("UDL_LOGGING_LEVEL")
	if loggingLevel != "" {
		level, err := zerolog.ParseLevel(loggingLevel)

		if err == nil {
			zerolog.SetGlobalLevel(level)
		}
	}
}

func doScanning() error {
	var envprovider envproviders.EnvironmentProvider = envproviders.EnvVarProvider{}

	var writer writers.Writer = writers.FileWriter{}
	var reader readers.Reader = readers.FileReader{}

	iniManipulator := inimanipulators.IniManipulator{
		Writer: writer,
		Reader: reader,
	}

	jsonManipulator := jsonmanipulators.JsonManipulator{
		Writer: writer,
		Reader: reader,
		MapManipulator: manipulators.CommonMapManipulator{
			Unmarshaller: jsonmanipulators.JsonUnmarshaller{},
		},
	}

	yamlManipulator := yamlmanipulators.YamlManipulator{
		Writer: writer,
		Reader: reader,
		MapManipulator: manipulators.CommonMapManipulator{
			Unmarshaller: yamlmanipulators.YamlUnmarshaller{},
		},
	}

	tomlManipulator := tomlmanipulators.TomlManipulator{
		Writer: writer,
		Reader: reader,
		MapManipulator: manipulators.CommonMapManipulator{
			Unmarshaller: tomlmanipulators.TomlUnmarshaller{},
		},
	}

	scanners := []envscanners.EnvScanner{

		envscanners.FileWriterEnvScanner{
			Writer: writer,
			Env:    envprovider,
		},

		envscanners.FileB64WriterEnvScanner{
			Writer: writer,
			Env:    envprovider,
		},

		envscanners.FileWriterEnvScannerTwo{
			Writer: writer,
			Env:    envprovider,
		},

		envscanners.FileB64WriterEnvScannerTwo{
			Writer: writer,
			Env:    envprovider,
		},

		envscanners.ManipulatorEnvScanner{
			Env: envprovider,
			Manipulator: []manipulators.Manipulator{
				iniManipulator,
				jsonManipulator,
				yamlManipulator,
				tomlManipulator,
			},
		},

		envscanners.ManipulatorEnvScannerTwo{
			Env: envprovider,
			Manipulator: []manipulators.Manipulator{
				iniManipulator,
				jsonManipulator,
				yamlManipulator,
				tomlManipulator,
			},
		},

		envscanners.ManipulatorSkipEmptyEnvScanner{
			Env: envprovider,
			Manipulator: []manipulators.Manipulator{
				iniManipulator,
				jsonManipulator,
				yamlManipulator,
				tomlManipulator,
			},
		},

		envscanners.ManipulatorSkipEmptyEnvScannerTwo{
			Env: envprovider,
			Manipulator: []manipulators.Manipulator{
				iniManipulator,
				jsonManipulator,
				yamlManipulator,
				tomlManipulator,
			},
		},
	}

	for _, scanner := range scanners {
		err := scanner.ProcessEnvVars()

		if err != nil {
			return err
		}
	}

	return nil
}

func systemExit() {
	var argparser argparsers.ArgParser = argparsers.SimpleArgParser{}
	var executor executors.Executor = executors.ExecuteAndWait{}

	// wrap a call to an external executable if supplied
	if argparser.HasExecutable() {
		err := executor.Execute(argparser.GetExecutable(), argparser.GetArguments())
		if exitCode := executor.ExitCode(err); exitCode != 0 {
			os.Exit(exitCode)
		}
	}
}

func main() {
	setLogging()
	err := doScanning()

	if err != nil {

		var customError *customerror.UdlError
		if errors.As(err, &customError) {
			var pathError *fs.PathError
			if errors.As(customError.Err, &pathError) {
				panic("Environment variable \"" + customError.EnvVar + "\"" +
					" ran operation \"" + pathError.Op + "\"" +
					" that failed at path \"" + pathError.Path + "\"" +
					" with error \"" + pathError.Err.Error() + "\"." +
					" This is usually a permission error. Make sure the Docker user has permission to this path.")
			}
		} else {
			panic(err.Error())
		}

	}

	systemExit()
}
