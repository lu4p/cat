package rtftxt

import (
	"os"
	"testing"
)

const txt = `Of course, we frequently hear about larger brands pushing out a ton of amazing content, and they're often used as examples of how to do content right. `
const red = `Restore The Selling Balance. Ad Technology doesn't have to be faceless. Our platform is designed to connect media companies directly to advertisers.`

func Test2Text(t *testing.T) {
	s, err := RtfTxt(`testdata/np.new.rtf`)
	if err != nil {
		t.Error(err)
	}
	if s != txt {
		t.Error("doesn't match", s)
	}
	s, err := RtfTxt(`testdata/ad.rtf`)
	if err != nil {
		t.Error(err)
	}
	if s != txt {
		t.Error("doesn't match", s)
	}
}
