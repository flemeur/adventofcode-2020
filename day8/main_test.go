package main

import (
	"strings"
	"testing"
)

func TestSolvePuzzle1(t *testing.T) {
	input := `
nop +0
acc +1
jmp +4
acc +3
jmp -3
acc -99
acc +1
jmp -4
acc +6
`
	program, err := readInput(strings.NewReader(input))
	if err != nil {
		t.Fatal(err)
	}

	answer := solvePuzzle1(program)
	if answer != 5 {
		t.Fatalf("expected 5, got %d", answer)
	}
}

func TestSolvePuzzle2(t *testing.T) {
	input := `
nop +0
acc +1
jmp +4
acc +3
jmp -3
acc -99
acc +1
jmp -4
acc +6
`
	program, err := readInput(strings.NewReader(input))
	if err != nil {
		t.Fatal(err)
	}

	answer := solvePuzzle2(program)
	if answer != 8 {
		t.Fatalf("expected 8, got %d", answer)
	}
}
