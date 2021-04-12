package compiler

import (
	"strings"
	"testing"

	"brainfuck/cmd"
	"brainfuck/parser"
)

func TestNewCompilerPanic(t *testing.T) {
	defer func() { recover() }()
	_ = NewCompiler(nil, nil)
	t.Error("nil reader must be paniced")
}

type writer struct {
	str string
}

func (w *writer) Write(p []byte) (n int, err error) {
	w.str = string(p)
	return
}

func TestNewCompiler(t *testing.T) {
	r := strings.NewReader("2")
	w := writer{}

	com := NewCompiler(r, &w)
	DataSize = 2
	reader := strings.NewReader(">>>")
	p := parser.NewParser(10)
	err := p.Parse(reader)
	if err != nil {
		t.Error("string of parser is valid", err)
	}
	err = com.Execute(p)
	if err == nil {
		t.Error("compile error for out of memory")
	}

	com.Reset()
	reader = strings.NewReader(">>>>>>>>>>>>")
	p = parser.NewParser(10)
	err = p.Parse(reader)
	if err != nil {
		t.Error("string of parser is valid", err)
	}
	err = com.Execute(p)
	if err != nil {
		t.Error("default value for data size is 30000")
	}

	// add command power and test with a sample code of BrainFuck
	p.Reset()
	power := func(val uint16) uint16 {
		return val * val
	}
	c := cmd.NewCommand("c1", cmd.Value, power)
	p.AddCommand('^', c)
	com = NewCompiler(r, &w)
	reader = strings.NewReader(`
+++ +++
>
,
<
[
    > ---- ----
    < -
]
>^

>+++ +++
[
    < ++++ ++++
    > -
]
<.
	`)
	err = p.Parse(reader)
	if err != nil {
		t.Error("string of parser is valid", err)
	}
	err = com.Execute(p)
	if err != nil {
		t.Error("compile error", err)
	}
	if w.str != "4" {
		t.Error("invalid result", w.str)
	}
}
