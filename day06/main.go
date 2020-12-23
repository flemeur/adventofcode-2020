package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	groups, err := readInput(f)
	if err != nil {
		log.Fatal(err)
	}

	answer := solvePuzzle1(groups)

	fmt.Printf("Puzzle 1 answer: %d\n", answer)

	answer = solvePuzzle2(groups)

	fmt.Printf("Puzzle 2 answer: %d\n", answer)
}

func solvePuzzle1(groups [][]string) int {
	var sum int
	for _, g := range groups {
		seen := make(map[rune]struct{})
		for _, l := range g {
			for _, r := range l {
				seen[r] = struct{}{}
			}
		}
		sum += len(seen)
	}
	return sum
}

func solvePuzzle2(groups [][]string) int {
	var sum int
	for _, g := range groups {
		seen := make(map[rune]int)
		for _, l := range g {
			for _, r := range l {
				seen[r]++
			}
		}

		for _, n := range seen {
			if n == len(g) {
				sum++
			}
		}
	}
	return sum
}

func readInput(r io.Reader) ([][]string, error) {
	input, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}

	var groups [][]string

	lineGroups := bytes.Split(input, []byte("\n\n"))
	for _, lg := range lineGroups {
		lg = bytes.TrimSpace(lg)
		if len(lg) == 0 {
			continue
		}

		group := strings.Split(string(lg), "\n")

		groups = append(groups, group)
	}
	return groups, nil
}
