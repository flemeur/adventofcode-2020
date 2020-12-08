package main

import (
	"strings"
	"testing"
)

func TestReadInput(t *testing.T) {
	input := `light red bags contain 1 bright white bag, 2 muted yellow bags.
dark orange bags contain 3 bright white bags, 4 muted yellow bags.
bright white bags contain 1 shiny gold bag.
muted yellow bags contain 2 shiny gold bags, 9 faded blue bags.
shiny gold bags contain 1 dark olive bag, 2 vibrant plum bags.
dark olive bags contain 3 faded blue bags, 4 dotted black bags.
vibrant plum bags contain 5 faded blue bags, 6 dotted black bags.
faded blue bags contain no other bags.
dotted black bags contain no other bags.`

	m, err := readInput(strings.NewReader(input))
	if err != nil {
		t.Fatal(err)
	}

	if len(m) != 9 {
		t.Fatalf("expected 9 map entries, got %d", len(m))
	}
}

/*
In the above rules, the following options would be available to you:

A bright white bag, which can hold your shiny gold bag directly.
A muted yellow bag, which can hold your shiny gold bag directly, plus some other bags.
A dark orange bag, which can hold bright white and muted yellow bags, either of which could then hold your shiny gold bag.
A light red bag, which can hold bright white and muted yellow bags, either of which could then hold your shiny gold bag.

So, in this example, the number of bag colors that can eventually contain at least one shiny gold bag is 4.
*/

func TestFind(t *testing.T) {
	input := `light red bags contain 1 bright white bag, 2 muted yellow bags.
	dark orange bags contain 3 bright white bags, 4 muted yellow bags.
	bright white bags contain 1 shiny gold bag.
	muted yellow bags contain 2 shiny gold bags, 9 faded blue bags.
	shiny gold bags contain 1 dark olive bag, 2 vibrant plum bags.
	dark olive bags contain 3 faded blue bags, 4 dotted black bags.
	vibrant plum bags contain 5 faded blue bags, 6 dotted black bags.
	faded blue bags contain no other bags.
	dotted black bags contain no other bags.`

	m, err := readInput(strings.NewReader(input))
	if err != nil {
		t.Fatal(err)
	}

	bags := m.findBagsContaining("shiny gold")

	expected := []string{
		"bright white",
		"muted yellow",
		"dark orange",
		"light red",
	}

	if len(bags) != len(expected) {
		t.Errorf("expected length of valid bags to be %d, got %d", len(expected), len(bags))
	}

	contains := func(bags []string, color string) bool {
		for _, b := range bags {
			if b == color {
				return true
			}
		}
		return false
	}

	for _, b := range expected {
		if !contains(bags, b) {
			t.Errorf("expected to find bag %q in list of valid bags", b)
		}
	}
}
