package main

import (
	"strings"
	"testing"
)

var input = `
..##.......
#...#...#..
.#....#..#.
..#.#...#.#
.#...##..#.
..#.##.....
.#.#.#....#
.#........#
#.##...#...
#...##....#
.#..#...#.#
`

func TestReadGrid(t *testing.T) {
	grid, err := readInput(strings.NewReader(input))
	if err != nil {
		t.Fatal(err)
	}

	rows := len(grid)
	if rows != 11 {
		t.Fatalf("expected 11 rows, got %d", rows)
	}

	columns := len(grid[0])
	if columns != 11 {
		t.Fatalf("expected 11 columns, got %d", columns)
	}
}

func TestSolvePuzzle1(t *testing.T) {
	grid, err := readInput(strings.NewReader(input))
	if err != nil {
		t.Fatal(err)
	}

	answer := solvePuzzle1(grid)

	if answer != 7 {
		t.Fatalf("expected 7, got %d", answer)
	}
}

func TestSolvePuzzle2(t *testing.T) {
	grid, err := readInput(strings.NewReader(input))
	if err != nil {
		t.Fatal(err)
	}

	answer := solvePuzzle2(grid)

	if answer != 336 {
		t.Fatalf("expected 336, got %d", answer)
	}
}
