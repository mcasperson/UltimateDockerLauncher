package main

import (
	b64 "encoding/base64"
	"github.com/mcasperson/UltimateDockerLauncher/cmd/internal/prefixes"
	"os"
	"strings"
	"testing"
)

func TestMainFunc(t *testing.T) {
	jsonExample := "{\"whatever\":\"value\"}"
	jsonProcessedExample := "{\"whatever\":\"5\"}"

	file, err := os.CreateTemp("", "file*.json")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(file.Name())

	// This is required to let the app know we are not trying to execute anything based on the passed argyuments
	t.Setenv("UDL_RUNNING_TEST", "true")
	t.Setenv("UDL_WRITEFILE["+file.Name()+"]", jsonExample)
	t.Setenv("UDL_SETVALUE["+file.Name()+"][whatever]", "5")

	// Test the main function
	main()

	if err != nil {
		t.Fatal(err)
	}

	contents, err := os.ReadFile(file.Name())
	if string(contents) != jsonProcessedExample {
		t.Fatal("File contents should have matched")
	}
}

func TestMainJson(t *testing.T) {
	jsonExample := "{\"whatever\":\"value\"}"
	jsonProcessedExample := "{\"whatever\":\"5\"}"

	file, err := os.CreateTemp("", "file*.json")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(file.Name())

	t.Setenv("UDL_WRITEFILE["+file.Name()+"]", jsonExample)
	t.Setenv("UDL_SETVALUE["+file.Name()+"][whatever]", "5")
	err = doScanning()

	if err != nil {
		t.Fatal(err)
	}

	contents, err := os.ReadFile(file.Name())
	if string(contents) != jsonProcessedExample {
		t.Fatal("File contents should have matched")
	}
}

func TestMainJsonWithPrefix(t *testing.T) {
	jsonExample := "{\"whatever\":\"value\"}"
	jsonProcessedExample := "{\"whatever\":\"5\"}"

	file, err := os.CreateTemp("", "file*.json")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(file.Name())

	t.Setenv(prefixes.EnvVarPrefixes[0]+"UDL_WRITEFILE["+file.Name()+"]", jsonExample)
	t.Setenv(prefixes.EnvVarPrefixes[0]+"UDL_SETVALUE["+file.Name()+"][whatever]", "5")
	err = doScanning()

	if err != nil {
		t.Fatal(err)
	}

	contents, err := os.ReadFile(file.Name())
	if string(contents) != jsonProcessedExample {
		t.Fatal("File contents should have matched")
	}
}

func TestMainJsonTwo(t *testing.T) {
	jsonExample := "{\"whatever\":\"value\"}"
	jsonProcessedExample := "{\"whatever\":\"5\"}"

	file, err := os.CreateTemp("", "file*.json")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(file.Name())

	t.Setenv("UDL_WRITEFILE["+file.Name()+"]", jsonExample)
	t.Setenv("UDL_SETVALUE_1", "["+file.Name()+"][whatever]5")
	err = doScanning()

	if err != nil {
		t.Fatal(err)
	}

	contents, err := os.ReadFile(file.Name())
	if string(contents) != jsonProcessedExample {
		t.Fatal("File contents should have matched")
	}
}

func TestMainJsonWithPrefixTwo(t *testing.T) {
	jsonExample := "{\"whatever\":\"value\"}"
	jsonProcessedExample := "{\"whatever\":\"5\"}"

	file, err := os.CreateTemp("", "file*.json")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(file.Name())

	t.Setenv(prefixes.EnvVarPrefixes[0]+"UDL_WRITEFILE["+file.Name()+"]", jsonExample)
	t.Setenv(prefixes.EnvVarPrefixes[0]+"UDL_SETVALUE_1", "["+file.Name()+"][whatever]5")
	err = doScanning()

	if err != nil {
		t.Fatal(err)
	}

	contents, err := os.ReadFile(file.Name())
	if string(contents) != jsonProcessedExample {
		t.Fatal("File contents should have matched")
	}
}

