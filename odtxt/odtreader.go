// Package odtxt extracts text from .odt documents
package odtxt

import (
	"archive/zip"
	"bytes"
	"encoding/xml"
	"errors"
	"io/ioutil"
)

//Odt zip struct
type Odt struct {
	zipFileReader *zip.ReadCloser
	zipReader     *zip.Reader
	Files         []*zip.File
	FilesContent  map[string][]byte
	Content       string
}

// BytesToStr converts a []byte representation of .odt document file to string
func BytesToStr(data []byte) (string, error) {
	reader := bytes.NewReader(data)
	d, err := OpenReader(reader)
	if err != nil {
		return "", err
	}
	content, err := d.GetTxt()
	if err != nil {
		return "", errors.New("Could not Get Content")
	}
	return content, nil
}

//OpenDocx open and load all files content
func OpenReader(bytesReader *bytes.Reader) (*Odt, error) {
	reader, err := zip.NewReader(bytesReader, bytesReader.Size())
	if err != nil {
		return nil, err
	}

	odtDoc := Odt{
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

// ToStr converts a .odt document file to string
func ToStr(filename string) (string, error) {
	d, err := Open(filename)
	if err != nil {
		return "", err
	}
	d.CloseZip()
	content, err := d.GetTxt()
	if err != nil {
		return "", errors.New("Could not Get Content")
	}
	return content, nil
}

func Open(path string) (*Odt, error) {
	reader, err := zip.OpenReader(path)
	if err != nil {
		return nil, err
	}

	odtDoc := Odt{
		zipFileReader: reader,
		Files:         reader.File,
		FilesContent:  map[string][]byte{},
	}

	for _, f := range odtDoc.Files {
		contents, _ := odtDoc.retrieveFileContents(f.Name)
		odtDoc.FilesContent[f.Name] = contents
	}

	return &odtDoc, nil
}

func (d *Odt) CloseZip() error {
	return d.zipFileReader.Close()
}

//Read all files contents
func (d *Odt) retrieveFileContents(filename string) ([]byte, error) {
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

func (d *Odt) GetTxt() (content string, err error) {
	xmlData := d.FilesContent["content.xml"]
	content, err = d.listP(xmlData)
	return
}

// listP for w:p tag value
func (d *Odt) listP(data []byte) (string, error) {
	v := new(Query)
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

type Query struct {
	XMLName xml.Name `xml:"document-content"`
	Body    Body     `xml:"body"`
}
type Body struct {
	Text []Text `xml:"text"`
}
type Text struct {
	P []string `xml:"p"`
}
