package envscanners

import (
	"github.com/mcasperson/UltimateDockerLauncher/cmd/internal/customerror"
	"github.com/mcasperson/UltimateDockerLauncher/cmd/internal/envproviders"
	"github.com/mcasperson/UltimateDockerLauncher/cmd/internal/prefixes"
	"github.com/mcasperson/UltimateDockerLauncher/cmd/internal/writers"
	"github.com/rs/zerolog/log"
	"strings"
)

type FileWriterEnvScanner struct {
	Env    envproviders.EnvironmentProvider
	Writer writers.Writer
}

func (f FileWriterEnvScanner) ProcessEnvVars() error {
	for _, e := range f.Env.GetAllEnvVars() {

		if i := strings.Index(e, "="); i >= 0 {
			key := e[:i]
			value := e[i+1:]

			for _, p := range prefixes.EnvVarPrefixes {
				prefix := p + "UDL_WRITEFILE["
				if strings.HasPrefix(key, prefix) && strings.HasSuffix(key, "]") {
					file := strings.TrimLeft(strings.TrimRight(key, "]"), prefix)

					log.Debug().Msg("Writing file \"" + file + "\" with content:")
					log.Debug().Msg(value)

					err := f.Writer.WriteString(file, value)

					if err != nil {
						return &customerror.UdlError{
							EnvVar: key,
							Err:    err,
						}
					}
				}
			}
		}
	}

	return nil
}
