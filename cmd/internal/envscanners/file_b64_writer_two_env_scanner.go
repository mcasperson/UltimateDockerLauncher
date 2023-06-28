package envscanners

import (
	b64 "encoding/base64"
	"errors"
	"github.com/mcasperson/UltimateDockerLauncher/cmd/internal/envproviders"
	"github.com/mcasperson/UltimateDockerLauncher/cmd/internal/writers"
	"github.com/rs/zerolog/log"
	"regexp"
	"strings"
)

type FileB64WriterEnvScannerTwo struct {
	Env    envproviders.EnvironmentProvider
	Writer writers.Writer
}

func (f FileB64WriterEnvScannerTwo) ProcessEnvVars() error {
	for _, e := range f.Env.GetAllEnvVars() {

		if i := strings.Index(e, "="); i >= 0 {
			key := e[:i]
			value := e[i+1:]

			if strings.HasPrefix(key, "UDL_WRITEB64FILE_") {
				file, encodedContents, err := f.getFilePath(value)

				if err != nil {
					return err
				}

				contents, err := b64.StdEncoding.DecodeString(encodedContents)

				if err != nil {
					log.Error().Msg(encodedContents + " is not a valid base64 encoded string. This operation is ignored.")
					return nil
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

func (f FileB64WriterEnvScannerTwo) getFilePath(key string) (string, string, error) {
	rgx := regexp.MustCompile("\\[([^\\[\\]]+)](.*)")
	rs := rgx.FindStringSubmatch(key)

	if rs != nil && len(rs) == 3 {
		return rs[1], rs[2], nil
	}

	return "", "", errors.New("failed to match the value to the regex")
}
