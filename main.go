// main.go

package main

import (
	"errors"
	"log"
	"monkey/file"
	"monkey/repl"
	"os"
)

func main() {
	var err error
	args := os.Args[1:]
	switch len(args) {
	case 0:
		err = repl.Start()
	case 1:
		err = file.Read(args[0])

	default:
		err = errors.New("incorrect number of arguments")
	}

	if err != nil {
		log.Fatal(err)
	}
}
