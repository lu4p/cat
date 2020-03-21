[![GoDoc](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white&style=flat-square)](https://pkg.go.dev/github.com/lu4p/cat)
[![License](https://img.shields.io/github/license/lu4p/cat.svg)](https://unlicense.org/)
[![CircleCI](https://circleci.com/gh/lu4p/cat.svg?style=svg)](https://circleci.com/gh/lu4p/cat)
[![Go Report Card](https://goreportcard.com/badge/github.com/lu4p/cat)](https://goreportcard.com/report/github.com/lu4p/cat)
[![codecov](https://codecov.io/gh/lu4p/cat/branch/master/graph/badge.svg)](https://codecov.io/gh/lu4p/cat)

# cat
This is a simple libary to extract text from plaintext, .docx, .odt, .pdf and .rtf files.

## Install
```go get -u github.com/lu4p/cat```

## Basic Usage
```golang 
package main
import (
  "fmt"
  "github.com/lu4p/cat"
)

func main(){
  txt, _ := cat.File("filename")
  fmt.Println(txt)
}
```
