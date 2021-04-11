package parser

import (
	"errors"
	"fmt"
	"io"

	"brainfuck/cmd"
)

// Instruction denotes each command into Parser.Instruction
type Instruction struct {
	Operator rune
	Operand  uint16
}

// Parser denotes the parser type in BrainFuck
type Parser struct {
	cmds        cmd.Commands
	readCounter uint16
	stackLoop   []uint16
	program     []Instruction
	bufSize     int
}

// NewParser returns a new parser with defualt commands
func NewParser(bufSize int) *Parser {
	var p Parser
	p.bufSize = bufSize
	p.stackLoop = make([]uint16, 0)
	p.program = make([]Instruction, 0)
	p.cmds = makeDefaultCommands()
	p.readCounter = 0

	return &p
}

// Parse tries to convert each command into its instruction
func (p *Parser) Parse(reader io.Reader) (err error) {
	if reader == nil {
		panic("Bad initialize...")
	}

	for {
		var input string
		input, err = readString(reader, p.bufSize)
		if err == nil {
			for _, char := range input {
				err = p.parse(char)
				if err != nil {
					return
				}
			}
		} else {
			if err == io.EOF && len(p.stackLoop) == 0 {
				return nil
			} else {
				return errors.New(fmt.Sprintf("Invalid loop order %s", err))
			}
		}
	}
}

// Program returns all instructions
func (p *Parser) Program() []Instruction {
	return p.program
}

// Command returns the command from available commands
func (p *Parser) Command(s rune) (c cmd.Command, ok bool) {
	c, ok = p.cmds[s]
	return
}

// AddCommand inserts a new command into command list
func (p *Parser) AddCommand(symbol rune, cmd cmd.Command) (err error) {
	_, ok := p.cmds[symbol]
	if ok {
		return errors.New("commad is already exist")
	}
	p.cmds[symbol] = cmd
	return
}

// RemoveCommand removes a command with its symbol
func (p *Parser) RemoveCommand(symbol rune) {
	delete(p.cmds, symbol)
}

func (p *Parser) startLoop(symbol rune) {
	p.program = append(p.program, Instruction{symbol, 0})
	p.stackLoop = append(p.stackLoop, p.readCounter)
}

func (p *Parser) endLoop(symbol rune) error {
	if len(p.stackLoop) == 0 {
		return errors.New("Invalid loop char")
	}

	jump := p.stackLoop[len(p.stackLoop)-1]
	p.stackLoop = p.stackLoop[:len(p.stackLoop)-1]
	p.program = append(p.program, Instruction{symbol, jump})
	p.program[jump].Operand = p.readCounter

	return nil
}

func readString(reader io.Reader, bf int) (str string, err error) {
	buf := make([]byte, bf)
	var n int
	for {
		n, err = reader.Read(buf)
		if err == io.EOF {
			break
		}

		if err != nil {
			break
		}

		if n > 0 {
			return string(buf[:n]), nil
		}
	}

	return
}

func (p *Parser) parse(symbol rune) error {
	switch symbol {
	case '.', ',':
		p.program = append(p.program, Instruction{symbol, 0})
	case '[':
		p.startLoop(symbol)
	case ']':
		err := p.endLoop(symbol)
		if err != nil {
			return err
		}
	default:
		_, ok := p.cmds[symbol]
		if ok {
			p.program = append(p.program, Instruction{symbol, 0})
		} else {
			p.readCounter--
		}
	}
	p.readCounter++

	return nil
}

func makeDefaultCommands() map[rune]cmd.Command {
	cmds := make(map[rune]cmd.Command, 4)
	cmds['+'] = cmd.NewCommand("Inc", cmd.Value, func(val uint16) uint16 { return val + 1 })
	cmds['-'] = cmd.NewCommand("Dec", cmd.Value, func(val uint16) uint16 { return val - 1 })
	cmds['>'] = cmd.NewCommand("Right", cmd.Pointer, func(val uint16) uint16 { return val + 1 })
	cmds['<'] = cmd.NewCommand("Left", cmd.Pointer, func(val uint16) uint16 { return val - 1 })

	return cmds
}
