package manipulators

type ValueType int

const (
	String  int = 0
	Boolean     = 1
	Number      = 2
)

type Manipulator interface {
	set(fileSpec string, valueSpec string, value string, valueType ValueType)
}
