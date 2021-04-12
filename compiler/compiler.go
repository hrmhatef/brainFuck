package compiler

import (
	"fmt"
	"io"

	"brainfuck/consts"
	"brainfuck/parser"
)

const maxDataSize = 30000

// DataSize denotes the memory size of compiler, it uses for memory limitation,
// zero means maxDataSize limit
var DataSize uint16 = 0

// Compiler denotes the compiler type in BrainFuck
type Compiler struct {
	reader         io.Reader
	writer         io.Writer
	executeCounter uint16
	data           []int16
	dataPtr        uint16
}

// NewCompiler returns a new compiler based on dataSize
func NewCompiler(r io.Reader, w io.Writer) *Compiler {
	if r == nil || w == nil {
		panic(consts.InvalidArgument)
	}

	var c Compiler
	c.data = make([]int16, maxDataSize)
	c.executeCounter = 0
	c.dataPtr = 0
	c.reader = r
	c.writer = w

	return &c
}

// Reset restes all data in compiler
func (c *Compiler) Reset() {
	c.executeCounter = 0
	c.dataPtr = 0
	DataSize = 0
}

// Execute runs the program of parser
func (c *Compiler) Execute(parser *parser.Parser) (err error) {
	p := parser.Program()
	for c.executeCounter < uint16(len(p)) {
		if c.dataPtr >= DataSize && (DataSize > 0 && DataSize <= maxDataSize) {
			return consts.InvalidMemory
		} else if c.dataPtr >= maxDataSize {
			return consts.InvalidMemory
		}

		switch p[c.executeCounter].Operator {
		case consts.Dot:
			err = c.output(byte(c.data[c.dataPtr]))
			if err != nil {
				return err
			}
		case consts.Comma:
			d, err := c.input()
			if err != nil && err != io.EOF {
				return err
			}
			c.data[c.dataPtr] = d
		case consts.Start:
			if c.data[c.dataPtr] == 0 {
				c.executeCounter = p[c.executeCounter].Operand
			}
		case consts.End:
			// make sure to ignore empty loop, like ++[]
			if c.data[c.dataPtr] > 0 && c.executeCounter-p[c.executeCounter].Operand > 1 {
				c.executeCounter = p[c.executeCounter].Operand
			}
		default:
			s := p[c.executeCounter].Operator
			cmd, ok := parser.Command(s)
			if ok {
				if cmd.IsValue() {
					c.data[c.dataPtr] = int16(cmd.Exec(uint16(c.data[c.dataPtr])))
				} else if cmd.IsPointer() {
					c.dataPtr = cmd.Exec(c.dataPtr)
				}
			} else {
				panic(fmt.Sprintf("%s {%s}", consts.InvalidOperator, string(s)))
			}
		}
		c.executeCounter++
	}

	return
}

func (c *Compiler) output(out byte) error {
	var b []byte
	b = append(b, out)
	_, err := c.writer.Write(b)
	if err != nil {
		return err
	}

	return nil
}

func (c *Compiler) input() (readVal int16, err error) {
	buf := make([]byte, 1)
	var n int
	for {
		n, err = c.reader.Read(buf)
		if err == io.EOF {
			break
		}

		if err != nil {
			break
		}

		if n > 0 {
			readVal = int16(buf[0])
			break
		}
	}

	return
}
