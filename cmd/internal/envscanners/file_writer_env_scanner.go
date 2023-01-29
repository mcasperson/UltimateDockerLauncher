package envscanners

import (
	"github.com/mcasperson/UltimateDockerLauncher/cmd/internal/envproviders"
	"github.com/mcasperson/UltimateDockerLauncher/cmd/internal/writers"
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

			if strings.HasPrefix(key, "UDL_WRITEFILE[") && strings.HasSuffix(key, "]") {
				file := strings.TrimLeft(strings.TrimRight(key, "]"), "UDL_WRITEFILE[")
				err := f.Writer.WriteString(file, value)

				if err != nil {
					return err
				}
			}
		}
	}

	return nil
}
