package main

import (
	"fmt"
	"io"
	"io/ioutil"
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

	answer := solvePuzzle1(numbers)

	fmt.Printf("Puzzle 1 answer: %d\n", answer)

	answer = solvePuzzle2(numbers)

	fmt.Printf("Puzzle 2 answer: %d\n", answer)
}

func solvePuzzle1(start []int) int {
	return findNthNumberSpoken(2020, start)
}

func findNthNumberSpoken(n int, start []int) int {
	pastIndices := make([]int, n)
	for i := range pastIndices {
		pastIndices[i] = -1
	}
	for i, n := range start {
		pastIndices[n] = i
	}

	lastNum := start[len(start)-1]

	for i := len(start); i < n; i++ {
		var nextNum int

		pastIndex := pastIndices[lastNum]
		if pastIndex > -1 && pastIndex < i-1 {
			nextNum = i - 1 - pastIndex
		}
		pastIndices[lastNum] = i - 1

		if i == n-1 {
			return nextNum
		}
		lastNum = nextNum
	}
	return -1
}

func solvePuzzle2(start []int) int {
	return findNthNumberSpoken(30000000, start)
}

func readInput(r io.Reader) ([]int, error) {
	var numbers []int

	b, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}

	nStrings := strings.Split(strings.TrimSpace(string(b)), ",")
	for _, s := range nStrings {
		n, err := strconv.Atoi(s)
		if err != nil {
			return nil, err
		}
		numbers = append(numbers, n)
	}
	return numbers, nil
}
