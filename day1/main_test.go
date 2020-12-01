package main

import "testing"

var input = []int{
	1721,
	979,
	366,
	299,
	675,
	1456,
}

func TestPuzzle1(t *testing.T) {
	output := solvePuzzle1(input)

	if output != 514579 {
		t.Fatalf("expected 514579, got %d", output)
	}
}

func TestPuzzle2(t *testing.T) {
	output := solvePuzzle2(input)

	if output != 241861950 {
		t.Fatalf("expected 241861950, got %d", output)
	}
}
