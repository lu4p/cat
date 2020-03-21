// Package odtxt extracts text from .odt documents
package odtxt

import (
	"archive/zip"
	"bytes"
	"encoding/xml"
	"errors"
	"io/ioutil"
)

// odt zip struct
type odt struct {
	zipFileReader *zip.ReadCloser
	Files         []*zip.File
	FilesContent  map[string][]byte
	Content       string
}

type query struct {
	XMLName xml.Name `xml:"document-content"`
	Body    body     `xml:"body"`
}
type body struct {
	Text []text `xml:"text"`
}
type text struct {
	P []string `xml:"p"`
}

// ToStr converts a .odt document file to string
func ToStr(filename string) (string, error) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return BytesToStr(content)
}

// BytesToStr converts a []byte representation of a .odt document file to string
func BytesToStr(data []byte) (string, error) {
	reader := bytes.NewReader(data)
	d, err := openReader(reader)
	if err != nil {
		return "", err
	}
	content, err := d.GetTxt()
	if err != nil {
		return "", errors.New("could not Get Content")
	}
	return content, nil
}

// openReader open and load all readers content
func openReader(bytesReader *bytes.Reader) (*odt, error) {
	reader, err := zip.NewReader(bytesReader, bytesReader.Size())
	if err != nil {
		return nil, err
	}

	odtDoc := odt{
		zipFileReader: nil,
		Files:         reader.File,
		FilesContent:  map[string][]byte{},
	}

	for _, f := range odtDoc.Files {
		contents, _ := odtDoc.retrieveFileContents(f.Name)
		odtDoc.FilesContent[f.Name] = contents
	}

	return &odtDoc, nil
}

// retrieveFileContents Read all files contents
func (d *odt) retrieveFileContents(filename string) ([]byte, error) {
	var file *zip.File
	for _, f := range d.Files {
		if f.Name == filename {
			file = f
			break
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

func (d *odt) GetTxt() (string, error) {
	xmlData := d.FilesContent["content.xml"]
	return d.listP(xmlData)
}

// listP for w:p tag value
func (d *odt) listP(data []byte) (string, error) {
	v := new(query)
	err := xml.Unmarshal(data, &v)
	if err != nil {
		return "", err
	}
	var result string
	for _, text := range v.Body.Text {
		for _, P := range text.P {
			for _, item := range P {
				result += string(item)
			}
			result += "\n"
		}
	}
	return result, nil
}
