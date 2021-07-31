package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1555A(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n int64
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		if n < 6 {
			Fprintln(out, 15)
		} else {
			Fprintln(out, n/6*15+(n%6+1)/2*5)
		}
	}
}

//func main() { CF1555A(os.Stdin, os.Stdout) }
