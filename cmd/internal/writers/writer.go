package writers

type Writer interface {
	write(destination string, value string)
}
