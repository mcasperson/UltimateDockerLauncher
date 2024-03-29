package envscanners

import (
	b64 "encoding/base64"
	"github.com/mcasperson/UltimateDockerLauncher/cmd/internal/customerror"
	"github.com/mcasperson/UltimateDockerLauncher/cmd/internal/envproviders"
	"github.com/mcasperson/UltimateDockerLauncher/cmd/internal/prefixes"
	"github.com/mcasperson/UltimateDockerLauncher/cmd/internal/writers"
	"github.com/rs/zerolog/log"
	"strings"
)

type FileB64WriterEnvScanner struct {
	Env    envproviders.EnvironmentProvider
	Writer writers.Writer
}

func (f FileB64WriterEnvScanner) ProcessEnvVars() error {
	for _, e := range f.Env.GetAllEnvVars() {

		if i := strings.Index(e, "="); i >= 0 {
			key := e[:i]
			value := e[i+1:]

			for _, p := range prefixes.EnvVarPrefixes {
				prefix := p + "UDL_WRITEB64FILE["
				if strings.HasPrefix(key, prefix) && strings.HasSuffix(key, "]") {
					file := strings.TrimLeft(strings.TrimRight(key, "]"), prefix)
					contents, err := b64.StdEncoding.DecodeString(value)

					if err != nil {
						log.Error().Msg(value + " is not a valid base64 encoded string. This operation is ignored.")
						return nil
					}

					log.Debug().Msg("Writing file \"" + file + "\" with content:")
					log.Debug().Msg(string(contents))

					err = f.Writer.WriteString(file, string(contents))

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
