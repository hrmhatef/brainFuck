package bfError

import "errors"

var (
	InvalidArgument  = errors.New("Invalid arguments")
	InvalidLoop      = errors.New("Invalid loop order")
	DuplicateCommand = errors.New("the command is already exist")
	InvalidMemory    = errors.New("Index of memory id out of range")
	InvalidOperator  = errors.New("Unknown operator")
)
