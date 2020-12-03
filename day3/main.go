package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	grid, err := readInput(f)
	if err != nil {
		log.Fatal(err)
	}

	answer := solvePuzzle1(grid)

	fmt.Printf("Puzzle 1 answer: %d\n", answer)

	answer = solvePuzzle2(grid)

	fmt.Printf("Puzzle 2 answer: %d\n", answer)
}

func solvePuzzle1(g grid) int {
	return followSlope(g, 3, 1)
}

func solvePuzzle2(g grid) int {
	answer := 1
	slopes := []struct {
		x, y int
	}{
		{x: 1, y: 1},
		{x: 3, y: 1},
		{x: 5, y: 1},
		{x: 7, y: 1},
		{x: 1, y: 2},
	}
	for _, s := range slopes {
		answer *= followSlope(g, s.x, s.y)
	}
	return answer
}

func followSlope(g grid, slopeX, slopeY int) int {
	var (
		nTrees     int
		posX, posY int
	)
	move := func() {
		posX += slopeX
		posY += slopeY
	}
	for !g.isBottom(posY) {
		if g.isTree(posX, posY) {
			nTrees++
		}
		move()
	}
	return nTrees
}

type grid [][]bool

func (g grid) isTree(x, y int) bool {
	if len(g) == 0 || y >= len(g) {
		return false
	}
	x = x % len(g[0])
	return g[y][x]
}

func (g grid) isBottom(y int) bool {
	return y >= len(g)
}

func readInput(r io.Reader) (grid, error) {
	var stride int
	var g grid
	s := bufio.NewScanner(r)
	for s.Scan() {
		line := strings.TrimSpace(s.Text())
		if line == "" {
			continue
		}
		if stride == 0 {
			stride = len(line)
		} else if len(line) != stride {
			return nil, fmt.Errorf("inconsistent line length, expected %d, got %d", stride, len(line))
		}

		row := make([]bool, stride)
		for i, c := range line {
			row[i] = c == '#' // A tree
		}
		g = append(g, row)
	}
	if err := s.Err(); err != nil {
		return nil, err
	}
	return g, nil
}
