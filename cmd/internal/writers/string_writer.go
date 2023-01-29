package writers

// StringWriter is used for testing write operations
type StringWriter struct {
	Output map[string]string
}

func (w *StringWriter) WriteString(file string, value string) error {
	if w.Output == nil {
		w.Output = map[string]string{}
	}

	w.Output[file] = value
	return nil
}
