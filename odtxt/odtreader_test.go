package odtxt

import (
	"testing"
)

func TestRead(t *testing.T) {
	string, err := OdtTxt("../test.odt")
	if err != nil {
		t.FailNow()
	}
	t.Log(string)
}
