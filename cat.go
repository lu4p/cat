// Package cat extracts text from multiple document formats
package cat

import (
	"errors"
	"os"
	"path/filepath"
	"strings"

	"github.com/lu4p/cat/docxtxt"
	"github.com/lu4p/cat/odtxt"
	"github.com/lu4p/cat/pdftxt"
	"github.com/lu4p/cat/plaintxt"
	"github.com/lu4p/cat/rtftxt"
)

// Cat reads a .odt, .docx, .rtf or plaintext file and returns the content as a string
func Cat(filename string) (txt string, err error) {
	filename, err = filepath.Abs(filename)
	if err != nil {
		return "", errors.New("abs " + err.Error())
	}
	_, err = os.Stat(filename)
	if err != nil {
		err = errors.New("File does not exist")
		return
	}
	if strings.HasSuffix(filename, ".odt") {
		txt, err = odtxt.ToStr(filename)
	} else if strings.HasSuffix(filename, ".docx") {
		txt, err = docxtxt.ToStr(filename)
	} else if strings.HasSuffix(filename, ".rtf") {
		txt, err = rtftxt.ToStr(filename)
	} else if strings.HasSuffix(filename, ".pdf") {
		txt, err = pdftxt.ToStr(filename)
	} else {
		txt, err = plaintxt.ToStr(filename)
	}
	return
}
