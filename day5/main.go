package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strings"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	lines, err := readInput(f)
	if err != nil {
		log.Fatal(err)
	}

	answer, err := solvePuzzle1(lines)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Puzzle 1 answer: %d\n", answer)

	answer, err = solvePuzzle2(lines)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Puzzle 2 answer: %d\n", answer)
}

func solvePuzzle1(lines []string) (int, error) {
	var maxID int
	for _, l := range lines {
		_, _, id, err := parseSeat(l)
		if err != nil {
			return 0, err
		}
		if id > maxID {
			maxID = id
		}
	}
	return maxID, nil
}

func solvePuzzle2(lines []string) (int, error) {
	var mySeatID int
	var seatIDs []int
	for _, l := range lines {
		_, _, id, err := parseSeat(l)
		if err != nil {
			return 0, err
		}
		seatIDs = append(seatIDs, id)
	}
	sort.Ints(seatIDs)
	var prevID int
	for _, id := range seatIDs {
		if prevID > 0 {
			if id > prevID+1 {
				mySeatID = prevID + 1
				break
			}
		}
		prevID = id
	}
	return mySeatID, nil
}

func parseSeat(str string) (row, column, id int, err error) {
	row, err = parseRow(str[:7])
	if err != nil {
		return 0, 0, 0, err
	}
	column, err = parseColumn(str[7:])
	if err != nil {
		return 0, 0, 0, err
	}
	return row, column, row*8 + column, nil
}

func parseRow(str string) (int, error) {
	high := 127
	low := 0
	for _, r := range str[:6] {
		if r == 'F' {
			high = (high-low)/2 + low
		} else if r == 'B' {
			low = (high-low)/2 + 1 + low
		} else {
			return 0, fmt.Errorf("invalid rune %c", r)
		}
	}
	if str[6] == 'F' {
		return low, nil
	} else if str[6] == 'B' {
		return high, nil
	}
	return 0, fmt.Errorf("invalid final rune %c", str[6])
}

func parseColumn(str string) (int, error) {
	high := 7
	low := 0
	for _, r := range str[:2] {
		if r == 'L' {
			high = (high-low)/2 + low
		} else if r == 'R' {
			low = (high-low)/2 + 1 + low
		} else {
			return 0, fmt.Errorf("invalid rune %c", r)
		}
	}
	if str[2] == 'L' {
		return low, nil
	} else if str[2] == 'R' {
		return high, nil
	}
	return 0, fmt.Errorf("invalid final rune %c", str[2])
}

func readInput(r io.Reader) ([]string, error) {
	var out []string
	s := bufio.NewScanner(r)
	for s.Scan() {
		txt := strings.TrimSpace(s.Text())
		if txt == "" {
			continue
		}

		out = append(out, txt)
	}
	if err := s.Err(); err != nil {
		return nil, err
	}
	return out, nil
}
