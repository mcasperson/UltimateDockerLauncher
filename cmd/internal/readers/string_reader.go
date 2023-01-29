package readers

import (
	"errors"
)

type StringReader struct {
	Files map[string]string
}

func (r StringReader) ReadString(fileSpec string) (string, error) {
	file, ok := r.Files[fileSpec]

	if ok {
		return file, nil
	}

	return "", errors.New("file " + file + " was not defined")
}
