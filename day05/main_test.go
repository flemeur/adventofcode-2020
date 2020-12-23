package main

import "testing"

func TestParseSeat(t *testing.T) {
	cases := []struct {
		input  string
		row    int
		column int
		id     int
	}{
		{"FBFBBFFRLR", 44, 5, 357},
		{"BFFFBBFRRR", 70, 7, 567},
		{"FFFBBBFRRR", 14, 7, 119},
		{"BBFFBBFRLL", 102, 4, 820},
	}

	for _, tc := range cases {
		t.Run("", func(t *testing.T) {
			row, column, id, err := parseSeat(tc.input)
			if err != nil {
				t.Fatal(err)
			}

			if row != tc.row {
				t.Fatalf("expected row %d, got %d", tc.row, row)
			}
			if column != tc.column {
				t.Fatalf("expected column %d, got %d", tc.column, column)
			}
			if id != tc.id {
				t.Fatalf("expected id %d, got %d", tc.id, id)
			}
		})
	}
}
