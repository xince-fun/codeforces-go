// Code generated by copypasta/template/leetcode/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/leetcode/testutil"
	testutil2 "github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func Test_d(t *testing.T) {
	targetCaseNum := -1
	if err := testutil.RunLeetCodeFuncWithFile(t, magnificentSets, "d.txt", targetCaseNum); err != nil {
		t.Fatal(err)
	}
}

// https://leetcode.cn/contest/weekly-contest-322/problems/divide-nodes-into-the-maximum-number-of-groups/

func TestCompareInf(t *testing.T) {
	testutil.DebugTLE = 0

	inputGenerator := func() (n int, edges [][]int) {
		//return
		rg := testutil2.NewRandGenerator()
		n = rg.Int(1, 5)
		m := rg.Int(n-1, n*(n-1)/2)
		edges = testutil2.TransEdges(rg.GraphEdges(n, m, 1, false))
		return
	}

	testutil.CompareInf(t, inputGenerator, magnificentSets, magnificentSetsWA)
}

func magnificentSetsWA(n int, edges [][]int) (ans int) {
	g := make([][]int, n)
	for _, e := range edges {
		x, y := e[0]-1, e[1]-1
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}

	time := make([]int, n) // 充当 vis 数组的作用（避免在 BFS 内部重复创建 vis 数组）
	clock := 0
	bfs := func(start int) (depth, u int) { // 返回从 start 出发的最大深度
		clock++
		time[start] = clock
		for q := []int{start}; len(q) > 0; depth++ {
			tmp := q
			q = nil
			for _, x := range tmp {
				u = x
				for _, y := range g[x] {
					if time[y] != clock { // 没有在同一次 BFS 中访问过
						time[y] = clock
						q = append(q, y)
					}
				}
			}
		}
		return
	}

	colors := make([]int8, n)
	var isBipartite func(int, int8) bool
	isBipartite = func(x int, c int8) bool { // 二分图判定，原理见视频讲解
		colors[x] = c
		for _, y := range g[x] {
			if colors[y] == c || colors[y] == 0 && !isBipartite(y, -c) {
				return false
			}
		}
		return true
	}
	for i, c := range colors {
		if c == 0 {
			if !isBipartite(i, 1) { // 如果不是二分图（有奇环），则无法分组
				return -1
			}
			_, u := bfs(i)
			res, _ := bfs(u)
			ans += res
		}
	}
	return
}