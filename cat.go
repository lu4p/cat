package cat

import (
	"errors"
	"io/ioutil"
	"os"
	"strings"

	"github.com/lu4p/cat/docxtxt"
	"github.com/lu4p/cat/odtxt"
	"github.com/lu4p/cat/rtftxt"
)

func Cat(filename string) (txt string, err error) {
	_, err = os.Stat(filename)
	if err != nil {
		err = errors.New("[!] File does not exist!")
		return
	}
	if strings.HasSuffix(filename, ".odt") {
		txt, err = odtxt.OdtTxt(filename)
	} else if strings.HasSuffix(filename, ".docx") {
		txt, err = docxtxt.DocxtoTxt(filename)
	} else if strings.HasSuffix(filename, ".rtf") {
		txt, err = rtftxt.RtfTxt(filename)
	} else {
		txt, err = readTxt(filename)
	}
	return
}

func readTxt(filename string) (string, error) {
	outbyte, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", errors.New("[!] Error while reading file")
	}
	return string(outbyte), nil
}
