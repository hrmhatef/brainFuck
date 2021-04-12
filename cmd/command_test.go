package cmd

import "testing"

func addOne(val uint16) uint16 {
	return val + 1
}

func power(val uint16) uint16 {
	return val * val
}

func TestCommand(t *testing.T) {
	one := NewCommand("c", Value, addOne)
	if !one.IsValue() {
		t.Error("the type of command should be value")
	}

	if one.Exec(2) != 3 {
		t.Error("the execution of func addOne is invalid")
	}

	if one.IsPointer() {
		t.Error("the type of command should not be value")
	}

	power := NewCommand("c", Pointer, power)
	if !power.IsPointer() {
		t.Error("the type of command should be pointer")
	}

	if power.Exec(2) != 4 {
		t.Error("the execution of func power is invalid")
	}
}
