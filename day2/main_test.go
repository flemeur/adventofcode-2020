package main

import (
	"bytes"
	"testing"
)

var inputTxt = `1-3 a: abcde
1-3 b: cdefg
2-9 c: ccccccccc`

func TestSolvePuzzle1(t *testing.T) {
	input, err := readInput(bytes.NewReader([]byte(inputTxt)))
	if err != nil {
		t.Fatal(err)
	}

	output := solvePuzzle1(input)

	if output != 2 {
		t.Fatalf("expected 2, got %d", output)
	}
}

func TestSolvePuzzle2(t *testing.T) {
	input, err := readInput(bytes.NewReader([]byte(inputTxt)))
	if err != nil {
		t.Fatal(err)
	}

	output := solvePuzzle2(input)

	if output != 1 {
		t.Fatalf("expected 1, got %d", output)
	}
}
