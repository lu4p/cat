<a href="https://unlicense.org/">![License](https://img.shields.io/github/license/lu4p/cat.svg)</a>
<a href="https://circleci.com/gh/lu4p/cat">![CircleCI (all branches)](https://img.shields.io/circleci/project/github/lu4p/cat.svg)
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