func TestMainToml(t *testing.T) {
	tomlExample := "whatever = 'value'"
	tomlExampleExample := "whatever = '5'"

	file, err := os.CreateTemp("", "file*.toml")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(file.Name())

	t.Setenv("UDL_WRITEFILE["+file.Name()+"]", tomlExample)
	t.Setenv("UDL_SETVALUE["+file.Name()+"][whatever]", "5")
	err = doScanning()

	if err != nil {
		t.Fatal(err)
	}

	contents, err := os.ReadFile(file.Name())
	contentsString := strings.TrimSpace(string(contents))
	if contentsString != tomlExampleExample {
		t.Fatal("File contents should have matched. Was " + contentsString + " expected " + tomlExampleExample)
	}
}

func TestMainTomlTwo(t *testing.T) {
	tomlExample := "whatever = 'value'"
	tomlExampleExample := "whatever = '5'"

	file, err := os.CreateTemp("", "file*.toml")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(file.Name())

	t.Setenv("UDL_WRITEFILE["+file.Name()+"]", tomlExample)
	t.Setenv("UDL_SETVALUE_1", "["+file.Name()+"][whatever]5")
	err = doScanning()

	if err != nil {
		t.Fatal(err)
	}

	contents, err := os.ReadFile(file.Name())
	contentsString := strings.TrimSpace(string(contents))
	if contentsString != tomlExampleExample {
		t.Fatal("File contents should have matched. Was " + contentsString + " expected " + tomlExampleExample)
	}
}

func TestMainYaml(t *testing.T) {
	yamlExample := "whatever: \"value\""
	yamlExampleProcessed := "whatever: \"5\""

	file, err := os.CreateTemp("", "file*.yaml")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(file.Name())

	t.Setenv("UDL_WRITEFILE["+file.Name()+"]", yamlExample)
	t.Setenv("UDL_SETVALUE["+file.Name()+"][whatever]", "5")
	err = doScanning()

	if err != nil {
		t.Fatal(err)
	}

	contents, err := os.ReadFile(file.Name())
	contentsString := strings.TrimSpace(string(contents))
	if contentsString != yamlExampleProcessed {
		t.Fatal("File contents should have matched. Was " + contentsString + " expected " + yamlExampleProcessed)
	}
}

func TestMainYamlTwo(t *testing.T) {
	yamlExample := "whatever: \"value\""
	yamlExampleProcessed := "whatever: \"5\""

	file, err := os.CreateTemp("", "file*.yaml")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(file.Name())

	t.Setenv("UDL_WRITEFILE["+file.Name()+"]", yamlExample)
	t.Setenv("UDL_SETVALUE_1", "["+file.Name()+"][whatever]5")
	err = doScanning()

	if err != nil {
		t.Fatal(err)
	}

	contents, err := os.ReadFile(file.Name())
	contentsString := strings.TrimSpace(string(contents))
	if contentsString != yamlExampleProcessed {
		t.Fatal("File contents should have matched. Was " + contentsString + " expected " + yamlExampleProcessed)
	}
}

func TestMainIni(t *testing.T) {
	iniExample := "whatever = value"
	iniExampleProcessed := "whatever = 5"

	file, err := os.CreateTemp("", "file*.ini")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(file.Name())

	t.Setenv("UDL_WRITEFILE["+file.Name()+"]", iniExample)
	t.Setenv("UDL_SETVALUE["+file.Name()+"][whatever]", "5")
	err = doScanning()

	if err != nil {
		t.Fatal(err)
	}

	contents, err := os.ReadFile(file.Name())
	contentsString := strings.TrimSpace(string(contents))
	if contentsString != iniExampleProcessed {
		t.Fatal("File contents should have matched. Was " + contentsString + " expected " + iniExampleProcessed)
	}
}

