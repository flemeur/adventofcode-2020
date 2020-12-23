package main

import (
	"fmt"
	"strings"
	"testing"
)

func TestReadInput(t *testing.T) {
	input := `ecl:gry pid:860033327 eyr:2020 hcl:#fffffd
byr:1937 iyr:2017 cid:147 hgt:183cm

iyr:2013 ecl:amb cid:350 eyr:2023 pid:028048884
hcl:#cfa07d byr:1929

hcl:#ae17e1 iyr:2013
eyr:2024
ecl:brn pid:760753108 byr:1931
hgt:179cm

hcl:#cfa07d eyr:2025 pid:166559648
iyr:2011 ecl:brn hgt:59in
`

	passports, err := readInput(strings.NewReader(input))
	if err != nil {
		t.Fatal(err)
	}

	if len(passports) != 4 {
		t.Fatalf("expected 4, got %d", len(passports))
	}
}

func TestPassportRequiredFieldsValidation(t *testing.T) {
	cases := []struct {
		passport passport
		valid    bool
	}{
		{
			passport: passport{
				EyeColor:       "gry",
				PassportID:     "860033327",
				ExpirationYear: 2020,
				HairColor:      "#fffffd",
				BirthYear:      1937,
				IssueYear:      2017,
				CountryID:      "147",
				Height:         "183cm",
			},
			valid: true,
		},
		{
			passport: passport{
				IssueYear:      2013,
				EyeColor:       "amb",
				CountryID:      "350",
				ExpirationYear: 2023,
				PassportID:     "028048884",
				HairColor:      "#cfa07d",
				BirthYear:      1929,
			},
			valid: false,
		},
		{
			passport: passport{
				HairColor:      "#ae17e1",
				IssueYear:      2013,
				ExpirationYear: 2024,
				EyeColor:       "brn",
				PassportID:     "760753108",
				BirthYear:      1931,
				Height:         "179cm",
			},
			valid: true,
		},
		{
			passport: passport{
				HairColor:      "#cfa07d",
				ExpirationYear: 2025,
				PassportID:     "166559648",
				IssueYear:      2011,
				EyeColor:       "brn",
				Height:         "59in",
			},
			valid: false,
		},
	}

	for i, tc := range cases {
		t.Run(fmt.Sprintf("passport %d", i), func(t *testing.T) {
			if tc.passport.hasRequiredFields() != tc.valid {
				t.Fatalf("expected %t, got %t", tc.valid, tc.passport.hasRequiredFields())
			}
		})
	}
}

func TestSolvePuzzle1(t *testing.T) {
	input := `ecl:gry pid:860033327 eyr:2020 hcl:#fffffd
byr:1937 iyr:2017 cid:147 hgt:183cm

iyr:2013 ecl:amb cid:350 eyr:2023 pid:028048884
hcl:#cfa07d byr:1929

hcl:#ae17e1 iyr:2013
eyr:2024
ecl:brn pid:760753108 byr:1931
hgt:179cm

hcl:#cfa07d eyr:2025 pid:166559648
iyr:2011 ecl:brn hgt:59in
`

	passports, err := readInput(strings.NewReader(input))
	if err != nil {
		t.Fatal(err)
	}

	answer := solvePuzzle1(passports)

	if answer != 2 {
		t.Fatalf("expected 2, got %d", answer)
	}
}

func TestSolvePuzzle2(t *testing.T) {
	input := `eyr:1972 cid:100
hcl:#18171d ecl:amb hgt:170 pid:186cm iyr:2018 byr:1926

iyr:2019
hcl:#602927 eyr:1967 hgt:170cm
ecl:grn pid:012533040 byr:1946

hcl:dab227 iyr:2012
ecl:brn hgt:182cm pid:021572410 eyr:2020 byr:1992 cid:277

hgt:59cm ecl:zzz
eyr:2038 hcl:74454a iyr:2023
pid:3556412378 byr:2007

pid:087499704 hgt:74in ecl:grn iyr:2012 eyr:2030 byr:1980
hcl:#623a2f

eyr:2029 ecl:blu cid:129 byr:1989
iyr:2014 pid:896056539 hcl:#a97842 hgt:165cm

hcl:#888785
hgt:164cm byr:2001 iyr:2015 cid:88
pid:545766238 ecl:hzl
eyr:2022

iyr:2010 hgt:158cm hcl:#b6652a ecl:blu byr:1944 eyr:2021 pid:093154719
`

	passports, err := readInput(strings.NewReader(input))
	if err != nil {
		t.Fatal(err)
	}

	answer := solvePuzzle2(passports)

	if answer != 4 {
		t.Fatalf("expected 4, got %d", answer)
	}
}
