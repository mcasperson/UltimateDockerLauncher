package main

import (
	b64 "encoding/base64"
	"os"
	"testing"
)

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
