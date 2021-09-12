package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type Node struct {
	a int64
	b int64
}
type Nodes []Node

func (u Nodes) Len() int {
	return len(u)
}
func (u Nodes) Less(i, j int) bool {
	return u[i].b < u[j].b
}
func (u Nodes) Swap(i, j int) {
	u[i], u[j] = u[j], u[i]
}
func main() {
	in := bufio.NewReader(os.Stdin)
	out := os.Stdout
	var n int
	fmt.Fscan(in, &n)
	var a int64
	var b int64
	var i int
	cool := make(Nodes, n+1)
	for i = 1; i <= n; i++ {
		fmt.Fscan(in, &a, &b)
		cool[i].a = a
		cool[i].b = b
	}
	sort.Sort(cool)
	buys := int64(0)
	cnt := int64(0)
	for i := 1; i <= n; i++ {
		if cool[i].a == 0 {
			break
		}
		if cool[i].b <= cnt {
			cnt += cool[i].a
			buys += cool[i].a
			continue
		}
		l := i
		r := n
		ans := 0
		for l <= r {
			mid := (l + r) / 2
			if cool[mid].a > 0 {
				ans = mid
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
		for j := ans; j >= i; j-- {
			fin := int64(-1)
			if cool[j].a < (cool[i].b - cnt) {
				fin = cool[j].a
				cool[j].a = 0
			} else {
				fin = cool[i].b - cnt
				cool[j].a -= (cool[i].b - cnt)
			}
			if cnt < cool[j].b && cnt+fin > cool[j].b {
				buys = (cool[j].b-cnt)*2 + (cnt + fin - cool[j].b)
			} else if cnt >= cool[j].b {
				buys += fin
			} else {
				buys += fin * 2
			}
			cnt += fin
			if cool[i].b == cnt {
				break
			}
		}
		cnt += cool[i].a
		buys += cool[i].a
	}
	fmt.Fprintln(out, buys)
}
