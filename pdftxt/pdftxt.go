//Package pdftxt extracts text from pdf documents
package pdftxt

import (
	"bytes"
	"os"

	"github.com/lu4p/unipdf/v3/extractor"
	pdf "github.com/lu4p/unipdf/v3/model"
)

// BytesToStr prints out contents of a []byte representation of PDF file to stdout.
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

// ToStr prints out contents of PDF file to stdout.
func ToStr(inputPath string) (string, error) {
	f, err := os.Open(inputPath)
	if err != nil {
		return "", err
	}

	defer f.Close()

	pdfReader, err := pdf.NewPdfReader(f)
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
