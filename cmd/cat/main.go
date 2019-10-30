package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/micaelAlastor/cat"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		log.Fatalln("Please provide a filename")
	}
	path := strings.Join(args[1:], " ")
	log.Println(path)

	fmt.Println(cat.Cat(path))
}
