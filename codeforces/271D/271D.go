package main

import (
	"bufio"
	"fmt"
	"os"
)

var f [300005][2]int
var pre [300005]int

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	var t int
	var k int
	f := [300005][2]int{}
	pre := [300005]int{}
	if k == 1 {
		f[1][1] = 1
	}
	f[1][0] = 1
	fmt.Fscan(in, &t, &k)
	f[k][1] = 1
	for i := 2; i <= 100000; i++ {
		f[i][0] = (f[i-1][0]%1000000007 + f[i-1][1]%1000000007) % 1000000007
		if i-k >= 1 {
			f[i][1] = (f[i-k][0]%1000000007 + f[i-k][1]%1000000007) % 1000000007
		}
	}
	for i := 1; i <= 100000; i++ {
		pre[i] = (pre[i-1]%1000000007 + (f[i][0]+f[i][1])%1000000007) % 1000000007
	}
	for c := 1; c <= t; c++ {
		var a int
		var b int
		fmt.Fscan(in, &a, &b)
		ans := 0
		ans = (pre[b] - pre[a-1] + 1000000007) % 1000000007
		fmt.Fprintln(out, ans)
	}
}
