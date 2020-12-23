package main

import (
	"strings"
	"testing"
)

func TestReadInput(t *testing.T) {
	input := `abc

a
b
c

ab
ac

a
a
a
a

b`

	groups, err := readInput(strings.NewReader(input))
	if err != nil {
		t.Fatal(err)
	}

	if len(groups) != 5 {
		t.Fatalf("expected 5 groups, got %d", len(groups))
	}

	if len(groups[0]) != 1 {
		t.Fatalf("expected len(groups[0]) = 1, got %d", len(groups[0]))
	}
	if len(groups[1]) != 3 {
		t.Fatalf("expected len(groups[1]) = 3, got %d", len(groups[1]))
	}
}

func TestSolvePuzzle1(t *testing.T) {
	input := `abc

a
b
c

ab
ac

a
a
a
a

b`

	groups, err := readInput(strings.NewReader(input))
	if err != nil {
		t.Fatal(err)
	}

	answer := solvePuzzle1(groups)
	if answer != 11 {
		t.Fatalf("expected answer 11, got %d", answer)
	}
}

func TestSolvePuzzle2(t *testing.T) {
	input := `abc

a
b
c

ab
ac

a
a
a
a

b`

	groups, err := readInput(strings.NewReader(input))
	if err != nil {
		t.Fatal(err)
	}

	answer := solvePuzzle2(groups)
	if answer != 6 {
		t.Fatalf("expected answer 6, got %d", answer)
	}
}
