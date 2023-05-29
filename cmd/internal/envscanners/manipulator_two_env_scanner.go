package envscanners

import (
	"errors"
	"github.com/mcasperson/UltimateDockerLauncher/cmd/internal/envproviders"
	"github.com/mcasperson/UltimateDockerLauncher/cmd/internal/manipulators"
	"github.com/rs/zerolog/log"
	"regexp"
	"sort"
	"strings"
)

// ManipulatorEnvScannerTwo exists because even though it is legal to have brackets in env var names in Linux,
// it is not legal to have them in Kubernetes. This manipulator uses plain env var names and defines the
// filename, accessor, and new value in the env var value.
type ManipulatorEnvScannerTwo struct {
	Env         envproviders.EnvironmentProvider
	Manipulator []manipulators.Manipulator
}

func (f ManipulatorEnvScannerTwo) getFilePath(key string) (string, string, string, error) {
	rgx := regexp.MustCompile("\\[([^\\[\\]]+)]\\[([^\\[\\]]+)](.*)")
	rs := rgx.FindStringSubmatch(key)

	if rs != nil && len(rs) == 4 {
		return rs[1], rs[2], rs[3], nil
	}

	return "", "", "", errors.New("failed to match the value to the regex")
}

func (f ManipulatorEnvScannerTwo) getVars() ([]int, map[int][]string) {
	orderedVars := map[int][]string{}
	orderedVarsKeys := []int{}

	// start by getting matching env vars, splitting the values by colon, and saving
	// the env var in a map keyed by the length of the accessor
	for _, e := range f.Env.GetAllEnvVars() {

		if i := strings.Index(e, "="); i >= 0 {
			key := e[:i]
			value := e[i-1:]

			if match, _ := regexp.MatchString("UDL_SETVALUE_[-._a-zA-Z0-9]+", key); !match {
				continue
			}

			_, accessor, _, err := f.getFilePath(value)

			if err == nil {
				splitPath := strings.Split(accessor, ":")
				if _, ok := orderedVars[len(splitPath)]; !ok {
					orderedVars[len(splitPath)] = []string{}
					orderedVarsKeys = append(orderedVarsKeys, len(splitPath))
				}

				orderedVars[len(splitPath)] = append(orderedVars[len(splitPath)], key)
			}
		}
	}

	// Sort the array of accessor lengths
	sort.Ints(orderedVarsKeys)

	return orderedVarsKeys, orderedVars
}

func (f ManipulatorEnvScannerTwo) ProcessEnvVars() error {
	orderedVarsKeys, orderedVars := f.getVars()

	// Starting with the accessors of the shortest length, process to injections.
	// This means that top level properties are set first, and deeper properties
	// are set last. It also means the deeper accessors can modify the previously
	// injected values.
	for _, length := range orderedVarsKeys {
		for _, key := range orderedVars[length] {
			value := f.Env.GetEnvVar(key)
			file, accessor, newValue, err := f.getFilePath(value)

			if err != nil {
				log.Debug().Msg("Could not parse " + value + " as a manipulation directive")
				continue
			}

			for _, manipulator := range f.Manipulator {

				log.Debug().Msg("Attempting to parse " + file + " as " + manipulator.GetFormatName() + " and modify value at " + accessor)

				if manipulator.CanManipulate(file) {
					log.Debug().Msg("Successfully parsed " + file + " as " + manipulator.GetFormatName())
					err := manipulator.SetValue(file, accessor, newValue)
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
