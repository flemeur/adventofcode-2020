package main

import (
	"strings"
	"testing"
)

func TestSolvePuzzle1(t *testing.T) {
	input := `
F10
N3
F7
R90
F11
`

	instructions, err := readInput(strings.NewReader(input))
	if err != nil {
		t.Fatal(err)
	}

	answer := solvePuzzle1(instructions)
	if answer != 25 {
		t.Fatalf("expected answer 25, got %d", answer)
	}
}

func TestSolvePuzzle2(t *testing.T) {
	input := `
F10
N3
F7
R90
F11
`

	instructions, err := readInput(strings.NewReader(input))
	if err != nil {
		t.Fatal(err)
	}

	// After these operations, the ship's Manhattan distance from its starting position is 214 + 72 = 286.
	answer := solvePuzzle2(instructions)
	if answer != 286 {
		t.Fatalf("expected answer 286, got %d", answer)
	}
}

func TestWaypointRotation(t *testing.T) {
	s := NewShip()
	if s.Waypoint.East != 10 && s.Waypoint.North != 1 {
		t.Fatalf("expected initial waypoint at E 10 N 1, got %s", s.Waypoint.String())
	}

	s.Exec(Instruction{Op: 'L', Arg: 90})
	if s.Waypoint.East != -1 && s.Waypoint.North != 10 {
		t.Fatalf("expected waypoint at E -1 N 10, got %s", s.Waypoint.String())
	}

	s = NewShip()

	s.Exec(Instruction{Op: 'R', Arg: 90})
	if s.Waypoint.East != 1 || s.Waypoint.North != -10 {
		t.Fatalf("expected waypoint at E 1 N -10, got %s", s.Waypoint.String())
	}
}

func TestShipMovement(t *testing.T) {
	s := NewShip()

	// F10 moves the ship to the waypoint 10 times (a total of 100 units east and 10 units north), leaving the ship at east 100, north 10. The waypoint stays 10 units east and 1 unit north of the ship.
	s.Exec(Instruction{Op: 'F', Arg: 10})
	if s.East != 100 || s.North != 10 {
		t.Fatalf("expected ship at E 100 N 10, got %s", s.Position.String())
	}

	// N3 moves the waypoint 3 units north to 10 units east and 4 units north of the ship. The ship remains at east 100, north 10.
	s.Exec(Instruction{Op: 'N', Arg: 3})
	if s.Waypoint.East != 10 || s.Waypoint.North != 4 {
		t.Fatalf("expected waypoint at E 10 N 4, got %s", s.Waypoint.String())
	}
	if s.East != 100 || s.North != 10 {
		t.Fatalf("expected ship at E 100 N 10, got %s", s.Position.String())
	}

	// F7 moves the ship to the waypoint 7 times (a total of 70 units east and 28 units north), leaving the ship at east 170, north 38. The waypoint stays 10 units east and 4 units north of the ship.
	s.Exec(Instruction{Op: 'F', Arg: 7})
	if s.Waypoint.East != 10 || s.Waypoint.North != 4 {
		t.Fatalf("expected waypoint at E 10 N 4, got %s", s.Waypoint.String())
	}
	if s.East != 170 || s.North != 38 {
		t.Fatalf("expected ship at E 170 N 38, got %s", s.Position.String())
	}

	// R90 rotates the waypoint around the ship clockwise 90 degrees, moving it to 4 units east and 10 units south of the ship. The ship remains at east 170, north 38.
	s.Exec(Instruction{Op: 'R', Arg: 90})
	if s.Waypoint.East != 4 && s.Waypoint.North != -10 {
		t.Fatalf("expected waypoint at E 4 N -10, got %s", s.Waypoint.String())
	}
	if s.East != 170 || s.North != 38 {
		t.Fatalf("expected ship at E 170 N 38, got %s", s.Position.String())
	}

	// F11 moves the ship to the waypoint 11 times (a total of 44 units east and 110 units south), leaving the ship at east 214, south 72. The waypoint stays 4 units east and 10 units south of the ship.
	s.Exec(Instruction{Op: 'F', Arg: 11})
	if s.Waypoint.East != 4 && s.Waypoint.North != -10 {
		t.Fatalf("expected waypoint at E 4 N -10, got %s", s.Waypoint.String())
	}
	if s.East != 214 || s.North != -72 {
		t.Fatalf("expected ship at E 214 N -72, got %s", s.Position.String())
	}
}
