package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func Test(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
7 3
2 6 4 3 6 8 9
outputCopy
16`
	testutil.AssertEqualCase(t, rawText, 0, run)
}
