package writers

import (
	"github.com/mcasperson/UltimateDockerLauncher/cmd/internal/readers"
	"os"
	"testing"
)

func TestFileWriting(t *testing.T) {
	file, err := os.CreateTemp("", "output")
	if err != nil {
		t.Fatal(err.Error())
	}
	defer os.Remove(file.Name())

	jsonExample := "{\"whatever\":\"value\"}"

	reader := readers.FileReader{}

	writer := FileWriter{}
	err = writer.WriteString(file.Name(), jsonExample)

	if err != nil {
		t.Fatal(err.Error())
	}

	value, err := reader.ReadString(file.Name())

	if err != nil {
		t.Fatal(err.Error())
	}

	if value != jsonExample {
		t.Fatal("Did not save the expected content")
	}
}
