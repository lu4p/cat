package cat

import (
	"testing"
)

func TestCat(t *testing.T)  {
	_, err := Cat("test.docx")
	if err != nil {
		t.Error(".docx failed:", err)
	}else {
		t.Log(".docx success")
	}

	_, err = Cat("test.odt")
	if err != nil {
		t.Error(".odt failed:", err)
	} else {
		t.Log(".odt success")
	}
	_, err = Cat("test.txt")
	if err != nil {
		t.Error(".txt failed:", err)
	}else {
		t.Log(".txt success")
	}
	_, err = Cat("./rtftxt/testdata/ad.rtf")
	if err != nil {
		t.Error(".rtf failed:", err)
	}else {
		t.Log(".rtf success")
	}
}