func TestMainIniTwo(t *testing.T) {
	iniExample := "whatever = value"
	iniExampleProcessed := "whatever = 5"

	file, err := os.CreateTemp("", "file*.ini")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(file.Name())

	t.Setenv("UDL_WRITEFILE["+file.Name()+"]", iniExample)
	t.Setenv("UDL_SETVALUE_1", "["+file.Name()+"][whatever]5")
	err = doScanning()

	if err != nil {
		t.Fatal(err)
	}

	contents, err := os.ReadFile(file.Name())
	contentsString := strings.TrimSpace(string(contents))
	if contentsString != iniExampleProcessed {
		t.Fatal("File contents should have matched. Was " + contentsString + " expected " + iniExampleProcessed)
	}
}

func TestMainWriteFile(t *testing.T) {
	jsonExample := "{\"whatever\":\"value\"}"

	file, err := os.CreateTemp("", "file.json")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(file.Name())

	t.Setenv("UDL_WRITEFILE["+file.Name()+"]", jsonExample)
	err = doScanning()

	if err != nil {
		t.Fatal(err)
	}

	contents, err := os.ReadFile(file.Name())
	if string(contents) != jsonExample {
		t.Fatal("File contents should have matched")
	}
}

func TestMainWriteFileTwo(t *testing.T) {
	jsonExample := "{\"whatever\":\"value\"}"

	file, err := os.CreateTemp("", "file.json")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(file.Name())

	t.Setenv("UDL_WRITEFILE_1", "["+file.Name()+"]"+jsonExample)
	err = doScanning()

	if err != nil {
		t.Fatal(err)
	}

	contents, err := os.ReadFile(file.Name())
	if string(contents) != jsonExample {
		t.Fatal("File contents should have matched")
	}
}

func TestMainBase64(t *testing.T) {
	jsonExample := "{\"whatever\":\"value\"}"
	jsonExampleEncoded := b64.StdEncoding.EncodeToString([]byte(jsonExample))

	file, err := os.CreateTemp("", "file.json")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(file.Name())

	t.Setenv("UDL_WRITEB64FILE["+file.Name()+"]", jsonExampleEncoded)
	err = doScanning()

	if err != nil {
		t.Fatal(err)
	}

	contents, err := os.ReadFile(file.Name())
	if string(contents) != jsonExample {
		t.Fatal("File contents should have matched")
	}
}

func TestMainBase64Two(t *testing.T) {
	jsonExample := "{\"whatever\":\"value\"}"
	jsonExampleEncoded := b64.StdEncoding.EncodeToString([]byte(jsonExample))

	file, err := os.CreateTemp("", "file.json")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(file.Name())

	t.Setenv("UDL_WRITEB64FILE_1", "["+file.Name()+"]"+jsonExampleEncoded)
	err = doScanning()

	if err != nil {
		t.Fatal(err)
	}

	contents, err := os.ReadFile(file.Name())
	if string(contents) != jsonExample {
		t.Fatal("File contents should have matched")
	}
}

func TestMainBase64WithProefix(t *testing.T) {
	jsonExample := "{\"whatever\":\"value\"}"
	jsonExampleEncoded := b64.StdEncoding.EncodeToString([]byte(jsonExample))

	file, err := os.CreateTemp("", "file.json")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(file.Name())

	t.Setenv(prefixes.EnvVarPrefixes[0]+"UDL_WRITEB64FILE["+file.Name()+"]", jsonExampleEncoded)
	err = doScanning()

	if err != nil {
		t.Fatal(err)
	}

	contents, err := os.ReadFile(file.Name())
	if string(contents) != jsonExample {
		t.Fatal("File contents should have matched")
	}
}

func TestMainBase64WithPrefixTwo(t *testing.T) {
	jsonExample := "{\"whatever\":\"value\"}"
	jsonExampleEncoded := b64.StdEncoding.EncodeToString([]byte(jsonExample))

	file, err := os.CreateTemp("", "file.json")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(file.Name())

	t.Setenv(prefixes.EnvVarPrefixes[0]+"UDL_WRITEB64FILE_1", "["+file.Name()+"]"+jsonExampleEncoded)
	err = doScanning()

	if err != nil {
		t.Fatal(err)
	}

	contents, err := os.ReadFile(file.Name())
	if string(contents) != jsonExample {
		t.Fatal("File contents should have matched")
	}
}
