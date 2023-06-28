package envscanners

import (
	"errors"
	"github.com/mcasperson/UltimateDockerLauncher/cmd/internal/envproviders"
	"github.com/mcasperson/UltimateDockerLauncher/cmd/internal/writers"
	"github.com/rs/zerolog/log"
	"regexp"
	"strings"
)

type FileWriterEnvScannerTwo struct {
	Env    envproviders.EnvironmentProvider
	Writer writers.Writer
}

func (f FileWriterEnvScannerTwo) ProcessEnvVars() error {
	for _, e := range f.Env.GetAllEnvVars() {

		if i := strings.Index(e, "="); i >= 0 {
			key := e[:i]
			value := e[i+1:]

			if strings.HasPrefix(key, "UDL_WRITEFILE_") {
				file, contents, err := f.getFilePath(value)

				if err != nil {
					return err
				}

				log.Debug().Msg("Writing file \"" + file + "\" with content:")
				log.Debug().Msg(contents)

				err = f.Writer.WriteString(file, contents)

				if err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func (f FileWriterEnvScannerTwo) getFilePath(key string) (string, string, error) {
	rgx := regexp.MustCompile("\\[([^\\[\\]]+)](.*)")
	rs := rgx.FindStringSubmatch(key)

	if rs != nil && len(rs) == 3 {
		return rs[1], rs[2], nil
	}

	return "", "", errors.New("failed to match the value to the regex")
}
