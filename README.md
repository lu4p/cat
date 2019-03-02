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
