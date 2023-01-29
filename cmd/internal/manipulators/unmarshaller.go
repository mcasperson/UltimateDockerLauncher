package manipulators

type Unmarshaller interface {
	UnmarshalMap(value string) (map[string]any, error)
	UnmarshalArray(value string) ([]any, error)
}
