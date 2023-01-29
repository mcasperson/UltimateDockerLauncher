package manipulators

type ValueType int

type Manipulator interface {
	CanManipulate(fileSpec string) bool
	SetValue(fileSpec string, valueSpec string, value string) error
}
