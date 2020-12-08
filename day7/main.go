package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strings"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	m, err := readInput(f)
	if err != nil {
		log.Fatal(err)
	}

	answer := solvePuzzle1(m)

	fmt.Printf("Puzzle 1 answer: %d\n", answer)
}

// How many bag colors can eventually contain at least one shiny gold bag?
func solvePuzzle1(m bagMap) int {
	bags := m.findBagsContaining("shiny gold")
	return len(bags)
}

type bagMap map[string][]string

func (m bagMap) findBagsContaining(search string) []string {
	found := make(map[string]struct{})
	var recursion func(key string, colors []string)
	recursion = func(key string, colors []string) {
		for _, c := range colors {
			if c == search {
				found[key] = struct{}{}
			}
			if len(m[c]) > 0 {
				recursion(key, m[c])
			}
		}
	}
	for k, v := range m {
		recursion(k, v)
	}

	var out []string
	for k := range found {
		out = append(out, k)
	}
	return out
}

func readInput(r io.Reader) (bagMap, error) {
	m := make(bagMap)
	s := bufio.NewScanner(r)
	for s.Scan() {
		l := strings.TrimSpace(s.Text())
		if l == "" {
			continue
		}
		color, colors := splitLine(l)
		m[color] = colors
	}
	if err := s.Err(); err != nil {
		return nil, err
	}
	return m, nil
}

var containRegex = regexp.MustCompile(`(?:\d+ ([^,]+) bags?(?:, )?)+?`)

func splitLine(l string) (string, []string) {
	tmp := strings.SplitN(l, "contain", 2)
	color := strings.TrimSuffix(strings.TrimSpace(tmp[0]), " bags")
	matches := containRegex.FindAllStringSubmatch(strings.TrimSpace(tmp[1]), -1)
	var colors []string
	for i := range matches {
		colors = append(colors, matches[i][1])
	}
	return color, colors
}
