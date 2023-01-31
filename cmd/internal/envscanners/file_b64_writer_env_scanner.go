package envscanners

import (
	b64 "encoding/base64"
	"github.com/mcasperson/UltimateDockerLauncher/cmd/internal/envproviders"
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

			if strings.HasPrefix(key, "UDL_WRITEB64FILE[") && strings.HasSuffix(key, "]") {
				file := strings.TrimLeft(strings.TrimRight(key, "]"), "UDL_WRITEB64FILE[")
				contents, err := b64.StdEncoding.DecodeString(value)

				if err != nil {
					return err
				}

				log.Debug().Msg("Writing file \"" + file + "\" with content:")
				log.Debug().Msg(string(contents))

				err = f.Writer.WriteString(file, string(contents))

				if err != nil {
					return err
				}
			}
		}
	}

	return nil
}
