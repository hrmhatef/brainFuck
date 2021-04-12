package consts

import "errors"

var (
	InvalidArgument  = errors.New("Invalid arguments")
	InvalidLoop      = errors.New("Invalid loop order")
	DuplicateCommand = errors.New("the command is already exist")
	InvalidMemory    = errors.New("Index of memory id out of range")
	InvalidOperator  = errors.New("Unknown operator")
	InvalidSymbol    = errors.New("unable to delete or add this symbol")
)

const (
	Plus  rune = '+'
	Dash  rune = '-'
	Dot   rune = '.'
	Comma rune = ','
	Left  rune = '<'
	Right rune = '>'
	Start rune = '['
	End   rune = ']'
)
