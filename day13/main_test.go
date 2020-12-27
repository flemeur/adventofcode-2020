package main

import (
	"strings"
	"testing"
)

func TestSolvePuzzle1(t *testing.T) {
	input := `
939
7,13,x,x,59,x,31,19
`

	timestamp, busIDs, err := readInput(strings.NewReader(input))
	if err != nil {
		t.Fatal(err)
	}

	answer := solvePuzzle1(timestamp, busIDs)
	if answer != 295 {
		t.Fatalf("expected answer 295, got %d", answer)
	}
}

func TestSolvePuzzle2(t *testing.T) {
	cases := []struct {
		input  string
		answer int
	}{
		{`0
7,13,x,x,59,x,31,19`, 1068781},
		{`0
17,x,13,19`, 3417},
		{`0
67,7,59,61`, 754018},
		{`0
67,x,7,59,61`, 779210},
		{`0
67,7,x,59,61`, 1261476},
		{`0
1789,37,47,1889`, 1202161486},
	}

	for _, tc := range cases {
		t.Run("", func(t *testing.T) {
			_, busIDs, err := readInput(strings.NewReader(tc.input))
			if err != nil {
				t.Fatal(err)
			}

			answer := solvePuzzle2(busIDs)
			if answer != tc.answer {
				t.Fatalf("expected answer %d, got %d", tc.answer, answer)
			}
		})
	}
}
