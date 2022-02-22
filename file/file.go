package file

import (
	"errors"
	"fmt"
	"io/ioutil"
	"monkey/evaluator"
	"monkey/lexer"
	"monkey/object"
	"monkey/parser"
	"strings"
)

// Read will read a .mky file
func Read(filename string) error {
	// Verify filename contains .mky
	if !strings.Contains(filename, ".mky") {
		return errors.New("not a .mky program")
	}
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	// Convert []byte to string and print to screen
	input := string(content)
	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()
	env := object.NewEnvironment()
	e := evaluator.Eval(program, env)
	fmt.Printf("%v\n", e.Inspect())

	return nil
}
