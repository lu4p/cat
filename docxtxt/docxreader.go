// Package docxtxt extracts text from .docx documents
package docxtxt

import (
	"archive/zip"
	"bytes"
	"errors"
	"io/ioutil"
	"regexp"
)

// docx zip struct
type docx struct {
	zipFileReader *zip.ReadCloser
	Files         []*zip.File
	FilesContent  map[string][]byte
	WordsList     []*words
}

type words struct {
	Content []string
}

// ToStr converts a .docx document file to string
func ToStr(filename string) (string, error) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return BytesToStr(content)
}

// BytesToStr converts a []byte representation of .docx document file to string
func BytesToStr(data []byte) (string, error) {
	reader := bytes.NewReader(data)
	d, err := openDocxReader(reader)
	if err != nil {
		return "", err
	}
	d.GenWordsList()
	var result string
	for _, word := range d.WordsList {
		for _, content := range word.Content {
			result += content
		}
		result += "\n"
	}
	return result, nil
}

// openDocxReader open and load all readers content
func openDocxReader(bytesReader *bytes.Reader) (*docx, error) {
	reader, err := zip.NewReader(bytesReader, bytesReader.Size())
	if err != nil {
		return nil, err
	}

	wordDoc := docx{
		zipFileReader: nil,
		Files:         reader.File,
		FilesContent:  map[string][]byte{},
	}

	for _, f := range wordDoc.Files {
		contents, _ := wordDoc.retrieveFileContents(f.Name)
		wordDoc.FilesContent[f.Name] = contents
	}

	return &wordDoc, nil
}

// Read all files contents
func (d *docx) retrieveFileContents(filename string) ([]byte, error) {
	var file *zip.File
	for _, f := range d.Files {
		if f.Name == filename {
			file = f
		}
	}

	if file == nil {
		return nil, errors.New(filename + " file not found")
	}

	reader, err := file.Open()
	if err != nil {
		return nil, err
	}

	return ioutil.ReadAll(reader)
}

// GenWordsList generate a list of all words
func (d *docx) GenWordsList() {
	xmlData := string(d.FilesContent["word/document.xml"])
	d.listP(xmlData)
}

// get w:t value
func (d *docx) getT(item string) {
	var subStr string
	data := item
	reRun := regexp.MustCompile(`(?U)(<w:r>|<w:r .*>)(.*)(</w:r>)`)
	re := regexp.MustCompile(`(?U)(<w:t>|<w:t .*>)(.*)(</w:t>)`)
	w := new(words)
	content := []string{}

	wrMatch := reRun.FindAllStringSubmatchIndex(data, -1)
	// loop r
	for _, rMatch := range wrMatch {
		rData := data[rMatch[4]:rMatch[5]]
		wtMatch := re.FindAllStringSubmatchIndex(rData, -1)
		for _, match := range wtMatch {
			subStr = rData[match[4]:match[5]]
			content = append(content, subStr)
		}
	}
	w.Content = content
	d.WordsList = append(d.WordsList, w)
}

// hasP identify the paragraph
func hasP(data string) bool {
	re := regexp.MustCompile(`(?U)<w:p .*>(.*)</w:p>`)
	result := re.MatchString(data)
	return result
}

// listP for w:p tag value
func (d *docx) listP(data string) {
	var result []string
	re := regexp.MustCompile(`(?U)<w:p>(.*)</w:p>`)
	for _, match := range re.FindAllStringSubmatch(data, -1) {
		result = append(result, match[1])
	}
	for _, item := range result {
		if hasP(item) {
			d.listP(item)
			continue
		}
		d.getT(item)
	}
}
