package manipulators

type ValueType int

type Manipulator interface {
	CanManipulate(fileSpec string) bool
	SetValue(fileSpec string, valueSpec string, value string) error
	GetFormatName() string
}

type MapManipulator interface {
	ProcessMap(result map[string]any, valueSpec string, value string) (map[string]any, error)
}
