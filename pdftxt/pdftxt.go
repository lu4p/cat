// Package pdftxt extracts text from .pdf documents
package pdftxt

import (
	"bytes"
	"io/ioutil"

	"github.com/lu4p/unipdf/v3/extractor"
	pdf "github.com/lu4p/unipdf/v3/model"
)

// ToStr converts a plaintext file to string
func ToStr(inputPath string) (string, error) {
	content, err := ioutil.ReadFile(inputPath)
	if err != nil {
		return "", err
	}
	return BytesToStr(content)
}

// BytesToStr converts a []byte representation of a .pdf document file to string
func BytesToStr(data []byte) (string, error) {
	reader := bytes.NewReader(data)

	pdfReader, err := pdf.NewPdfReader(reader)
	if err != nil {
		return "", err
	}

	numPages, err := pdfReader.GetNumPages()
	if err != nil {
		return "", err
	}

	var out string

	for i := 0; i < numPages; i++ {
		pageNum := i + 1

		page, err := pdfReader.GetPage(pageNum)
		if err != nil {
			return "", err
		}

		ex, err := extractor.New(page)
		if err != nil {
			return "", err
		}

		text, err := ex.ExtractText()
		if err != nil {
			return "", err
		}

		out += text + "\n\n"
	}

	return out, nil
}
