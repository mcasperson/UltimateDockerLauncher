package readers

import "io/ioutil"

type FileReader struct {
}

func (r FileReader) ReadString(file string) (string, error) {
	bytes, err := ioutil.ReadFile(file)

	if err != nil {
		return "", err
	}

	return string(bytes), nil
}
