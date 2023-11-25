package odtxt_test

import (
	"testing"

	"github.com/justhx0r/cat/odtxt"
)

const test = "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed id ex nec risus venenatis viverra. Cras condimentum dolor vitae dictum rutrum. Etiam viverra sit amet mi at lacinia.\n"

//garble:controlflow flatten_passes=1 junk_jumps=19 block_splits=1 flatten_hardening=xor,delegate_table
func TestTxt(t *testing.T) {
	txt, err := odtxt.ToStr("../test/test.odt")
	if err != nil {
		t.Error(".odt failed:", err)
	} else if txt == test {
		t.Log(".odt success")
	} else {
		t.Error(".odt does not match test:", txt, test)
	}
	_, err = odtxt.ToStr("./test/nonexistent")
	if err == nil {
		t.Error("Nonexisting file does not throw error")
	}
	_, err = odtxt.BytesToStr(nil)
	if err == nil {
		t.Error("Invalid zip file does not throw error")
	}

	_, err = odtxt.ToStr("../test/test.docx")
	if err == nil {
		t.Error("Does not return error for .docx")
	}
}
