package pdftxt_test

import (
	"io/ioutil"
	"os"
	"strings"
	"testing"

	"github.com/lu4p/cat/pdftxt"
)

const test = `Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed id ex nec risus venenatis viverra. Cras 
condimentum dolor vitae dictum rutrum. Etiam viverra sit amet mi at lacinia.`

func TestToStr(t *testing.T) {
	txt, err := pdftxt.ToStr("../test/test.pdf")
	txt = strings.TrimSpace(txt)
	if err != nil {
		t.Error(".pdf failed:", err)
	} else if txt == test {
		t.Log(".pdf success")
	} else {
		t.Error(".pdf does not match test:", txt, test)
	}
}

func TestBytesToStr(t *testing.T) {
	file, err := os.Open("../test/test.pdf")
	if err != nil {
		t.Error("can't open test.pdf file")
	}
	data, err := ioutil.ReadAll(file)
	txt, err := pdftxt.BytesToStr(data)
	txt = strings.TrimSpace(txt)
	if err != nil {
		t.Error(".pdf failed:", err)
	} else if txt == test {
		t.Log(".pdf success")
	} else {
		t.Error(".pdf does not match test:", txt, test)
	}
}
