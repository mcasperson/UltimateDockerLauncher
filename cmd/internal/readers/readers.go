package readers

type Reader interface {
	ReadString(file string) (string, error)
}
