package envscanners

import (
	"github.com/mcasperson/UltimateDockerLauncher/cmd/internal/envproviders"
	"github.com/mcasperson/UltimateDockerLauncher/cmd/internal/manipulators"
	"github.com/mcasperson/UltimateDockerLauncher/cmd/internal/stringutil"
	"github.com/rs/zerolog/log"
	"regexp"
	"sort"
	"strings"
)

type ManipulatorEnvScanner struct {
	Env         envproviders.EnvironmentProvider
	Manipulator []manipulators.Manipulator
}

func (f ManipulatorEnvScanner) getFilePath(key string) (string, string) {
	file := stringutil.Substr(key, len("UDL_SETVALUE["), strings.Index(key, "]")-len("UDL_SETVALUE["))
	path := stringutil.Substr(key, strings.Index(key, "][")+2, len(key)-strings.Index(key, "][")-3)
	return file, path
}

func (f ManipulatorEnvScanner) getVars() ([]int, map[int][]string) {
	orderedVars := map[int][]string{}
	orderedVarsKeys := []int{}

	// start by getting matching env vars, splitting the values by colon, and saving
	// the env var in a map keyed by the length of the accessor
	for _, e := range f.Env.GetAllEnvVars() {

		if i := strings.Index(e, "="); i >= 0 {
			key := e[:i]
			match, _ := regexp.MatchString("UDL_SETVALUE\\[[^\\[\\]]+]\\[[^\\[\\]]+]", key)

			if match {
				_, path := f.getFilePath(key)
				splitValue := strings.Split(path, ":")
				if _, ok := orderedVars[len(splitValue)]; !ok {
					orderedVars[len(splitValue)] = []string{}
					orderedVarsKeys = append(orderedVarsKeys, len(splitValue))
				}

				orderedVars[len(splitValue)] = append(orderedVars[len(splitValue)], key)
			}
		}
	}

	// Sort the array of accessor lengths
	sort.Ints(orderedVarsKeys)

	return orderedVarsKeys, orderedVars
}

func (f ManipulatorEnvScanner) ProcessEnvVars() error {
	orderedVarsKeys, orderedVars := f.getVars()

	// Starting with the accessors of the shortest length, process to injections.
	// This means that top level properties are set first, and deeper properties
	// are set last. It also means the deeper accessors can modify the previously
	// injected values.
	for _, length := range orderedVarsKeys {
		for _, key := range orderedVars[length] {
			value := f.Env.GetEnvVar(key)
			file, path := f.getFilePath(key)

			for _, manipulator := range f.Manipulator {

				log.Debug().Msg("Attempting to parse " + file + " as " + manipulator.GetFormatName() + " and modify value at " + path)

				if manipulator.CanManipulate(file) {
					log.Debug().Msg("Successfully parsed " + file + " as " + manipulator.GetFormatName())
					err := manipulator.SetValue(file, path, value)
					if err != nil {
						return err
					}
				} else {
					log.Debug().Msg("Could not parse " + file + " as " + manipulator.GetFormatName())
				}
			}
		}
	}

	return nil
}
