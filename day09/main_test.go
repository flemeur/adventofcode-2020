package main

import (
	"strings"
	"testing"
)

func TestSolvePuzzle1(t *testing.T) {
	input := `
35
20
15
25
47
40
62
55
65
95
102
117
150
182
127
219
299
277
309
576
`

	numbers, err := readInput(strings.NewReader(input))
	if err != nil {
		t.Fatal(err)
	}

	answer := solvePuzzle1(numbers, 5)
	if answer != 127 {
		t.Fatalf("expected answer 127, got %d", answer)
	}
}

func TestSolvePuzzle2(t *testing.T) {
	input := `
35
20
15
25
47
40
62
55
65
95
102
117
150
182
127
219
299
277
309
576
`

	numbers, err := readInput(strings.NewReader(input))
	if err != nil {
		t.Fatal(err)
	}

	answer1 := solvePuzzle1(numbers, 5)

	answer := solvePuzzle2(numbers, answer1)
	if answer != 62 {
		t.Fatalf("expected answer 62, got %d", answer)
	}
}
