package parser

import (
	"strings"
	"testing"

	"brainfuck/cmd"
)

func TestDefaultParser(t *testing.T) {
	parser := NewParser(10)

	c, ok := parser.Command('+')
	if !ok {
		t.Error("unable to find default command 'Inc' in parser")
	}
	if c.Exec(1) != 2 {
		t.Error("invalid result for command 'Inc'")
	}
	if !c.IsValue() {
		t.Error("invalid type for command 'Inc'")
	}

	c, ok = parser.Command('-')
	if !ok {
		t.Error("unable to find default command 'Dec' in parser")
	}
	if c.Exec(1) != 0 {
		t.Error("invalid result for command 'Dec'")
	}
	if !c.IsValue() {
		t.Error("invalid type for command 'Dec'")
	}

	c, ok = parser.Command('>')
	if !ok {
		t.Error("unable to find default command 'Right' in parser")
	}
	if c.Exec(1) != 2 {
		t.Error("invalid result for command 'Right'")
	}
	if !c.IsPointer() {
		t.Error("invalid type for command 'Right'")
	}

	c, ok = parser.Command('<')
	if !ok {
		t.Error("unable to find default command 'Left' in parser")
	}
	if c.Exec(1) != 0 {
		t.Error("invalid result for command 'Left'")
	}
	if !c.IsPointer() {
		t.Error("invalid type for command 'Left'")
	}

	p := parser.Program()
	if len(p) != 0 {
		t.Error("invalid default value for array program")
	}
}

func TestAddRemoveCommand(t *testing.T) {
	parser := NewParser(10)

	method := func(uint16) uint16 {
		return 1
	}

	c := cmd.NewCommand("c", cmd.Value, method)
	err := parser.AddCommand('+', c)
	if err == nil {
		t.Error("there is no way to add duplicate command")
	}

	err = parser.AddCommand('t', c)
	if err != nil {
		t.Error("the valid command should be inserted in parser")
	}

	parser.RemoveCommand('t')
	_, ok := parser.Command('t')
	if ok {
		t.Error("deleted command should be remove from parser")
	}
}

func TestPanic(t *testing.T) {
	parser := NewParser(10)
	defer func() { recover() }()
	parser.Parse(nil)
	t.Error("nil reader must be paniced")
}

func TestParser(t *testing.T) {
	reader := strings.NewReader("+++>[[+-]")
	parser := NewParser(10)
	err := parser.Parse(reader)
	if err == nil {
		t.Error("invalid loop in unacceptable")
	}

	parser.Reset()
	reader = strings.NewReader("")
	err = parser.Parse(reader)
	if err != nil {
		t.Error("empty string is not invalid", err)
	}

	parser.Reset()
	reader = strings.NewReader("s;ldkfjsd;lkfj")
	err = parser.Parse(reader)
	if err != nil {
		t.Error("there is no command", err)
	}

	parser.Reset()
	reader = strings.NewReader("s;ldkfjsd;lkfj]")
	err = parser.Parse(reader)
	if err == nil {
		t.Error("there is an invlaid loop char")
	}

	parser.Reset()
	reader = strings.NewReader("++[]")
	err = parser.Parse(reader)
	if err != nil {
		t.Error("string is valid", err)
	}

	parser.Reset()
	reader = strings.NewReader(">>[[[-]]]")
	err = parser.Parse(reader)
	if err != nil {
		t.Error("neasted loop is valid", err)
	}
}
