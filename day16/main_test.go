package main

import (
	"strings"
	"testing"
)

func TestSolvePuzzle1(t *testing.T) {
	input := `
class: 1-3 or 5-7
row: 6-11 or 33-44
seat: 13-40 or 45-50

your ticket:
7,1,14

nearby tickets:
7,3,47
40,4,50
55,2,20
38,6,12
`

	rules, _, tickets, err := readInput(strings.NewReader(input))
	if err != nil {
		t.Fatal(err)
	}

	answer := solvePuzzle1(rules, tickets)
	if answer != 71 {
		t.Fatalf("expected answer 71, got %d", answer)
	}
}

func TestFindTicketFields(t *testing.T) {
	input := `
class: 0-1 or 4-19
row: 0-5 or 8-19
seat: 0-13 or 16-19

your ticket:
11,12,13

nearby tickets:
3,9,18
15,1,5
5,14,9
`

	rules, ticket, tickets, err := readInput(strings.NewReader(input))
	if err != nil {
		t.Fatal(err)
	}

	mapping := findTicketFields(rules, ticket, tickets)

	if mapping["row"] != 11 {
		t.Fatalf("expected field %s to be %d, got %d", "row", 11, mapping["row"])
	}
	if mapping["class"] != 12 {
		t.Fatalf("expected field %s to be %d, got %d", "class", 12, mapping["class"])
	}
	if mapping["seat"] != 13 {
		t.Fatalf("expected field %s to be %d, got %d", "seat", 13, mapping["seat"])
	}
}
