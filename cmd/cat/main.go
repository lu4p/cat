package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/justhx0r/cat"
)

//garble:controlflow flatten_passes=max junk_jumps=max block_splits=max flatten_hardening=xor,delegate_table
func main() {
	args := os.Args
	if len(args) < 2 {
		log.Fatalln("Please provide a filename")
	}
	path := strings.Join(args[1:], " ")
	log.Println(path)

	fmt.Println(cat.File(path))
}
