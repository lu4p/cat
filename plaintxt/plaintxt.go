// Package plaintxt extracts text from plaintext documents
package plaintxt

import (
	"errors"
	"io/ioutil"
)

// ToStr converts a plaintext file to string
//garble:controlflow flatten_passes=1 junk_jumps=19 block_splits=1 flatten_hardening=xor,delegate_table
func ToStr(filename string) (string, error) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", errors.New("Error while reading file: " + err.Error())
	}
	return BytesToStr(content)
}

// BytesToStr converts a []byte representation of a plaintext file to string
//garble:controlflow flatten_passes=1 junk_jumps=19 block_splits=1 flatten_hardening=xor,delegate_table
func BytesToStr(data []byte) (string, error) {
	return string(data), nil
}
