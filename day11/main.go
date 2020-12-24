package main

import (
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

	seats, err := readInput(f)
	if err != nil {
		log.Fatal(err)
	}

	answer := solvePuzzle1(seats)

	fmt.Printf("Puzzle 1 answer: %d\n", answer)

	answer = solvePuzzle2(seats)

	fmt.Printf("Puzzle 2 answer: %d\n", answer)
}

func solvePuzzle1(seats []string) int {
	var previous []string
	for i := 0; i < 100000; i++ {
		previous = seats
		seats = iterate1(seats)
		if equalSeats(previous, seats) {
			break
		}
	}
	return countOccupied(seats)
}

func countOccupied(seats []string) int {
	var count int
	for _, row := range seats {
		count += strings.Count(row, string(occupied))
	}
	return count
}

func equalSeats(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// Seat states
const (
	empty    = 'L'
	occupied = '#'
)

func iterate1(seats []string) []string {
	newSeats := make([]string, len(seats))
	for y := range seats {
		var row string
		for x, s := range seats[y] {
			if s == empty && willSit1(seats, y, x) {
				row += string(occupied)
			} else if s == occupied && willLeave1(seats, y, x) {
				row += string(empty)
			} else {
				row += string(s)
			}
		}
		newSeats[y] = row
	}
	return newSeats
}

func willSit1(seats []string, y, x int) bool {
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if i == 0 && j == 0 ||
				y+i < 0 || y+i > len(seats)-1 ||
				x+j < 0 || x+j > len(seats[y])-1 {
				continue
			}
			if seats[y+i][x+j] == occupied {
				return false
			}
		}
	}
	return true
}

func willLeave1(seats []string, y, x int) bool {
	var nOccupied int
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if i == 0 && j == 0 ||
				y+i < 0 || y+i > len(seats)-1 ||
				x+j < 0 || x+j > len(seats[y])-1 {
				continue
			}
			if seats[y+i][x+j] == occupied {
				nOccupied++
			}
		}
	}
	return nOccupied >= 4
}

func solvePuzzle2(seats []string) int {
	var previous []string
	for i := 0; i < 100000; i++ {
		previous = seats
		seats = iterate2(seats)
		if equalSeats(previous, seats) {
			break
		}
	}
	return countOccupied(seats)
}

func iterate2(seats []string) []string {
	newSeats := make([]string, len(seats))
	for y := range seats {
		var row string
		for x, s := range seats[y] {
			if s == empty && willSit2(seats, y, x) {
				row += string(occupied)
			} else if s == occupied && willLeave2(seats, y, x) {
				row += string(empty)
			} else {
				row += string(s)
			}
		}
		newSeats[y] = row
	}
	return newSeats
}

type direction uint

// Directions
const (
	up    direction = iota // y-i, x
	down                   // y+i, x
	left                   // y, x-i
	right                  // y, x+i
	nw                     // y-i, x-i
	sw                     // y+i, x-i
	ne                     // y-i, x+i
	se                     // y+i, x+i
)

var directions = map[direction][]int{
	up:    {-1, 0},
	down:  {1, 0},
	left:  {0, -1},
	right: {0, 1},
	nw:    {-1, -1},
	sw:    {1, -1},
	ne:    {-1, 1},
	se:    {1, 1},
}

func willSit2(seats []string, y, x int) bool {
	blocked := make(map[direction]bool)
	for i := 1; ; i++ {
		for d, dm := range directions {
			if blocked[d] {
				continue
			}

			y1, x1 := y+i*dm[0], x+i*dm[1]
			if y1 >= 0 && y1 <= len(seats)-1 && x1 >= 0 && x1 <= len(seats[0])-1 {
				if seats[y1][x1] == empty {
					blocked[d] = true
					continue
				}
				if seats[y1][x1] == occupied {
					return false
				}
			}
		}

		// Out of bounds
		if y-i <= 0 && x-i <= 0 && y+i >= len(seats)-1 && x+i >= len(seats[0])-1 {
			break
		}
	}
	return true
}

func willLeave2(seats []string, y, x int) bool {
	var nOccupied int
	blocked := make(map[direction]bool)
	for i := 1; ; i++ {
		for d, dm := range directions {
			if blocked[d] {
				continue
			}

			y1, x1 := y+i*dm[0], x+i*dm[1]
			if y1 >= 0 && y1 <= len(seats)-1 && x1 >= 0 && x1 <= len(seats[0])-1 {
				if seats[y1][x1] == empty {
					blocked[d] = true
					continue
				}
				if seats[y1][x1] == occupied {
					blocked[d] = true
					nOccupied++
				}
			}
		}

		// Out of bounds
		if y-i <= 0 && x-i <= 0 && y+i >= len(seats)-1 && x+i >= len(seats[0])-1 {
			break
		}
	}
	return nOccupied >= 5
}

func readInput(r io.Reader) ([]string, error) {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}
	return strings.Split(strings.TrimSpace(string(b)), "\n"), nil
}
