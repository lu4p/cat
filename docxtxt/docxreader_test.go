package docxtxt

import (
	"testing"
)

func TestDocx(t *testing.T) {
	string, err := DocxtoTxt("../test.docx")
	if err != nil {
		t.FailNow()
	}
	t.Log(string)
}
