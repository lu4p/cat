package cat_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/lu4p/cat"
)

const test = "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed id ex nec risus venenatis viverra. Cras condimentum dolor vitae dictum rutrum. Etiam viverra sit amet mi at lacinia."

const red = "Restore The Selling Balance. Ad Technology doesn't have to be faceless. Our platform is designed to connect media companies directly to advertisers."

func TestCat(t *testing.T) {
	filetypes := []string{".docx", ".odt", ".txt"}
	for _, filetype := range filetypes {
		txt, err := cat.File("./test/test" + filetype)
		if err != nil {
			t.Error(filetype, "failed:", err)
		}
		txt = strings.TrimSpace(txt)
		txt = strings.ReplaceAll(txt, "\n", "")
		if txt == test {
			t.Log(filetype, "success")
		} else {
			t.Error(filetype, "does not match test: got:", txt, "expected:", test)
		}
	}
	rtf, err := cat.File("./test/ad.rtf")
	if err != nil {
		t.Error(".rtf failed:", err)
	} else if rtf == red {
		t.Log(".rtf success")
	} else {
		t.Log(".rtf does not match test:", rtf, red)
	}

	_, err = cat.File("./test/nonexistent")
	if err == nil {
		t.Error("Nonexisting file does not throw error")
	}
}

func Example() {
	txt, _ := cat.File("./test/test.docx")
	fmt.Println(txt)
	// Output: Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed id ex nec risus venenatis viverra. Cras condimentum dolor vitae dictum rutrum. Etiam viverra sit amet mi at lacinia.
}
