package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strconv"
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

	answer = solvePuzzle2(m)

	fmt.Printf("Puzzle 2 answer: %d\n", answer)
}

// How many bag colors can eventually contain at least one shiny gold bag?
func solvePuzzle1(m bagMap) int {
	bags := m.findBagsContaining("shiny gold")
	return len(bags)
}

func solvePuzzle2(m bagMap) int {
	return m.countBags("shiny gold")
}

type bagMap map[string][]bag

type bag struct {
	count int
	color string
}

func (m bagMap) countBags(search string) int {
	var recursion func(key string, bags []bag) int
	recursion = func(key string, bags []bag) int {
		var count int
		for _, b := range bags {
			count += b.count + b.count*recursion(b.color, m[b.color])
		}
		return count
	}
	if _, ok := m[search]; !ok {
		return 0
	}
	return recursion(search, m[search])
}

func (m bagMap) findBagsContaining(search string) []string {
	found := make(map[string]struct{})
	var recursion func(key string, bags []bag)
	recursion = func(key string, bags []bag) {
		for _, b := range bags {
			if b.color == search {
				found[key] = struct{}{}
			}
			if len(m[b.color]) > 0 {
				recursion(key, m[b.color])
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
		color, bags := splitLine(l)
		m[color] = bags
	}
	if err := s.Err(); err != nil {
		return nil, err
	}
	return m, nil
}

var containRegex = regexp.MustCompile(`(?:(\d+) ([^,]+) bags?(?:, )?)+?`)

func splitLine(l string) (string, []bag) {
	tmp := strings.SplitN(l, "contain", 2)
	color := strings.TrimSuffix(strings.TrimSpace(tmp[0]), " bags")
	matches := containRegex.FindAllStringSubmatch(strings.TrimSpace(tmp[1]), -1)
	var bags []bag
	for i := range matches {
		count, _ := strconv.Atoi(matches[i][1])
		bags = append(bags, bag{
			count: count,
			color: matches[i][2],
		})

	}
	return color, bags
}
