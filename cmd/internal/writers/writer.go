package writers

type Writer interface {
	WriteString(file string, value string) error
}
