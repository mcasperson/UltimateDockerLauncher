package envscanners

import (
	b64 "encoding/base64"
	"github.com/mcasperson/UltimateDockerLauncher/cmd/internal/envproviders"
	"github.com/mcasperson/UltimateDockerLauncher/cmd/internal/writers"
	"testing"
)

func TestB64FileWritingTwo(t *testing.T) {
	jsonExample := "{\"whatever\":\"value\"}"
	jsonExampleEncoded := b64.StdEncoding.EncodeToString([]byte("{\"whatever\":\"value\"}"))
	writer := writers.StringWriter{}
	scanner := FileB64WriterEnvScannerTwo{
		Env: envproviders.StringProvider{
			Vars: map[string]string{
				"UDL_WRITEB64FILE_ABC": "[/etc/myapp/settings.json]" + jsonExampleEncoded,
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

func TestB64FileWritingTwoInvalid(t *testing.T) {
	writer := writers.StringWriter{}
	scanner := FileB64WriterEnvScannerTwo{
		Env: envproviders.StringProvider{
			Vars: map[string]string{
				"UDL_WRITEB64FILE_ABC": "[/etc/myapp/settings.json]someinvalidb64encoded",
			},
		},
		Writer: &writer,
	}

	err := scanner.ProcessEnvVars()

	if err != nil {
		t.Fatal(err.Error())
	}

	if writer.Output != nil {
		_, ok := (*writer.Output)["/etc/myapp/settings.json"]

		if ok {
			t.Fatal("The file should not exist")
		}
	}
}
