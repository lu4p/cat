package cat

import (
	"fmt"
	"testing"
)

const test = "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed id ex nec risus venenatis viverra. Cras condimentum dolor vitae dictum rutrum. Etiam viverra sit amet mi at lacinia.\n"

const red = "Restore The Selling Balance. Ad Technology doesn't have to be faceless. Our platform is designed to connect media companies directly to advertisers."

func TestCat(t *testing.T) {
	filetypes := []string{".docx", ".odt", ".txt"}
	for _, filetype := range filetypes {
		txt, err := Cat("./test/test" + filetype)
		if err != nil {
			t.Error(filetype, "failed:", err)
		} else if txt == test {
			t.Log(filetype, "success")
		} else {

			t.Error(filetype, "does not match test:", txt, test)
		}
	}
	rtf, err := Cat("./test/ad.rtf")
	if err != nil {
		t.Error(".rtf failed:", err)
	} else if rtf == red {
		t.Log(".rtf success")
	} else {
		t.Log(".rtf does not match test:", rtf, red)
	}
}

func Example() {
	txt, _ := Cat("./test/test.docx")
	fmt.Println(txt)
	// Output: Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed id ex nec risus venenatis viverra. Cras condimentum dolor vitae dictum rutrum. Etiam viverra sit amet mi at lacinia.
}
