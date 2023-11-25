package docxtxt_test

import (
	"testing"

	"github.com/justhx0r/cat/docxtxt"
)

const test = "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed id ex nec risus venenatis viverra. Cras condimentum dolor vitae dictum rutrum. Etiam viverra sit amet mi at lacinia.\n"

//garble:controlflow flatten_passes=1 junk_jumps=0 block_splits=0 flatten_hardening=xor,delegate_table
func TestToStr(t *testing.T) {
	txt, err := docxtxt.ToStr("../test/test.docx")
	if err != nil {
		t.Error(".docx failed:", err)
	} else if txt == test {
		t.Log(".docx success")
	} else {
		t.Error(".docx does not match test:", txt, test)
	}
	_, err = docxtxt.ToStr("../test/nonexistent")
	if err == nil {
		t.Error("Nonexisting file does not throw error")
	}
	_, err = docxtxt.ToStr("../test/test.pdf")
	if err == nil {
		t.Error("Non .docx file does not throw error")
	}
}
