package docxtxt_test

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/micaelAlastor/cat/docxtxt"
)

const test = "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed id ex nec risus venenatis viverra. Cras condimentum dolor vitae dictum rutrum. Etiam viverra sit amet mi at lacinia.\n"

func TestToStr(t *testing.T) {
	txt, err := docxtxt.ToStr("../test/test.docx")
	if err != nil {
		t.Error(".docx failed:", err)
	} else if txt == test {
		t.Log(".docx success")
	} else {
		t.Error(".docx does not match test:", txt, test)
	}
}

func TestBytesToTxt(t *testing.T) {
	file, err := os.Open("../test/test.docx")
	if err != nil {
		t.Error("can't open test.docx file")
	}
	data, err := ioutil.ReadAll(file)
	if err != nil {
		t.Error(err)
	}
	txt, err := docxtxt.BytesToStr(data)
	if err != nil {
		t.Error(".docx failed:", err)
	} else if txt == test {
		t.Log(".docx success")
	} else {
		t.Error(".docx does not match test:", txt, test)
	}
}
