package main

import (
	"strings"
	"testing"
)

func TestSolvePuzzle1(t *testing.T) {
	cases := []struct {
		input  string
		answer int
	}{
		{`
16
10
15
5
1
11
7
19
6
12
4
`, 7 * 5},
		{`
28
33
18
42
31
14
46
20
48
47
24
23
49
45
19
38
39
11
1
32
25
35
8
17
7
9
4
2
34
10
3
`, 22 * 10},
	}

	for _, tc := range cases {
		t.Run("", func(t *testing.T) {
			adapters, err := readInput(strings.NewReader(tc.input))
			if err != nil {
				t.Fatal(err)
			}

			answer := solvePuzzle1(adapters)
			if answer != tc.answer {
				t.Fatalf("expected answer %d, got %d", tc.answer, answer)
			}
		})
	}
}

func TestSolvePuzzle2(t *testing.T) {
	cases := []struct {
		input  string
		answer int
	}{

		{`
16
10
15
5
1
11
7
19
6
12
4
`, 8},

		{`
28
33
18
42
31
14
46
20
48
47
24
23
49
45
19
38
39
11
1
32
25
35
8
17
7
9
4
2
34
10
3
`, 19208},
	}

	for _, tc := range cases {
		t.Run("", func(t *testing.T) {
			adapters, err := readInput(strings.NewReader(tc.input))
			if err != nil {
				t.Fatal(err)
			}

			answer := solvePuzzle2(adapters)
			if answer != tc.answer {
				t.Fatalf("expected answer %d, got %d", tc.answer, answer)
			}
		})
	}
}
