package main

import (
	"bufio"
	"fmt"
	"os"
)

var trie [5000005]Node
var ans = 0
var k = 0
var q = []int{}

type Node struct {
	fail int
	cnt  int
	next [30]int
}

func build_trie(id int, s []byte) {
	j := 0
	for i := 0; i < len(s); i++ {
		j = int(s[i] - 'a')
		if trie[id].next[j] == 0 {
			k++
			trie[id].next[j] = k
		}
		id = trie[id].next[j]
	}
	trie[id].cnt++
}
func build_fail(id int) {
	for len(q) != 0 {
		q = q[1:]
	}
	for i := 0; i < 26; i++ {
		j := trie[id].next[i]
		if j != 0 {
			q = append(q, j)
			trie[j].fail = id
		}
	}
	for len(q) != 0 {
		now := q[0]
		q = q[1:]
		for i := 0; i < 26; i++ {
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
func solve_trie(id int, s []byte) {
	j := 0
	for i := 0; i < len(s); i++ {
		j = trie[id].next[int(s[i]-'a')]
		for j != 0 && trie[j].cnt != -1 {
			ans += trie[j].cnt
			trie[j].cnt = -1
			j = trie[j].fail
		}
		id = trie[id].next[s[i]-'a']
	}
}
func main() {
	in := bufio.NewReader(os.Stdin)
	out := os.Stdout
	var n int
	fmt.Fscan(in, &n)
	for i := 0; i < n; i++ {
		var c string
		fmt.Fscan(in, &c)
		build_trie(0, []byte(c))
	}
	build_fail(0)
	var s string
	fmt.Fscan(in, &s)
	solve_trie(0, []byte(s))
	fmt.Fprintln(out, ans)
}
