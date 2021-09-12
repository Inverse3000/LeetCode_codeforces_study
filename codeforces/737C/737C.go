package main

import (
	"fmt"
	"os"
)

type ll int64

func main() {
	in := os.Stdin
	out := os.Stdout
	const mod ll = 1000000007
	var t ll
	fmt.Fscan(in, &t)
	for ; t > 0; t-- {
		var n, k int
		fmt.Fscan(in, &n, &k)
		pow := [300005]ll{}
		pow[1] = 1
		for i := 2; i <= n; i++ {
			pow[i] = (pow[i-1] * 2) % mod
		}
		dp := [300005][2]ll{}
		var m ll
		m = 1
		dp[0][0] = m
		for i := 1; i <= k; i++ {
			if n%2 == 1 {
				dp[i][0] = (dp[i-1][0] * (pow[n] + 1)) % mod
				dp[i][1] = (dp[i-1][1] * pow[n] * 2) % mod
			} else if n%2 == 0 {
				dp[i][0] = (dp[i-1][0] * (pow[n] - 1)) % mod
				dp[i][1] = (dp[i-1][0] + dp[i-1][1]*pow[n]*2) % mod
			}
		}
		fmt.Fprintln(out, (dp[k][0]+dp[k][1])%mod)
	}
}
