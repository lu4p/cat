package plaintxt

import (
	"errors"
	"io/ioutil"
)

// ToStr converts a plaintext file to string
func ToStr(filename string) (string, error) {
	outbyte, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", errors.New("Error while reading file")
	}
	return string(outbyte), nil
}
