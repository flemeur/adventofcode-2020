package main

import (
	"errors"
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

	pd, err := readInput(f)
	if err != nil {
		log.Fatal(err)
	}

	answer := solvePuzzle1(pd)

	fmt.Printf("Puzzle 1 answer: %d\n", answer)

	answer = solvePuzzle2(pd)

	fmt.Printf("Puzzle 2 answer: %d\n", answer)
}

func solvePuzzle1(pd PocketDimension) int {
	for i := 0; i < 6; i++ {
		pd = pd.Evolve()
	}
	return pd.CountActive()
}

func solvePuzzle2(pd PocketDimension) int {
	for i := 0; i < 6; i++ {
		pd = pd.Evolve4D()
	}
	return pd.CountActive()
}

type Point struct {
	X, Y, Z int
	W       int // Part 2
}

type PocketDimension map[Point]bool

func (pd PocketDimension) Evolve() PocketDimension {
	npd := make(PocketDimension)
	for p0 := range pd {
		for i := -1; i <= 1; i++ {
			for j := -1; j <= 1; j++ {
				for k := -1; k <= 1; k++ {
					p := Point{Z: p0.Z + i, Y: p0.Y + j, X: p0.X + k}
					nActiveNeighbors := pd.CountActiveNeighbors(p.X, p.Y, p.Z)
					if pd[p] {
						if nActiveNeighbors == 2 || nActiveNeighbors == 3 {
							npd[p] = true
						} else {
							npd[p] = false
						}
					} else {
						if nActiveNeighbors == 3 {
							npd[p] = true
						} else {
							npd[p] = false
						}
					}
				}
			}
		}

	}
	return npd
}

func (pd PocketDimension) Evolve4D() PocketDimension {
	npd := make(PocketDimension)
	for p0 := range pd {
		for i := -1; i <= 1; i++ {
			for j := -1; j <= 1; j++ {
				for k := -1; k <= 1; k++ {
					for n := -1; n <= 1; n++ {
						p := Point{X: p0.X + k, Y: p0.Y + j, Z: p0.Z + i, W: p0.W + n}
						nActiveNeighbors := pd.CountActiveNeighbors4D(p.X, p.Y, p.Z, p.W)
						if pd[p] {
							if nActiveNeighbors == 2 || nActiveNeighbors == 3 {
								npd[p] = true
							} else {
								npd[p] = false
							}
						} else {
							if nActiveNeighbors == 3 {
								npd[p] = true
							} else {
								npd[p] = false
							}
						}
					}
				}
			}
		}

	}
	return npd
}

func (pd PocketDimension) CountActive() int {
	var count int
	for p := range pd {
		if pd[p] {
			count++
		}
	}
	return count
}

func (pd PocketDimension) CountActiveNeighbors(x, y, z int) int {
	var count int
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			for k := -1; k <= 1; k++ {
				if i == 0 && j == 0 && k == 0 {
					continue
				}

				p := Point{X: x + k, Y: y + j, Z: z + i}
				if pd[p] {
					count++
				}
			}
		}
	}
	return count
}

func (pd PocketDimension) CountActiveNeighbors4D(x, y, z, w int) int {
	var count int
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			for k := -1; k <= 1; k++ {
				for n := -1; n <= 1; n++ {
					if i == 0 && j == 0 && k == 0 && n == 0 {
						continue
					}

					p := Point{X: x + k, Y: y + j, Z: z + i, W: w + n}
					if pd[p] {
						count++
					}
				}
			}
		}
	}
	return count
}

func readInput(r io.Reader) (PocketDimension, error) {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(strings.TrimSpace(string(b)), "\n")

	if len(lines) == 0 {
		return nil, errors.New("empty input")
	}

	pd := make(PocketDimension)
	for y, line := range lines {
		for x, c := range line {
			p := Point{X: x, Y: y}
			pd[p] = c == '#'
		}
	}
	return pd, nil
}
