package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	adapters, err := readInput(f)
	if err != nil {
		log.Fatal(err)
	}

	answer := solvePuzzle1(adapters)

	fmt.Printf("Puzzle 1 answer: %d\n", answer)

	answer = solvePuzzle2(adapters)

	fmt.Printf("Puzzle 2 answer: %d\n", answer)
}

func solvePuzzle1(adapters []int) int {
	var countDiff1, countDiff3 int
	var currentRating int
	for _, rating := range adapters {
		diff := rating - currentRating
		if diff == 1 {
			countDiff1++
		} else if diff == 3 {
			countDiff3++
		}
		currentRating = rating
	}
	return countDiff1 * countDiff3
}

func solvePuzzle2(adapters []int) int {
	return countArrangements(adapters, 0)
}

func countArrangements(adapters []int, i int) int {
	memo := make(map[int]int)
	var countFunc func([]int, int, map[int]int) int
	countFunc = func(adapters []int, i int, memo map[int]int) int {
		if i == len(adapters)-1 {
			return 1
		}
		if count, ok := memo[i]; ok {
			return count
		}
		var count int
		for j := 1; j <= 3 && i+j < len(adapters); j++ {
			if adapters[i+j]-adapters[i] > 3 {
				break
			}
			count += countFunc(adapters, i+j, memo)
		}
		memo[i] = count
		return count
	}
	return countFunc(adapters, i, memo)
}

func readInput(r io.Reader) ([]int, error) {
	// Start with 0 jolts
	adapters := []int{0}
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

		adapters = append(adapters, n)
	}
	if err := s.Err(); err != nil {
		return nil, err
	}
	// Sort the adapters
	sort.Ints(adapters)
	// Add the final +3 jolts
	adapters = append(adapters, adapters[len(adapters)-1]+3)
	return adapters, nil
}
