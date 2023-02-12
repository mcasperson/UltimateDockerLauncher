package envscanners

import (
	"github.com/mcasperson/UltimateDockerLauncher/cmd/internal/envproviders"
	"github.com/mcasperson/UltimateDockerLauncher/cmd/internal/writers"
	"testing"
)

func TestFileWriting(t *testing.T) {
	jsonExample := "{\"whatever\":\"value\"}"
	writer := writers.StringWriter{}
	scanner := FileWriterEnvScanner{
		Env: envproviders.StringProvider{
			Vars: map[string]string{
				"UDL_WRITEFILE[/etc/myapp/settings.json]": jsonExample,
			},
		},
		Writer: &writer,
	}

	err := scanner.ProcessEnvVars()

	if err != nil {
		t.Fatal(err.Error())
	}

	value, ok := (*writer.Output)["/etc/myapp/settings.json"]

	if !ok {
		t.Fatal("Did not create the expected file")
	}

	if value != jsonExample {
		t.Fatal("Did not save the expected content")
	}
}
