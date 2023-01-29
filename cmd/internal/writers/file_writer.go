package writers

import "os"

type FileWriter struct {
}

func (w FileWriter) WriteString(file string, value string) error {
	f, err := os.Create(file)

	if err != nil {
		return err
	}

	defer f.Close()

	_, err = f.WriteString(value)

	return err
}
