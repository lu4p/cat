package rtftxt_test

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/micaelAlastor/cat/rtftxt"
)

const txt = `Of course, we frequently hear about larger brands pushing out a ton of amazing content, and they're often used as examples of how to do content right. `
const red = `Restore The Selling Balance. Ad Technology doesn't have to be faceless. Our platform is designed to connect media companies directly to advertisers.`

func Test2Text(t *testing.T) {
	s, err := rtftxt.ToStr(`../test/np.new.rtf`)
	if err != nil {
		t.Error(err)
	}
	if s != txt {
		t.Error("doesn't match", s)
	}
	s, err = rtftxt.ToStr(`../test/ad.rtf`)
	if err != nil {
		t.Error(err)
	}
	if s != red {
		t.Error("doesn't match", s)
	}
}

func TestBytes2Text(t *testing.T) {
	file, err := os.Open("../test/np.new.rtf")
	if err != nil {
		t.Error("can't open test.rtf file")
	}
	data, err := ioutil.ReadAll(file)
	s, err := rtftxt.BytesToStr(data)
	if err != nil {
		t.Error(err)
	}
	if s != txt {
		t.Error("doesn't match", s)
	}
	s, err = rtftxt.ToStr(`../test/ad.rtf`)
	if err != nil {
		t.Error(err)
	}
	if s != red {
		t.Error("doesn't match", s)
	}
}
