package yamlmanipulators

import (
	"gopkg.in/yaml.v3"
)

type YamlUnmarshaller struct {
}

func (u YamlUnmarshaller) UnmarshalMap(value string) (map[string]any, error) {
	var objectValue map[string]any
	err := yaml.Unmarshal([]byte(value), &objectValue)
	if err != nil {
		return nil, err
	}
	return objectValue, nil
}

func (u YamlUnmarshaller) UnmarshalArray(value string) ([]any, error) {
	var objectValue []any
	err := yaml.Unmarshal([]byte(value), &objectValue)
	if err != nil {
		return nil, err
	}
	return objectValue, nil
}
