package docxtxt_test

import (
	"testing"

	"github.com/lu4p/cat/docxtxt"
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
