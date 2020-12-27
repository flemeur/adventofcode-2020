package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math"
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

	timestamp, busIDs, err := readInput(f)
	if err != nil {
		log.Fatal(err)
	}

	answer := solvePuzzle1(timestamp, busIDs)

	fmt.Printf("Puzzle 1 answer: %d\n", answer)

	answer = solvePuzzle2(busIDs)

	fmt.Printf("Puzzle 2 answer: %d\n", answer)
}

func solvePuzzle1(timestamp int, busIDs []int) int {
	times := make(map[int]int)
	for _, id := range busIDs {
		// Skip the Xs
		if id <= 0 {
			continue
		}

		var t int
		for {
			t += id
			if t >= timestamp {
				break
			}
		}
		times[id] = t
	}
	minDiff := math.MaxInt64
	var bestBus int
	for id, t := range times {
		diff := t - timestamp
		if diff < minDiff {
			minDiff = diff
			bestBus = id
		}
	}
	return minDiff * bestBus
}

func solvePuzzle2(busIDs []int) int {
	t := busIDs[0]
	tstep := t

	for i := 1; i < len(busIDs); i++ {
		if busIDs[i] <= 0 {
			continue
		}

		for ; (t+i)%busIDs[i] != 0; t += tstep {
			// Nothing
		}

		tstep *= busIDs[i]
	}
	return t
}

func readInput(r io.Reader) (int, []int, error) {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return 0, nil, err
	}

	parts := strings.Split(strings.TrimSpace(string(b)), "\n")

	timestamp, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, nil, err
	}

	var busIDs []int
	busses := strings.Split(parts[1], ",")
	for _, b := range busses {
		var (
			id  int
			err error
		)
		if b == "x" {
			id = -1
		} else {
			id, err = strconv.Atoi(b)
			if err != nil {
				return 0, nil, err
			}
		}
		busIDs = append(busIDs, id)
	}

	return timestamp, busIDs, nil
}
