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
	defer f.Close()

	program, err := readInput(f)
	if err != nil {
		log.Fatal(err)
	}

	answer := solvePuzzle1(program)

	fmt.Printf("Puzzle 1 answer: %d\n", answer)

	answer = solvePuzzle2(program)

	fmt.Printf("Puzzle 2 answer: %d\n", answer)
}

func solvePuzzle1(p *program) int {
	seen := make(map[int]struct{})
	for {
		seen[p.nextInstruction] = struct{}{}
		p.step()
		if _, ok := seen[p.nextInstruction]; ok {
			break
		}
	}
	return p.accumulator
}

func solvePuzzle2(prog *program) int {
	var acc int
Outer:
	for i := range prog.instructions {
		p := prog.copy()

		if p.instructions[i].operation == "nop" {
			p.instructions[i].operation = "jmp"
		} else if p.instructions[i].operation == "jmp" {
			p.instructions[i].operation = "nop"
		}

		seen := make(map[int]struct{})
		for {
			seen[p.nextInstruction] = struct{}{}
			p.step()
			if _, ok := seen[p.nextInstruction]; ok {
				break
			}
			if p.nextInstruction == len(p.instructions) {
				acc = p.accumulator
				break Outer
			}
		}
	}
	return acc
}

type program struct {
	instructions    []instruction
	accumulator     int
	nextInstruction int
}

func (p *program) step() {
	instruction := p.instructions[p.nextInstruction]
	switch instruction.operation {
	case "acc":
		p.accumulator += instruction.argument
	case "jmp":
		p.nextInstruction += instruction.argument
		return
	}
	p.nextInstruction++
}

func (p *program) copy() *program {
	i := make([]instruction, len(p.instructions))
	copy(i, p.instructions)
	return &program{instructions: i}
}

type instruction struct {
	operation string
	argument  int
}

func readInput(r io.Reader) (*program, error) {
	var instructions []instruction
	s := bufio.NewScanner(r)
	for s.Scan() {
		txt := strings.TrimSpace(s.Text())
		if txt == "" {
			continue
		}

		split := strings.Split(txt, " ")
		arg, _ := strconv.Atoi(split[1])

		instructions = append(instructions, instruction{
			operation: split[0],
			argument:  arg,
		})
	}
	if err := s.Err(); err != nil {
		return nil, err
	}
	return &program{instructions: instructions}, nil
}
