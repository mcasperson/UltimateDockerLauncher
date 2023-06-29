package inimanipulators

import (
	"github.com/mcasperson/UltimateDockerLauncher/cmd/internal/readers"
	"github.com/mcasperson/UltimateDockerLauncher/cmd/internal/writers"
	"gopkg.in/ini.v1"
	"testing"
)

func TestInvalidFile(t *testing.T) {
	iniExample := "whatever = value"
	writer := writers.StringWriter{}
	reader := readers.StringReader{
		Files: &map[string]string{
			"/etc/config.ini": iniExample,
		},
	}
	manipulator := IniManipulator{
		Writer: &writer,
		Reader: reader,
	}

	if manipulator.CanManipulate("/etc/config.doesnotexist") {
		t.Fatal("This should have failed")
	}
}

func TestInvalidFileExtension(t *testing.T) {
	jsonExample := "hi= there"
	writer := writers.StringWriter{}
	reader := readers.StringReader{
		Files: &map[string]string{
			"/etc/config.json": jsonExample,
		},
	}
	manipulator := IniManipulator{
		Writer: &writer,
		Reader: reader,
	}

	if manipulator.CanManipulate("/etc/config.json") {
		t.Fatal("Must not be able to process files other that .ini")
	}
}

func TestSetInvalidStringField(t *testing.T) {
	jsonExample := "whatever = value"
	writer := writers.StringWriter{}
	reader := readers.StringReader{
		Files: &map[string]string{
			"/etc/config.ini": jsonExample,
		},
	}
	manipulator := IniManipulator{
		Writer: &writer,
		Reader: reader,
	}

	if !manipulator.CanManipulate("/etc/config.ini") {
		t.Fatal("Must be able to manipulate INI files")
	}

	err := manipulator.SetValue("/etc/config.ini", "whatever:blah:foo", "newvalue")

	if err == nil {
		t.Fatal("This should have failed")
	}
}

func TestSetIniStringField(t *testing.T) {
	jsonExample := "whatever = value"
	writer := writers.StringWriter{}
	reader := readers.StringReader{
		Files: &map[string]string{
			"/etc/config.ini": jsonExample,
		},
	}
	manipulator := IniManipulator{
		Writer: &writer,
		Reader: reader,
	}

	if !manipulator.CanManipulate("/etc/config.ini") {
		t.Fatal("Must be able to manipulate INI files")
	}

	err := manipulator.SetValue("/etc/config.ini", "whatever", "newvalue")

	if err != nil {
		t.Fatal("Failed to manipulate INI file: " + err.Error())
	}

	result, err := ini.Load([]byte((*writer.Output)["/etc/config.ini"]))

	value := result.Section("").Key("whatever").Value()

	if value != "newvalue" {
		t.Fatal("Value must be set to \"newvalue\" (was: \"" + value + "\"")
	}
}

func TestSetIniStringGroupField(t *testing.T) {
	jsonExample := "[group]\nwhatever = value"
	writer := writers.StringWriter{}
	reader := readers.StringReader{
		Files: &map[string]string{
			"/etc/config.ini": jsonExample,
		},
	}
	manipulator := IniManipulator{
		Writer: &writer,
		Reader: reader,
	}

	if !manipulator.CanManipulate("/etc/config.ini") {
		t.Fatal("Must be able to manipulate INI files")
	}

	err := manipulator.SetValue("/etc/config.ini", "group:whatever", "newvalue")

	if err != nil {
		t.Fatal("Failed to manipulate INI file: " + err.Error())
	}

	result, err := ini.Load([]byte((*writer.Output)["/etc/config.ini"]))

	value := result.Section("group").Key("whatever").Value()

	if value != "newvalue" {
		t.Fatal("Value must be set to \"newvalue\" (was: \"" + value + "\"")
	}
}

func TestSetIniStringGroupColonField(t *testing.T) {
	jsonExample := "[group]\nwhatever: value"
	writer := writers.StringWriter{}
	reader := readers.StringReader{
		Files: &map[string]string{
			"/etc/config.ini": jsonExample,
		},
	}
	manipulator := IniManipulator{
		Writer: &writer,
		Reader: reader,
	}

	if !manipulator.CanManipulate("/etc/config.ini") {
		t.Fatal("Must be able to manipulate INI files")
	}

	err := manipulator.SetValue("/etc/config.ini", "group:whatever", "newvalue")

	if err != nil {
		t.Fatal("Failed to manipulate INI file: " + err.Error())
	}

	result, err := ini.Load([]byte((*writer.Output)["/etc/config.ini"]))

	value := result.Section("group").Key("whatever").Value()

	if value != "newvalue" {
		t.Fatal("Value must be set to \"newvalue\" (was: \"" + value + "\"")
	}
}

func TestSetIniIntField(t *testing.T) {
	jsonExample := "whatever = 10"
	writer := writers.StringWriter{}
	reader := readers.StringReader{
		Files: &map[string]string{
			"/etc/config.ini": jsonExample,
		},
	}
	manipulator := IniManipulator{
		Writer: &writer,
		Reader: reader,
	}

	if !manipulator.CanManipulate("/etc/config.ini") {
		t.Fatal("Must be able to manipulate INI files")
	}

	err := manipulator.SetValue("/etc/config.ini", "whatever", "5")

	if err != nil {
		t.Fatal("Failed to manipulate INI file: " + err.Error())
	}

	result, err := ini.Load([]byte((*writer.Output)["/etc/config.ini"]))

	value := result.Section("").Key("whatever").Value()

	if value != "5" {
		t.Fatal("Value must be set to \"5\" (was: \"" + value + "\"")
	}
}

func TestSetIniBoolField(t *testing.T) {
	jsonExample := "whatever = false"
	writer := writers.StringWriter{}
	reader := readers.StringReader{
		Files: &map[string]string{
			"/etc/config.ini": jsonExample,
		},
	}
	manipulator := IniManipulator{
		Writer: &writer,
		Reader: reader,
	}

	if !manipulator.CanManipulate("/etc/config.ini") {
		t.Fatal("Must be able to manipulate INI files")
	}

	err := manipulator.SetValue("/etc/config.ini", "whatever", "true")

	if err != nil {
		t.Fatal("Failed to manipulate INI file: " + err.Error())
	}

	result, err := ini.Load([]byte((*writer.Output)["/etc/config.ini"]))

	value := result.Section("").Key("whatever").Value()

	if value != "true" {
		t.Fatal("Value must be set to \"true\" (was: \"" + value + "\"")
	}
}
