package pdftxt_test

import (
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
	_, err = pdftxt.ToStr("../test/nonexistent")
	if err == nil {
		t.Error("Nonexisting file does not throw error")
	}
	_, err = pdftxt.ToStr("../test/test.docx")
	if err == nil {
		t.Error("Non .pdf file does not throw error")
	}
}
