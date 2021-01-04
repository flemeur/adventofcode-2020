package main

import (
	"strings"
	"testing"
)

func TestSolvePuzzle1(t *testing.T) {
	input := `
.#.
..#
###
`

	pd, err := readInput(strings.NewReader(input))
	if err != nil {
		t.Fatal(err)
	}

	answer := solvePuzzle1(pd)

	if answer != 112 {
		t.Fatalf("expected answer 112, got %d", answer)
	}
}

func TestPocketDimension(t *testing.T) {
	input := `
.#.
..#
###
`

	pd, err := readInput(strings.NewReader(input))
	if err != nil {
		t.Fatal(err)
	}

	nActive := pd.CountActive()
	if nActive != 5 {
		t.Fatalf("expected 5 active points, got %d", nActive)
	}

	pd = pd.Evolve()

	nActive = pd.CountActive()
	if nActive != 11 {
		t.Fatalf("expected 11 active points, got %d", nActive)
	}
}

func TestSolvePuzzle2(t *testing.T) {
	input := `
.#.
..#
###
`

	pd, err := readInput(strings.NewReader(input))
	if err != nil {
		t.Fatal(err)
	}

	answer := solvePuzzle2(pd)

	if answer != 848 {
		t.Fatalf("expected answer 848, got %d", answer)
	}
}
