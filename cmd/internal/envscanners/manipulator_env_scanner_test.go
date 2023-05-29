package envscanners

import (
	"encoding/json"
	"fmt"
	"github.com/mcasperson/UltimateDockerLauncher/cmd/internal/envproviders"
	"github.com/mcasperson/UltimateDockerLauncher/cmd/internal/manipulators"
	"github.com/mcasperson/UltimateDockerLauncher/cmd/internal/manipulators/jsonmanipulators"
	"github.com/mcasperson/UltimateDockerLauncher/cmd/internal/readers"
	"github.com/mcasperson/UltimateDockerLauncher/cmd/internal/writers"
	"testing"
)

func TestJsonManipulation(t *testing.T) {
	jsonExample := "{\"whatever\":\"hello\"}"
	writer := writers.StringWriter{}
	reader := readers.StringReader{
		Files: &map[string]string{
			"/tmp/myapp/config.json": jsonExample,
		},
	}
	manipulator := ManipulatorEnvScanner{
		Env: envproviders.StringProvider{
			Vars: map[string]string{
				"UDL_SETVALUE[/tmp/myapp/config.json][whatever]": "world",
			},
		},
		Manipulator: jsonmanipulators.JsonManipulator{
			Reader: reader,
			Writer: &writer,
			MapManipulator: manipulators.CommonMapManipulator{
				Unmarshaller: jsonmanipulators.JsonUnmarshaller{},
			},
		},
	}

	err := manipulator.ProcessEnvVars()

	if err != nil {
		t.Fatal(err.Error())
	}

	value, ok := (*writer.Output)["/tmp/myapp/config.json"]

	if !ok {
		t.Fatal("Did not create the expected file")
	}

	var result map[string]any
	err = json.Unmarshal([]byte((*writer.Output)["/tmp/myapp/config.json"]), &result)

	value, ok = result["whatever"].(string)

	if !ok {
		t.Fatal("value must be a string")
	}

	if value != "world" {
		t.Fatal("value must be set to \"world\"")
	}
}

// TestMultipleJsonManipulation verifies that top level properties are set first, followed
// by deeper properties
func TestMultipleJsonManipulation(t *testing.T) {
	jsonExample := "{\"whatever\":[\"hello\"]}"
	files := map[string]string{
		"/tmp/myapp/config.json": jsonExample,
	}

	writer := writers.StringWriter{
		Output: &files,
	}
	reader := readers.StringReader{
		Files: &files,
	}
	manipulator := ManipulatorEnvScanner{
		Env: envproviders.StringProvider{
			Vars: map[string]string{
				"UDL_SETVALUE[/tmp/myapp/config.json][whatever]":   "[1, 2, 3, 4]",
				"UDL_SETVALUE[/tmp/myapp/config.json][whatever:0]": "5",
				"UDL_SETVALUE[/tmp/myapp/config.json][whatever:1]": "6",
				"UDL_SETVALUE[/tmp/myapp/config.json][whatever:2]": "7",
				"UDL_SETVALUE[/tmp/myapp/config.json][whatever:3]": "8",
			},
		},
		Manipulator: jsonmanipulators.JsonManipulator{
			Reader: reader,
			Writer: &writer,
			MapManipulator: manipulators.CommonMapManipulator{
				Unmarshaller: jsonmanipulators.JsonUnmarshaller{},
			},
		},
	}

	err := manipulator.ProcessEnvVars()

	if err != nil {
		t.Fatal(err.Error())
	}

	value, ok := files["/tmp/myapp/config.json"]

	if !ok {
		t.Fatal("Did not create the expected file")
	}

	var result map[string]any
	err = json.Unmarshal([]byte(value), &result)

	property, ok := result["whatever"].([]any)

	if !ok {
		t.Fatal("value must be a string")
	}

	if property[0] != float64(5) {
		t.Fatal("first item must be set to 5, was " + fmt.Sprint(property[0]))
	}

	if property[1] != float64(6) {
		t.Fatal("first item must be set to 6")
	}

	if property[2] != float64(7) {
		t.Fatal("first item must be set to 7")
	}

	if property[3] != float64(8) {
		t.Fatal("first item must be set to 8")
	}
}
