package main

import (
	"bufio"
	"fmt"
	"os"
)

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	var t int
	fmt.Fscan(in, &t)
	for c := 1; c <= t; c++ {
		var n int
		var k int
		fmt.Fscan(in, &n, &k)
		p := make([]int, n+5)
		l := make([]int, n+5)
		v := make([]int, n+5)
		f := make([]int, n+5)
		b := make([]int, n+5)
		for i := 1; i <= k; i++ {
			fmt.Fscan(in, &p[i])
		}
		for i := 1; i <= k; i++ {
			fmt.Fscan(in, &v[i])
		}
		for i := 1; i <= k; i++ {
			l[p[i]] = v[i]
		}
		for i := 1; i <= n; i++ {
			if l[i] != 0 && f[i-1] == 0 {
				f[i] = l[i]
				continue
			}
			if l[i] != 0 && f[i-1] != 0 {
				f[i] = min(f[i-1]+1, l[i])
				continue
			}
			if f[i-1] != 0 {
				f[i] = f[i-1] + 1
			}
		}
		for i := n; i >= 1; i-- {
			if l[i] != 0 && b[i+1] == 0 {
				b[i] = l[i]
				continue
			}
			if l[i] != 0 && b[i+1] != 0 {
				b[i] = min(b[i+1]+1, l[i])
				continue
			}
			if b[i+1] != 0 {
				b[i] = b[i+1] + 1
			}
		}
		for i := 1; i <= n; i++ {
			var ans int
			if f[i] == 0 {
				ans = b[i]
			}
			if b[i] == 0 {
				ans = f[i]
			}
			if f[i] != 0 && b[i] != 0 {
				ans = min(f[i], b[i])
			}
			fmt.Fprintf(out, "%d ", ans)
		}
		fmt.Fprintf(out, "\n")
	}
}
