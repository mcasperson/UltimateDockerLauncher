package yamlmanipulators

import (
	"github.com/mcasperson/UltimateDockerLauncher/cmd/internal/manipulators"
	"github.com/mcasperson/UltimateDockerLauncher/cmd/internal/readers"
	"github.com/mcasperson/UltimateDockerLauncher/cmd/internal/writers"
	"gopkg.in/yaml.v3"
	"strings"
)

type YamlManipulator struct {
	Writer         writers.Writer
	Reader         readers.Reader
	MapManipulator manipulators.MapManipulator
}

func (m YamlManipulator) GetFormatName() string {
	return "YAML"
}

func (m YamlManipulator) CanManipulate(fileSpec string) bool {
	content, err := m.Reader.ReadString(fileSpec)
	if err != nil {
		return false
	}

	var result map[string]any
	err = yaml.Unmarshal([]byte(content), &result)
	return err == nil && (strings.HasSuffix(fileSpec, ".yml") || strings.HasSuffix(fileSpec, ".yaml"))
}

func (m YamlManipulator) SetValue(fileSpec string, valueSpec string, value string) error {
	content, err := m.Reader.ReadString(fileSpec)
	if err != nil {
		return err
	}

	var result map[string]any
	err = yaml.Unmarshal([]byte(content), &result)
	if err != nil {
		return err
	}

	result, err = m.MapManipulator.ProcessMap(result, valueSpec, value)
	if err != nil {
		return err
	}

	json, err := yaml.Marshal(result)
	if err != nil {
		return err
	}
	err = m.Writer.WriteString(fileSpec, string(json))
	return err
}
