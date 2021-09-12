package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

type Node struct {
	fail int
	cnt  int
	next [30]int
}

var trie [5000005]Node
var ans = 0
var k = 0
var q = []int{}
var dp [3005][5005]int

func build_trie(id int, s []byte) {
	j := 0
	for i := 0; i < len(s); i++ {
		j = int(s[i] - 'A')
		if trie[id].next[j] == 0 {
			k++
			trie[id].next[j] = k
		}
		id = trie[id].next[j]
	}
	trie[id].cnt++
}

func build_fail(id int) {
	for i := 0; i < 3; i++ {
		j := trie[id].next[i]
		if j != 0 {
			q = append(q, j)
			trie[j].fail = id
		}
	}
	for len(q) != 0 {
		now := q[0]
		q = q[1:]
		trie[now].cnt += trie[trie[now].fail].cnt
		for i := 0; i < 3; i++ {
			j := trie[now].next[i]
			if j == 0 {
				trie[now].next[i] = trie[trie[now].fail].next[i]
				continue
			}
			trie[j].fail = trie[trie[now].fail].next[i]
			q = append(q, j)
		}
	}
}
func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
func main() {
	in := bufio.NewReader(os.Stdin)
	out := os.Stdout
	var n, m int
	fmt.Fscan(in, &n, &m)
	for i := 1; i <= n; i++ {
		var s string
		fmt.Fscan(in, &s)
		build_trie(0, []byte(s))
	}
	build_fail(0)
	for i := 0; i <= m; i++ {
		for j := 1; j <= k; j++ {
			dp[i][j] = -math.MaxInt32
		}
	}
	for i := 1; i <= m; i++ {
		for j := 0; j <= k; j++ {
			for v := 0; v < 3; v++ {
				dp[i][trie[j].next[v]] = max(dp[i][trie[j].next[v]], dp[i-1][j]+trie[trie[j].next[v]].cnt)
			}
		}
	}
	for i := 0; i <= k; i++ {
		ans = max(ans, dp[m][i])
	}
	fmt.Fprintln(out, ans)
}
