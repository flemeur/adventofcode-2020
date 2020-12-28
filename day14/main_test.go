package main

import (
	"strings"
	"testing"
)

func TestSolvePuzzle1(t *testing.T) {
	input := `
mask = XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X
mem[8] = 11
mem[7] = 101
mem[8] = 0
`

	lines, err := readInput(strings.NewReader(input))
	if err != nil {
		t.Fatal(err)
	}

	answer, err := solvePuzzle1(lines)
	if err != nil {
		t.Fatal(err)
	}

	if answer != 165 {
		t.Fatalf("expected answer 165, got %d", answer)
	}
}

func TestSolvePuzzle2(t *testing.T) {
	input := `
mask = 000000000000000000000000000000X1001X
mem[42] = 100
mask = 00000000000000000000000000000000X0XX
mem[26] = 1
`

	lines, err := readInput(strings.NewReader(input))
	if err != nil {
		t.Fatal(err)
	}

	answer, err := solvePuzzle2(lines)
	if err != nil {
		t.Fatal(err)
	}

	if answer != 208 {
		t.Fatalf("expected answer 208, got %d", answer)
	}
}

func TestGenerateBinaryString(t *testing.T) {
	binStrings := generateBinaryStrings("1X0XX")

	t.Logf("%+v", binStrings)

	if len(binStrings) != 8 {
		t.Fatalf("expected 8 strings, got %d", len(binStrings))
	}
}
