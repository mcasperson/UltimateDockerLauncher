package manipulators

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// CommonMapManipulator provides value substitution against a generic map. Since most parsing libraries can unmarhsall
// documents to maps, this struct provides the base for value substitutions.
type CommonMapManipulator struct {
	Unmarshaller Unmarshaller
}

func (m CommonMapManipulator) ProcessMap(result map[string]any, valueSpec string, value string) (map[string]any, error) {
	path := strings.Split(valueSpec, ":")

	var current any = result
	for i, p := range path {

		// If this part of the path is a number, it represents an array index
		if index, err := strconv.ParseInt(p, 10, 16); err == nil {

			if i != len(path)-1 {
				return nil, errors.New("integer indexes must be the final element in the path (index was element " + fmt.Sprint(i+1) + " in a path with " + fmt.Sprint(len(path)) + " elements)")
			}

			objectType := m.getType(current)
			if objectType != "array" {
				return nil, errors.New("integer indexes must be used against an existing array (object type was " + objectType + ")")
			}

			array := current.([]any)

			if int64(len(array)) <= index {
				return nil, errors.New("integer indexes must be within the existing array's bounds (array has " + fmt.Sprint(len(array)) + " elements, index was " + fmt.Sprint(index) + ")")
			}

			switch m.getType(array[index]) {
			case "number":
				number, err := strconv.ParseFloat(value, 64)
				if err != nil {
					array[index] = value
				} else {
					array[index] = number
				}
			case "boolean":
				bool, err := strconv.ParseBool(value)
				if err != nil {
					array[index] = value
				} else {
					array[index] = bool
				}
			case "object":
				var objectValue map[string]any
				err = json.Unmarshal([]byte(value), &objectValue)
				if err != nil {
					array[index] = value
				} else {
					array[index] = objectValue
				}
			case "array":
				var objectValue []any
				err = json.Unmarshal([]byte(value), &objectValue)
				if err != nil {
					array[index] = value
				} else {
					array[index] = objectValue
				}
			default:
				array[index] = value
			}
		} else {
			currentMap, ok := current.(map[string]any)
			if ok {
				if i < len(path)-1 {
					current = currentMap[p]
				} else {
					// Attempt to match the destination type, falling back to a string if the supplied value
					// does not match the destination.
					switch m.getType(currentMap[p]) {
					case "number":
						number, err := strconv.ParseFloat(value, 64)
						if err != nil {
							currentMap[p] = value
						} else {
							currentMap[p] = number
						}
					case "boolean":
						bool, err := strconv.ParseBool(value)
						if err != nil {
							currentMap[p] = value
						} else {
							currentMap[p] = bool
						}
					case "object":
						objectValue, err := m.Unmarshaller.UnmarshalMap(value)
						if err != nil {
							currentMap[p] = value
						} else {
							currentMap[p] = objectValue
						}
					case "array":
						objectValue, err := m.Unmarshaller.UnmarshalArray(value)
						if err != nil {
							currentMap[p] = value
						} else {
							currentMap[p] = objectValue
						}
					default:
						currentMap[p] = value
					}
				}
			} else {
				return nil, errors.New("failed to navigate through JSON object to desired location")
			}
		}
	}

	return result, nil
}

func (m CommonMapManipulator) getType(object any) string {
	if _, ok := object.(int); ok {
		return "number"
	}

	if _, ok := object.(int8); ok {
		return "number"
	}

	if _, ok := object.(int16); ok {
		return "number"
	}

	if _, ok := object.(int32); ok {
		return "number"
	}

	if _, ok := object.(int64); ok {
		return "number"
	}

	if _, ok := object.(float64); ok {
		return "number"
	}

	if _, ok := object.(bool); ok {
		return "boolean"
	}

	if _, ok := object.([]any); ok {
		return "array"
	}

	if _, ok := object.(map[string]any); ok {
		return "object"
	}

	return "string"
}
