[![GoDoc](https://godoc.org/github.com/lu4p/cat?status.svg)](https://godoc.org/github.com/lu4p/cat)
[![License](https://img.shields.io/github/license/lu4p/cat.svg)](https://unlicense.org/)
[![CircleCI](https://circleci.com/gh/lu4p/cat.svg?style=svg)](https://circleci.com/gh/lu4p/cat)
[![Go Report Card](https://goreportcard.com/badge/github.com/lu4p/cat)](https://goreportcard.com/report/github.com/lu4p/cat)
# cat
This is a simple libary to extract text from plaintext, .docx, .odt and .rtf files.

Work in progess...

## Install
```go get -u github.com/lu4p/cat```

## Usage
```golang 
package main
import (
  "fmt"
  "github.com/lu4p/cat"
)

func main(){
  txt, _ := cat.Cat("filename")
  fmt.Println(txt)
}
```
