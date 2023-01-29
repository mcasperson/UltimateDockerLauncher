package envscanners

import (
	"encoding/json"
	"github.com/mcasperson/UltimateDockerLauncher/cmd/internal/envproviders"
	"github.com/mcasperson/UltimateDockerLauncher/cmd/internal/manipulators"
	"github.com/mcasperson/UltimateDockerLauncher/cmd/internal/readers"
	"github.com/mcasperson/UltimateDockerLauncher/cmd/internal/writers"
	"testing"
)

func TestJsonManipulation(t *testing.T) {
	jsonExample := "{\"whatever\":\"hello\"}"
	writer := writers.StringWriter{}
	reader := readers.StringReader{
		Files: map[string]string{
			"/etc/myapp/config.json": jsonExample,
		},
	}
	manipulator := ManipulatorEnvScanner{
		Env: envproviders.StringProvider{
			Vars: map[string]string{
				"UDL_SETVALUE[/etc/myapp/config.json][whatever]": "world",
			},
		},
		Manipulator: manipulators.JsonManipulator{
			Reader: reader,
			Writer: &writer,
		},
	}

	err := manipulator.ProcessEnvVars()

	if err != nil {
		t.Fatal(err.Error())
	}

	value, ok := writer.Output["/etc/myapp/config.json"]

	if !ok {
		t.Fatal("Did not create the expected file")
	}

	var result map[string]any
	err = json.Unmarshal([]byte(writer.Output["/etc/myapp/config.json"]), &result)

	value, ok = result["whatever"].(string)

	if !ok {
		t.Fatal("value must be a string")
	}

	if value != "world" {
		t.Fatal("value must be set to \"world\"")
	}
}
