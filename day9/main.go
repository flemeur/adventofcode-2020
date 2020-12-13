package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	numbers, err := readInput(f)
	if err != nil {
		log.Fatal(err)
	}

	answer := solvePuzzle1(numbers, 25)

	fmt.Printf("Puzzle 1 answer: %d\n", answer)

	answer = solvePuzzle2(numbers, answer)

	fmt.Printf("Puzzle 2 answer: %d\n", answer)
}

func solvePuzzle1(numbers []int, preamble int) int {
	for i, n := range numbers[preamble:] {
		if !checkNum(n, numbers[i:i+preamble]) {
			return n
		}
	}
	return -1
}

func checkNum(n int, preamble []int) bool {
	for i := 0; i < len(preamble); i++ {
		for j := i + 1; j < len(preamble); j++ {
			if preamble[i]+preamble[j] == n {
				return true
			}
		}
	}
	return false
}

func solvePuzzle2(numbers []int, target int) int {
	var set []int
Outer:
	for i := range numbers {
		set = []int{}
		for _, n := range numbers[i:] {
			if len(set) >= 2 && sum(set) == target {
				break Outer
			} else if sum(set) > target {
				continue Outer
			}
			set = append(set, n)
		}
	}
	min, max := minMax(set)
	return min + max
}

func sum(numbers []int) int {
	var sum int
	for _, n := range numbers {
		sum += n
	}
	return sum
}

func minMax(numbers []int) (min int, max int) {
	min = numbers[0]
	for _, n := range numbers {
		if n < min {
			min = n
		}
		if n > max {
			max = n
		}
	}
	return min, max
}

func readInput(r io.Reader) ([]int, error) {
	var numbers []int
	s := bufio.NewScanner(r)
	for s.Scan() {
		txt := strings.TrimSpace(s.Text())
		if txt == "" {
			continue
		}

		n, err := strconv.Atoi(txt)
		if err != nil {
			return nil, err
		}

		numbers = append(numbers, n)
	}
	if err := s.Err(); err != nil {
		return nil, err
	}
	return numbers, nil
}
