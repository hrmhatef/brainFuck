package cmd

const (
	Pointer = CommandType(1)
	Value   = CommandType(2)
)

// CommandType denotes the int representation of the type of the Command that
// can be Poiner(changes the index) or Value(changes the value)
type CommandType int

// Commands is a map of char and Command struct
type Commands map[rune]Command

// Method denotes the function type that used in BrainFuck
type Method func(uint16) uint16

// Command denotes the command type in BrainFuck
type Command struct {
	name    string
	cmdType CommandType
	method  Method
}

// IsValue checks the type of the command is Value
func (c *Command) IsValue() bool {
	return c.cmdType == Value
}

// IsPointer checks the type of the command is Pointer
func (c *Command) IsPointer() bool {
	return c.cmdType == Pointer
}

// Exec executes the method of the command
func (c *Command) Exe(in uint16) uint16 {
	return c.method(in)
}

// NewCommand returns a new Command
func NewCommand(name string, ct CommandType, m Method) Command {
	return Command{
		name:    name,
		cmdType: ct,
		method:  m,
	}
}
