package main

import (
	"bufio"
	"fmt"
	"io"
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

	instructions, err := readInput(f)
	if err != nil {
		log.Fatal(err)
	}

	answer := solvePuzzle1(instructions)

	fmt.Printf("Puzzle 1 answer: %d\n", answer)

	answer = solvePuzzle2(instructions)

	fmt.Printf("Puzzle 2 answer: %d\n", answer)
}

func solvePuzzle1(instructions []Instruction) int {
	dir := 90
	var posN, posE int
	for _, inst := range instructions {
		switch inst.Op {
		case 'N':
			posN += inst.Arg
		case 'S':
			posN -= inst.Arg
		case 'E':
			posE += inst.Arg
		case 'W':
			posE -= inst.Arg

		case 'L':
			dir = mod(dir-inst.Arg, 360)
		case 'R':
			dir = mod(dir+inst.Arg, 360)

		case 'F':
			switch dir {
			case 0:
				posN += inst.Arg
			case 90:
				posE += inst.Arg
			case 180:
				posN -= inst.Arg
			case 270:
				posE -= inst.Arg
			}
		}
	}
	return abs(posE) + abs(posN)
}

func solvePuzzle2(instructions []Instruction) int {
	s := NewShip()
	for _, inst := range instructions {
		s.Exec(inst)
	}
	return abs(s.East) + abs(s.North)
}

func NewShip() *Ship {
	return &Ship{
		// The waypoint starts 10 units east and 1 unit north relative to the ship.
		Waypoint: Waypoint{
			Position: Position{
				North: 1,
				East:  10,
			},
		},
	}
}

type Ship struct {
	Position
	Waypoint
}

func (s *Ship) Exec(inst Instruction) {
	switch inst.Op {
	case 'F':
		// Action F means to move forward to the waypoint a number of times equal to the given value.
		s.East += inst.Arg * s.Waypoint.East
		s.North += inst.Arg * s.Waypoint.North

	default:
		s.Waypoint.Exec(inst)
	}
}

func (s Ship) String() string {
	return fmt.Sprintf("ship: %s, waypoint: %s", s.Position.String(), s.Waypoint.Position.String())
}

// The waypoint is relative to the ship; that is, if the ship moves, the waypoint moves with it.
type Waypoint struct {
	Position
}

func (w *Waypoint) Exec(inst Instruction) {
	switch inst.Op {
	case 'N':
		w.North += inst.Arg
	case 'S':
		w.North -= inst.Arg
	case 'E':
		w.East += inst.Arg
	case 'W':
		w.East -= inst.Arg

	case 'L':
		// Action L means to rotate the waypoint around the ship left (counter-clockwise) the given number of degrees.
		pos := w.Position
		switch inst.Arg {
		case 0, 360:
			// Nothing
		case 90, 180, 270:
			rad := float64(inst.Arg) * math.Pi / 180
			// Could also be https://en.wikipedia.org/wiki/Rotation_matrix#Common_rotations
			w.East = pos.East*int(math.Cos(rad)) - pos.North*int(math.Sin(rad))
			w.North = pos.East*int(math.Sin(rad)) + pos.North*int(math.Cos(rad))
		default:
			panic(fmt.Sprintf("invalid rotation arg %d", inst.Arg))
		}
	case 'R':
		// Action R means to rotate the waypoint around the ship right (clockwise) the given number of degrees.
		pos := w.Position
		switch inst.Arg {
		case 0, 360:
			// Nothing
		case 90, 180, 270:
			rad := float64(inst.Arg) * math.Pi / 180
			// Could also be https://en.wikipedia.org/wiki/Rotation_matrix#Common_rotations
			w.East = pos.East*int(math.Cos(-rad)) - pos.North*int(math.Sin(-rad))
			w.North = pos.East*int(math.Sin(-rad)) + pos.North*int(math.Cos(-rad))
		default:
			panic(fmt.Sprintf("invalid rotation arg %d", inst.Arg))
		}

	default:
		panic(fmt.Sprintf("invalid op %c", inst.Op))
	}
}

type Position struct {
	North int
	East  int
}

func (p Position) String() string {
	return fmt.Sprintf("east %d, north %d", p.East, p.North)
}

func mod(a, b int) int {
	return (a%b + b) % b
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

type Instruction struct {
	Op  rune
	Arg int
}

func readInput(r io.Reader) ([]Instruction, error) {
	var instructions []Instruction
	s := bufio.NewScanner(r)
	for s.Scan() {
		txt := strings.TrimSpace(s.Text())
		if txt == "" {
			continue
		}

		op := rune(txt[0])

		arg, err := strconv.Atoi(txt[1:])
		if err != nil {
			return nil, err
		}

		instructions = append(instructions, Instruction{
			Op:  op,
			Arg: arg,
		})
	}
	if err := s.Err(); err != nil {
		return nil, err
	}
	return instructions, nil
}
