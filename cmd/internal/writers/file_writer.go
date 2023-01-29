package writers

import (
	"os"
	"path/filepath"
)

type FileWriter struct {
}

func (w FileWriter) WriteString(file string, value string) error {
	dir, _ := filepath.Split(file)

	if dir != "" {
		err := os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			return err
		}
	}

	f, err := os.Create(file)

	if err != nil {
		return err
	}

	defer f.Close()

	_, err = f.WriteString(value)

	return err
}
