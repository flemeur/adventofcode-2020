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

	input, err := readInput(f)
	if err != nil {
		log.Fatal(err)
	}

	answer := solvePuzzle1(input)

	fmt.Printf("Puzzle 1 answer: %d\n", answer)

	answer = solvePuzzle2(input)

	fmt.Printf("Puzzle 2 answer: %d\n", answer)
}

func solvePuzzle1(input []int) int {
	var sum, output int
	for i, a := range input {
		for j, b := range input {
			if i == j {
				continue
			}

			sum = a + b
			if sum == 2020 {
				output = a * b
				break
			}
		}
	}
	return output
}

func solvePuzzle2(input []int) int {
	var sum, output int
	for i, a := range input {
		for j, b := range input {
			for k, c := range input {
				if i == j && i == k {
					continue
				}

				sum = a + b + c
				if sum == 2020 {
					output = a * b * c
					break
				}
			}
		}
	}
	return output
}

func readInput(r io.Reader) ([]int, error) {
	var parsed []int

	s := bufio.NewScanner(r)

	for s.Scan() {
		txt := strings.TrimSpace(s.Text())
		if txt == "" {
			continue
		}

		i, err := strconv.Atoi(txt)
		if err != nil {
			continue
		}

		parsed = append(parsed, i)
	}

	if err := s.Err(); err != nil {
		return nil, err
	}

	return parsed, nil
}
