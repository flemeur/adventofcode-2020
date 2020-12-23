package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	passports, err := readInput(f)
	if err != nil {
		log.Fatal(err)
	}

	answer := solvePuzzle1(passports)

	fmt.Printf("Puzzle 1 answer: %d\n", answer)

	answer = solvePuzzle2(passports)

	fmt.Printf("Puzzle 2 answer: %d\n", answer)
}

func solvePuzzle1(passports []passport) int {
	var count int
	for _, p := range passports {
		if p.hasRequiredFields() {
			count++
		}
	}
	return count
}

func solvePuzzle2(passports []passport) int {
	var count int
	for _, p := range passports {
		if p.valid() {
			count++
		}
	}
	return count
}

type passport struct {
	BirthYear      int        // byr
	IssueYear      int        // iyr
	ExpirationYear int        // eyr
	Height         height     // hgt
	HairColor      hairColor  // hcl
	EyeColor       eyeColor   // ecl
	PassportID     passportID // pid
	CountryID      string     // cid - Should be optional
}

func (p *passport) UnmarshalText(b []byte) error {
	fields := strings.Split(strings.ReplaceAll(string(b), "\n", " "), " ")
	for _, f := range fields {
		kv := strings.SplitN(f, ":", 2)
		if len(kv) < 2 {
			return fmt.Errorf("malformed key-value pair: %q", f)
		}

		switch kv[0] {
		case "byr":
			year, err := strconv.Atoi(kv[1])
			if err != nil {
				return fmt.Errorf("could not parse %s %q: %w", kv[0], kv[1], err)
			}
			p.BirthYear = year

		case "iyr":
			year, err := strconv.Atoi(kv[1])
			if err != nil {
				return fmt.Errorf("could not parse %s %q: %w", kv[0], kv[1], err)
			}
			p.IssueYear = year

		case "eyr":
			year, err := strconv.Atoi(kv[1])
			if err != nil {
				return fmt.Errorf("could not parse %s %q: %w", kv[0], kv[1], err)
			}
			p.ExpirationYear = year

		case "hgt":
			p.Height = height(kv[1])

		case "hcl":
			p.HairColor = hairColor(kv[1])

		case "ecl":
			p.EyeColor = eyeColor(kv[1])

		case "pid":
			p.PassportID = passportID(kv[1])

		case "cid":
			p.CountryID = kv[1]
		}
	}
	return nil
}

func (p passport) hasRequiredFields() bool {
	return p.BirthYear != 0 &&
		p.IssueYear != 0 &&
		p.ExpirationYear != 0 &&
		p.Height != "" &&
		p.HairColor != "" &&
		p.EyeColor != "" &&
		p.PassportID != ""
}

func (p passport) valid() bool {
	return p.hasRequiredFields() &&
		// byr (Birth Year) - four digits; at least 1920 and at most 2002.
		p.BirthYear >= 1920 && p.BirthYear <= 2002 &&
		// iyr (Issue Year) - four digits; at least 2010 and at most 2020.
		p.IssueYear >= 2010 && p.IssueYear <= 2020 &&
		// eyr (Expiration Year) - four digits; at least 2020 and at most 2030.
		p.ExpirationYear >= 2020 && p.ExpirationYear <= 2030 &&
		// hgt (Height) - a number followed by either cm or in:
		// If cm, the number must be at least 150 and at most 193.
		// If in, the number must be at least 59 and at most 76.
		p.Height.valid() &&
		// hcl (Hair Color) - a # followed by exactly six characters 0-9 or a-f.
		p.HairColor.valid() &&
		// ecl (Eye Color) - exactly one of: amb blu brn gry grn hzl oth.
		p.EyeColor.valid() &&
		// pid (Passport ID) - a nine-digit number, including leading zeroes.
		p.PassportID.valid()

}

type height string

func (h height) valid() bool {
	// hgt (Height) - a number followed by either cm or in:
	// If cm, the number must be at least 150 and at most 193.
	// If in, the number must be at least 59 and at most 76.
	if strings.HasSuffix(string(h), "cm") {
		hInt, _ := strconv.Atoi(string(h[:strings.LastIndex(string(h), "cm")]))
		return hInt >= 150 && hInt <= 193
	} else if strings.HasSuffix(string(h), "in") {
		hInt, _ := strconv.Atoi(string(h[:strings.LastIndex(string(h), "in")]))
		return hInt >= 59 && hInt <= 76
	}
	return false
}

type hairColor string

var hclRegex = regexp.MustCompile(`^#[0-9a-f]{6}$`)

func (c hairColor) valid() bool {
	// hcl (Hair Color) - a # followed by exactly six characters 0-9 or a-f.
	return hclRegex.MatchString(string(c))
}

type eyeColor string

var eclRegex = regexp.MustCompile(`^(amb|blu|brn|gry|grn|hzl|oth)$`)

func (c eyeColor) valid() bool {
	// ecl (Eye Color) - exactly one of: amb blu brn gry grn hzl oth.
	return eclRegex.MatchString(string(c))
}

type passportID string

var pidRegex = regexp.MustCompile(`^[0-9]{9}$`)

func (id passportID) valid() bool {
	// pid (Passport ID) - a nine-digit number, including leading zeroes.
	return pidRegex.MatchString(string(id))
}

func readInput(r io.Reader) ([]passport, error) {
	batch, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}

	var passports []passport

	lines := bytes.Split(batch, []byte("\n\n"))
	for _, l := range lines {
		l = bytes.TrimSpace(l)
		if len(l) == 0 {
			continue
		}
		var p passport
		if err := p.UnmarshalText(l); err != nil {
			return nil, err
		}
		passports = append(passports, p)
	}
	return passports, nil
}
