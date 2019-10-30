package odtxt_test

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/lu4p/cat/odtxt"
)

const test = "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed id ex nec risus venenatis viverra. Cras condimentum dolor vitae dictum rutrum. Etiam viverra sit amet mi at lacinia.\n"

func TestTxt(t *testing.T) {
	txt, err := odtxt.ToStr("../test/test.odt")
	if err != nil {
		t.Error(".odt failed:", err)
	} else if txt == test {
		t.Log(".odt success")
	} else {
		t.Error(".odt does not match test:", txt, test)
	}
}

func TestBytesToTxt(t *testing.T) {
	file, err := os.Open("../test/test.odt")
	if err != nil {
		t.Error("can't open test.odt file")
	}
	data, err := ioutil.ReadAll(file)
	if err != nil {
		t.Error(err)
	}
	txt, err := odtxt.BytesToStr(data)
	if err != nil {
		t.Error(".odt failed:", err)
	} else if txt == test {
		t.Log(".odt success")
	} else {
		t.Error(".odt does not match test:", txt, test)
	}
}
