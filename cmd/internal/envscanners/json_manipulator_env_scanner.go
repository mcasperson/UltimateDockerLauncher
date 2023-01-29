package envscanners

import (
	"github.com/mcasperson/UltimateDockerLauncher/cmd/internal/envproviders"
	"github.com/mcasperson/UltimateDockerLauncher/cmd/internal/manipulators"
	"github.com/mcasperson/UltimateDockerLauncher/cmd/internal/stringutil"
	"github.com/rs/zerolog/log"
	"strings"
)

type ManipulatorEnvScanner struct {
	Env         envproviders.EnvironmentProvider
	Manipulator manipulators.Manipulator
}

func (f ManipulatorEnvScanner) ProcessEnvVars() error {
	for _, e := range f.Env.GetAllEnvVars() {

		if i := strings.Index(e, "="); i >= 0 {
			key := e[:i]
			value := e[i+1:]

			if strings.HasPrefix(key, "UDL_SETVALUE[") && strings.HasSuffix(key, "]") {
				file := stringutil.Substr(key, len("UDL_SETVALUE["), strings.Index(key, "]")-len("UDL_SETVALUE["))
				path := stringutil.Substr(key, strings.Index(key, "][")+2, len(key)-strings.Index(key, "][")-3)

				log.Debug().Msg("Attempting to parse " + file + " as JSON and modify value at " + path)

				if f.Manipulator.CanManipulate(file) {
					log.Debug().Msg("Successfully parsed " + file + " as JSON")
					err := f.Manipulator.SetValue(file, path, value)
					if err != nil {
						return err
					}
				} else {
					log.Debug().Msg("Could not parse " + file + " as JSON")
				}
			}
		}
	}

	return nil
}
