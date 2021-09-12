package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := os.Stdout
	var t int
	fmt.Fscan(in, &t)
	for c := 0; c < t; c++ {
		var n, k int
		fmt.Fscan(in, &n, &k)
		l := []int{}
		or := []int{}
		for i := 1; i <= n; i++ {
			var m int
			fmt.Fscan(in, &m)
			l = append(l, m)
			or = append(or, m)
		}
		sort.Ints(l)
		mp := map[int]int{}
		for i := 0; i < n-1; i++ {
			mp[l[i]] = l[i+1]
		}
		cnt := 1
		for i := 0; i < n-1; i++ {
			if mp[or[i]] != or[i+1] || (or[i] == l[n-1]) {
				cnt++
			}
		}
		if cnt > k {
			fmt.Fprintln(out, "No")
		} else {
			fmt.Fprintln(out, "Yes")
		}
	}
}
