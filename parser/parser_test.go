package parser

import (
	"brainfuck/cmd"
	"testing"
)

func TestDefaultParser(t *testing.T) {
	parser := NewParser(10)

	c, ok := parser.Command('+')
	if !ok {
		t.Error("unable to find default command 'Inc' in parser")
	}
	if c.Exe(1) != 2 {
		t.Error("invalid result for command 'Inc'")
	}
	if !c.IsValue() {
		t.Error("invalid type for command 'Inc'")
	}

	c, ok = parser.Command('-')
	if !ok {
		t.Error("unable to find default command 'Dec' in parser")
	}
	if c.Exe(1) != 0 {
		t.Error("invalid result for command 'Dec'")
	}
	if !c.IsValue() {
		t.Error("invalid type for command 'Dec'")
	}

	c, ok = parser.Command('>')
	if !ok {
		t.Error("unable to find default command 'Right' in parser")
	}
	if c.Exe(1) != 2 {
		t.Error("invalid result for command 'Right'")
	}
	if !c.IsPointer() {
		t.Error("invalid type for command 'Right'")
	}

	c, ok = parser.Command('<')
	if !ok {
		t.Error("unable to find default command 'Left' in parser")
	}
	if c.Exe(1) != 0 {
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
