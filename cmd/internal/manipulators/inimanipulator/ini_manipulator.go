package inimanipulators

import (
	"errors"
	"github.com/mcasperson/UltimateDockerLauncher/cmd/internal/readers"
	"github.com/mcasperson/UltimateDockerLauncher/cmd/internal/writers"
	"gopkg.in/ini.v1"
	"strings"
)

type IniManipulator struct {
	Writer writers.Writer
	Reader readers.Reader
}

func (m IniManipulator) CanManipulate(fileSpec string) bool {
	content, err := m.Reader.ReadString(fileSpec)
	if err != nil {
		return false
	}

	_, err = ini.Load([]byte(content))
	return err == nil && strings.HasSuffix(fileSpec, ".ini")
}

func (m IniManipulator) SetValue(fileSpec string, valueSpec string, value string) error {
	content, err := m.Reader.ReadString(fileSpec)
	if err != nil {
		return err
	}

	result, err := ini.Load([]byte(content))
	if err != nil {
		return err
	}

	path := strings.Split(valueSpec, ":")

	if !(len(path) == 1 || len(path) == 2) {
		return errors.New("path must be a single key or section and key separated by a colon")
	}

	section := ""
	key := ""
	if len(path) == 2 {
		section = path[0]
		key = path[1]
	} else {
		key = path[0]
	}

	result.Section(section).Key(key).SetValue(value)

	stringWriter := writers.StringIOWriter{}
	_, err = result.WriteTo(&stringWriter)

	return m.Writer.WriteString(fileSpec, stringWriter.Output)
}
