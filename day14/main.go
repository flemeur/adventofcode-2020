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

func solvePuzzle1(lines []string) (uint, error) {
	var mask string
	mem := make(map[uint]uint)
	for _, l := range lines {
		parts := strings.SplitN(l, " = ", 2)
		if parts[0] == "mask" {
			mask = parts[1]
			continue
		}

		var addr uint
		_, err := fmt.Sscanf(parts[0], "mem[%d]", &addr)
		if err != nil {
			return 0, err
		}

		n, err := strconv.Atoi(parts[1])
		if err != nil {
			return 0, err
		}

		binBytes := []byte(fmt.Sprintf("%036b", n))

		for i := range mask {
			if mask[i] != 'X' {
				binBytes[i] = mask[i]
			}
		}

		val, err := strconv.ParseUint(string(binBytes), 2, 64)
		if err != nil {
			return 0, err
		}

		mem[addr] = uint(val)
	}

	var sum uint
	for _, v := range mem {
		sum += v
	}
	return sum, nil
}

func solvePuzzle2(lines []string) (uint, error) {
	var mask string
	mem := make(map[uint]uint)
	for _, l := range lines {
		parts := strings.SplitN(l, " = ", 2)
		if parts[0] == "mask" {
			mask = parts[1]
			continue
		}

		var addr uint
		_, err := fmt.Sscanf(parts[0], "mem[%d]", &addr)
		if err != nil {
			return 0, err
		}

		addrBytes := []byte(fmt.Sprintf("%036b", addr))
		for i := range addrBytes {
			if mask[i] != '0' {
				addrBytes[i] = mask[i]
			}
		}

		addrTemplate := string(addrBytes)
		binStrings := generateBinaryStrings(addrTemplate)

		var addresses []uint
		for _, binStr := range binStrings {
			v, err := strconv.ParseUint(binStr, 2, 64)
			if err != nil {
				return 0, err
			}
			addresses = append(addresses, uint(v))
		}

		val, err := strconv.Atoi(parts[1])
		if err != nil {
			return 0, err
		}

		for _, addr := range addresses {
			mem[addr] = uint(val)
		}
	}

	var sum uint
	for _, v := range mem {
		sum += v
	}
	return sum, nil
}

func generateBinaryStrings(template string) []string {
	var binStrings []string

	xIndex := strings.LastIndexByte(template, 'X')
	if xIndex == -1 {
		binStrings = append(binStrings, template)
		return binStrings
	}

	for j := 0; j < 2; j++ {
		tmp := []byte(template)
		tmp[xIndex] = byte('0' + j)
		binStrings = append(binStrings, generateBinaryStrings(string(tmp))...)
	}

	return binStrings
}

func readInput(r io.Reader) ([]string, error) {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}
	return strings.Split(strings.TrimSpace(string(b)), "\n"), nil
}
