package main

import (
	"strings"
	"testing"
)

func TestSolvePuzzle1(t *testing.T) {
	cases := []struct {
		input  string
		answer int
	}{
		{"0,3,6", 436},
		{"1,3,2", 1},
		{"2,1,3", 10},
		{"1,2,3", 27},
		{"2,3,1", 78},
		{"3,2,1", 438},
		{"3,1,2", 1836},
	}

	for _, tc := range cases {
		t.Run("", func(t *testing.T) {
			startNumbers, err := readInput(strings.NewReader(tc.input))
			if err != nil {
				t.Fatal(err)
			}

			answer := solvePuzzle1(startNumbers)
			if answer != tc.answer {
				t.Fatalf("expected answer %d, got %d", tc.answer, answer)
			}
		})
	}
}

func TestSolvePuzzle2(t *testing.T) {
	cases := []struct {
		input  string
		answer int
	}{
		{"0,3,6", 175594},
		{"1,3,2", 2578},
		{"2,1,3", 3544142},
		{"1,2,3", 261214},
		{"2,3,1", 6895259},
		{"3,2,1", 18},
		{"3,1,2", 362},
	}

	for _, tc := range cases {
		t.Run("", func(t *testing.T) {
			startNumbers, err := readInput(strings.NewReader(tc.input))
			if err != nil {
				t.Fatal(err)
			}

			answer := solvePuzzle2(startNumbers)
			if answer != tc.answer {
				t.Fatalf("expected answer %d, got %d", tc.answer, answer)
			}
		})
	}
}
