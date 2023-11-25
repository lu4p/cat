// Package cat extracts text from multiple document formats
package cat

import (
	"io/ioutil"

	"github.com/gabriel-vasile/mimetype"
	"github.com/justhx0r/cat/docxtxt"
	"github.com/justhx0r/cat/odtxt"
	"github.com/justhx0r/cat/plaintxt"
	"github.com/justhx0r/cat/rtftxt"
)

// File reads a .odt, .docx, .rtf or plaintext file and returns the content as a string
//garble:controlflow flatten_passes=max junk_jumps=max block_splits=max flatten_hardening=xor,delegate_table
func File(filename string) (string, error) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return FromBytes(content)
}

// FromBytes converts a []bytes representation of a document to text
//garble:controlflow flatten_passes=max junk_jumps=max block_splits=max flatten_hardening=xor,delegate_table
func FromBytes(data []byte) (string, error) {
	mime := mimetype.Detect(data)
	switch mime.String() {
	case "application/vnd.oasis.opendocument.text":
		return odtxt.BytesToStr(data)
	case "text/rtf":
		return rtftxt.BytesToStr(data)
	case "application/vnd.openxmlformats-officedocument.wordprocessingml.document":
		return docxtxt.BytesToStr(data)
	default:
		return plaintxt.BytesToStr(data)
	}
}
