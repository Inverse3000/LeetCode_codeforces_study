package main

import (
	"bufio"
	"fmt"
	"os"
)

func max(x, y int64) int64 {
	if x > y {
		return x
	}
	return y
}
func min(x, y int64) int64 {
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
		var n int64
		var k int64
		fmt.Fscan(in, &n, &k)
		var a [100005]int64
		var i int64
		for i = 1; i <= n; i++ {
			fmt.Fscan(in, &a[i])
		}
		st := max(n-2*k, 1)
		var ans int64
		ans = -99999999
		for i = st; i <= n; i++ {
			for j := i + 1; j <= n; j++ {
				ans = max(ans, i*j-(a[i]|a[j])*k)
			}
		}
		fmt.Fprintln(out, ans)
	}
}
