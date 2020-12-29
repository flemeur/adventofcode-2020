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

	rules, ticket, tickets, err := readInput(f)
	if err != nil {
		log.Fatal(err)
	}

	answer := solvePuzzle1(rules, tickets)

	fmt.Printf("Puzzle 1 answer: %d\n", answer)

	answer = solvePuzzle2(rules, ticket, tickets)

	fmt.Printf("Puzzle 2 answer: %d\n", answer)
}

func solvePuzzle1(rules map[string]Rule, tickets []Ticket) int {
	var invalid []int
	for _, t := range tickets {
		invalid = append(invalid, checkTicket(t, rules)...)
	}
	var sum int
	for _, n := range invalid {
		sum += n
	}
	return sum
}

func solvePuzzle2(rules map[string]Rule, ownTicket Ticket, tickets []Ticket) int {
	var validTickets []Ticket
	for _, t := range tickets {
		invalid := checkTicket(t, rules)
		if len(invalid) == 0 {
			validTickets = append(validTickets, t)
		}
	}

	mapping := findTicketFields(rules, ownTicket, validTickets)

	return mapping["departure location"] *
		mapping["departure station"] *
		mapping["departure platform"] *
		mapping["departure track"] *
		mapping["departure date"] *
		mapping["departure time"]
}

func findTicketFields(rules map[string]Rule, ownTicket Ticket, tickets []Ticket) map[string]int {
	possibleKeys := make([]map[string]struct{}, len(ownTicket))
	for i := range possibleKeys {
		possibleKeys[i] = make(map[string]struct{})
		for k := range rules {
			possibleKeys[i][k] = struct{}{}
		}
	}

	for _, t := range tickets {
		for i, n := range t {
			for k, r := range rules {
				if !r.Check(n) {
					delete(possibleKeys[i], k)
				}
			}
		}
	}

	seen := make(map[string]struct{})
	for {
		for i := range possibleKeys {
			if len(possibleKeys[i]) == 1 {
				for k := range possibleKeys[i] {
					seen[k] = struct{}{}
				}
				continue
			}

			for k := range possibleKeys[i] {
				if _, ok := seen[k]; ok {
					delete(possibleKeys[i], k)
				}
			}

		}

		if len(seen) == len(ownTicket) {
			break
		}
	}

	mapping := make(map[string]int)

	for i := range possibleKeys {
		for k := range possibleKeys[i] {
			if _, ok := mapping[k]; !ok {
				mapping[k] = ownTicket[i]
			}
		}
	}

	return mapping
}

func checkTicket(ticket Ticket, rules map[string]Rule) []int {
	var invalid []int
Outer:
	for _, n := range ticket {
		for _, r := range rules {
			if r.Check(n) {
				continue Outer
			}
		}
		invalid = append(invalid, n)
	}
	return invalid
}

func NewRule(r string) (Rule, error) {
	var rule Rule
	parts := strings.Split(r, " or ")
	for _, p := range parts {
		var min, max int
		_, err := fmt.Sscanf(p, "%d-%d", &min, &max)
		if err != nil {
			return rule, err
		}
		rule = append(rule, RulePart{
			min: min,
			max: max,
		})
	}
	return rule, nil
}

type Rule []RulePart

func (r Rule) Check(n int) bool {
	for _, p := range r {
		if n >= p.min && n <= p.max {
			return true
		}
	}
	return false
}

type RulePart struct {
	min, max int
}

type Ticket []int

func readInput(r io.Reader) (rules map[string]Rule, ticket Ticket, tickets []Ticket, err error) {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, nil, nil, err
	}
	parts := strings.SplitN(strings.TrimSpace(string(b)), "\n\n", 3)

	// Parse rules
	rules, err = parseRules(parts[0])
	if err != nil {
		return nil, nil, nil, err
	}

	// Parse own ticket
	lines := strings.SplitN(parts[1], "\n", 2)
	ticket, err = parseTicket(lines[1])
	if err != nil {
		return nil, nil, nil, err
	}

	// Parse nearby tickets
	lines = strings.SplitN(parts[2], "\n", 2)
	tickets, err = parseTickets(lines[1])
	if err != nil {
		return nil, nil, nil, err
	}

	return rules, ticket, tickets, nil
}

func parseRules(input string) (map[string]Rule, error) {
	rules := make(map[string]Rule)
	lines := strings.Split(input, "\n")
	for _, l := range lines {
		kv := strings.SplitN(l, ": ", 2)
		r, err := NewRule(kv[1])
		if err != nil {
			return nil, err
		}
		rules[kv[0]] = r
	}
	return rules, nil
}

func parseTicket(input string) (Ticket, error) {
	var t Ticket
	parts := strings.Split(input, ",")
	for _, p := range parts {
		n, err := strconv.Atoi(p)
		if err != nil {
			return nil, err
		}
		t = append(t, n)
	}
	return t, nil
}

func parseTickets(input string) ([]Ticket, error) {
	var tickets []Ticket
	lines := strings.Split(input, "\n")
	for _, l := range lines {
		t, err := parseTicket(l)
		if err != nil {
			return nil, err
		}
		tickets = append(tickets, t)
	}
	return tickets, nil
}
