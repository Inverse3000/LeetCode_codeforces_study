package main

import (
	"bufio"
	"fmt"
	"os"
)

type Node struct {
	index int
	h     int
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
func main() {
	in := bufio.NewReader(os.Stdin)
	out := os.Stdout
	var n int
	fmt.Fscan(in, &n)
	var a [300005]int
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i])
	}
	ans := 0
	st := []Node{}
	for i := 1; i <= n; i++ {
		if i == 1 {
			st = append(st, Node{i, a[i]})
			continue
		}
		lst := st[len(st)-1].index
		for len(st) >= 1 && st[len(st)-1].h > a[i] {
			if len(st) == 1 {
				ans = max(ans, min(st[len(st)-1].h, lst))
				st = st[:len(st)-1]
				continue
			}
			if lst-st[len(st)-2].index >= st[len(st)-1].h {
				ans = max(ans, st[len(st)-1].h)
			} else {
				ans = max(ans, lst-st[len(st)-2].index)
			}
			st = st[:len(st)-1]
		}
		st = append(st, Node{i, a[i]})
	}
	lst := st[len(st)-1].index
	for len(st) >= 1 {
		if len(st) == 1 {
			ans = max(ans, min(st[len(st)-1].h, lst))
			st = st[:len(st)-1]
			continue
		}
		if lst-st[len(st)-2].index >= st[len(st)-1].h {
			ans = max(ans, st[len(st)-1].h)
		} else {
			ans = max(ans, lst-st[len(st)-2].index)
		}
		st = st[:len(st)-1]
	}
	fmt.Fprintln(out, ans)
}
