// Generated by copypasta/template/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1985/H2
// https://codeforces.com/problemset/status/1985/problem/H2
func Test_cf1985H2(t *testing.T) {
	testCases := [][2]string{
		{
			`6
1 1
.
4 2
..
#.
#.
.#
3 5
.#.#.
..#..
.#.#.
5 5
#...#
....#
#...#
.....
...##
6 6
.#..#.
#..#..
.#...#
#.#.#.
.#.##.
###..#
6 8
..#....#
.####.#.
###.#..#
.##.#.##
.#.##.##
#..##.#.`,
			`1
7
11
16
22
36`,
		},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, cf1985H2)
}
