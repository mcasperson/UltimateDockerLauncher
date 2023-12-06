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

func TestJsonManipulationSkipEmpty(t *testing.T) {
	jsonExample := "{\"whatever\":\"hello\"}"
	writer := writers.StringWriter{
		Output: &map[string]string{},
	}
	reader := readers.StringReader{
		Files: &map[string]string{
			"/tmp/myapp/config.json": jsonExample,
		},
	}
	manipulator := ManipulatorSkipEmptyEnvScanner{
		Env: envproviders.StringProvider{
			Vars: map[string]string{
				"UDL_SKIPEMPTY_SETVALUE[/tmp/myapp/config.json][whatever]": "there",
				"UDL_SKIPEMPTY_SETVALUE[/tmp/myapp/config.json][blah]":     "",
			},
		},
		Manipulator: []manipulators.Manipulator{
			jsonmanipulators.JsonManipulator{
				Reader: reader,
				Writer: &writer,
				MapManipulator: manipulators.CommonMapManipulator{
					Unmarshaller: jsonmanipulators.JsonUnmarshaller{},
				},
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

	if value != "there" {
		t.Fatal("value must be set to \"there\"")
	}

	if _, ok = result["blah"].(string); ok {
		t.Fatal("value \"blah\" must be empty")
	}
}

// TestMultipleJsonManipulation verifies that top level properties are set first, followed
// by deeper properties
func TestMultipleJsonManipulationSkipEmpty(t *testing.T) {
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
	manipulator := ManipulatorSkipEmptyEnvScanner{
		Env: envproviders.StringProvider{
			Vars: map[string]string{
				"UDL_SKIPEMPTY_SETVALUE[/tmp/myapp/config.json][whatever]":   "[1, 2, 3, 4]",
				"UDL_SKIPEMPTY_SETVALUE[/tmp/myapp/config.json][whatever:0]": "5",
				"UDL_SKIPEMPTY_SETVALUE[/tmp/myapp/config.json][whatever:1]": "6",
				"UDL_SKIPEMPTY_SETVALUE[/tmp/myapp/config.json][whatever:2]": "7",
				"UDL_SKIPEMPTY_SETVALUE[/tmp/myapp/config.json][whatever:3]": "8",
				"UDL_SKIPEMPTY_SETVALUE[/tmp/myapp/config.json][whatever:4]": "",
				"UDL_SKIPEMPTY_SETVALUE[/tmp/myapp/config.json][whatever:5]": "  ",
			},
		},
		Manipulator: []manipulators.Manipulator{
			jsonmanipulators.JsonManipulator{
				Reader: reader,
				Writer: &writer,
				MapManipulator: manipulators.CommonMapManipulator{
					Unmarshaller: jsonmanipulators.JsonUnmarshaller{},
				},
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

	if len(property) != 4 {
		t.Fatal("must have a length of 4")
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
