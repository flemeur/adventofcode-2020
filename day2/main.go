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

	input, err := readInput(f)
	if err != nil {
		log.Fatal(err)
	}

	answer := solvePuzzle1(input)

	fmt.Printf("Puzzle 1 answer: %d\n", answer)

	answer = solvePuzzle2(input)

	fmt.Printf("Puzzle 2 answer: %d\n", answer)
}

type line struct {
	n, m     int
	letter   rune
	password string
}

func solvePuzzle1(lines []line) int {
	var count int
	for _, l := range lines {
		n := strings.Count(l.password, string(l.letter))
		if n >= l.n && n <= l.m {
			count++
		}
	}
	return count
}

func solvePuzzle2(lines []line) int {
	var count int
	for _, l := range lines {
		passwdRunes := []rune(l.password)
		atPosN := passwdRunes[l.n-1] == l.letter
		atPosM := passwdRunes[l.m-1] == l.letter
		if atPosN != atPosM {
			count++
		}
	}
	return count
}

func readInput(r io.Reader) ([]line, error) {
	var parsed []line

	s := bufio.NewScanner(r)

	for s.Scan() {
		txt := strings.TrimSpace(s.Text())
		if txt == "" {
			continue
		}

		var l line
		_, err := fmt.Sscanf(txt, "%d-%d %c: %s", &l.n, &l.m, &l.letter, &l.password)
		if err != nil {
			log.Printf("could not scan line: %q: %v", txt, err)
			continue
		}

		parsed = append(parsed, l)
	}

	if err := s.Err(); err != nil {
		return nil, err
	}

	return parsed, nil
}
