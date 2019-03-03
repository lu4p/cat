package odtxt

import (
	"testing"
)

const test = "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed id ex nec risus venenatis viverra. Cras condimentum dolor vitae dictum rutrum. Etiam viverra sit amet mi at lacinia.\n"

func TestDocx(t *testing.T) {
	txt, err := OdtTxt("../test/test.odt")
	if err != nil {
		t.Error(".odt failed:", err)
	} else if txt == test {
		t.Log(".odt success")
	} else {
		t.Error(".odt does not match test:", txt, test)
	}
}
