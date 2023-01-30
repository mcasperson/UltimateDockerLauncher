package writers

type StringIOWriter struct {
	Output string
}

func (w *StringIOWriter) Write(p []byte) (n int, err error) {
	w.Output = string(p)
	return len(p), nil
}
