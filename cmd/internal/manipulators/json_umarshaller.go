package manipulators

import "encoding/json"

type JsonUnmarshaller struct {
}

func (u JsonUnmarshaller) UnmarshalMap(value string) (map[string]any, error) {
	var objectValue map[string]any
	err := json.Unmarshal([]byte(value), &objectValue)
	if err != nil {
		return nil, err
	}
	return objectValue, nil
}

func (u JsonUnmarshaller) UnmarshalArray(value string) ([]any, error) {
	var objectValue []any
	err := json.Unmarshal([]byte(value), &objectValue)
	if err != nil {
		return nil, err
	}
	return objectValue, nil
}
