package manipulators

import "github.com/pelletier/go-toml/v2"

type TomlUnmarshaller struct {
}

func (u TomlUnmarshaller) UnmarshalMap(value string) (map[string]any, error) {
	var objectValue map[string]any
	err := toml.Unmarshal([]byte(value), &objectValue)
	if err != nil {
		return nil, err
	}
	return objectValue, nil
}

func (u TomlUnmarshaller) UnmarshalArray(value string) ([]any, error) {
	// We assign the array to a property and then extract it again
	var objectValue map[string][]any
	err := toml.Unmarshal([]byte("array = "+value), &objectValue)
	if err != nil {
		return nil, err
	}
	return objectValue["array"], nil
}
