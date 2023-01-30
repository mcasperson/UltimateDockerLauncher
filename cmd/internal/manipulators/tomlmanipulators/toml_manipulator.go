package tomlmanipulators

import (
	"github.com/mcasperson/UltimateDockerLauncher/cmd/internal/manipulators"
	"github.com/mcasperson/UltimateDockerLauncher/cmd/internal/readers"
	"github.com/mcasperson/UltimateDockerLauncher/cmd/internal/writers"
	"github.com/pelletier/go-toml/v2"
)

type TomlManipulator struct {
	Writer         writers.Writer
	Reader         readers.Reader
	MapManipulator manipulators.MapManipulator
}

func (m TomlManipulator) CanManipulate(fileSpec string) bool {
	content, err := m.Reader.ReadString(fileSpec)
	if err != nil {
		return false
	}

	var result map[string]any
	err = toml.Unmarshal([]byte(content), &result)
	return err == nil
}

func (m TomlManipulator) SetValue(fileSpec string, valueSpec string, value string) error {
	content, err := m.Reader.ReadString(fileSpec)
	if err != nil {
		return err
	}

	var result map[string]any
	err = toml.Unmarshal([]byte(content), &result)
	if err != nil {
		return err
	}

	result, err = m.MapManipulator.ProcessMap(result, valueSpec, value)
	if err != nil {
		return err
	}

	json, err := toml.Marshal(result)
	if err != nil {
		return err
	}
	err = m.Writer.WriteString(fileSpec, string(json))
	return err
}
