package main

import (
	"bufio"
	"fmt"
	"os"

	"brainfuck/cmd"
	"brainfuck/compiler"
	"brainfuck/parser"
)

type writer struct {
}

func (w writer) Write(p []byte) (n int, err error) {
	fmt.Print(string(p))
	return
}

func main() {
	args := os.Args
	if len(args) != 2 {
		fmt.Printf("Usage: %s filename\n", args[0])
	}
	filename := args[1]
	f, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Error reading %s\n", filename)
		return
	}
	re := bufio.NewReader(f)

	p := parser.NewParser(100)
	r := bufio.NewReader(os.Stdin)
	w := writer{}
	c := compiler.NewCompiler(r, w)
	power := cmd.NewCommand("power", cmd.Value, func(val uint16) uint16 { return val * val })
	p.AddCommand('p', power)
	enter := cmd.NewCommand("enter", cmd.Value, func(uint16) uint16 { return 10 })
	p.AddCommand('c', enter)
	doubleMove := cmd.NewCommand("double", cmd.Pointer, func(val uint16) uint16 { return val + 2 })
	p.AddCommand('d', doubleMove)

	err = p.Parse(re)
	if err != nil {
		fmt.Println("parser error:", err)
	}

	err = c.Execute(p)
	if err != nil {
		fmt.Println("exe error:", err)
	}
}
