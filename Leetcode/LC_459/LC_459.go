package main

import (
	"fmt"
	"strings"
)

var primes [10005]int
var vis [10005]bool

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
func repeatedSubstringPattern(s string) bool {
	n := 10000
	cnt := 0
	for i := 2; i <= n; i++ {
		if !vis[i] {
			cnt++
			primes[cnt] = i
		}
		for j := 1; j <= cnt && i*primes[j] < n; j++ {
			vis[i*primes[j]] = true
			if i%primes[j] == 0 {
				break
			}
		}
	}
	m := len(s)
	remain := 1
	ans := 0
	for i := 1; i <= cnt; i++ {
		for m%primes[i] == 0 && m/primes[i] != 1 {
			l := m / primes[i]
			r := primes[i]
			remain = remain * r
			if strings.Repeat(s[:l-1], r) == s {
				ans = max(ans, l)
			}
			if strings.Repeat(s[:remain-1], len(s)/remain) == s {
				ans = max(ans, remain)
			}
			m = m / primes[i]
		}
		if m/primes[i] == 1 {
			break
		}
	}
	fmt.Println(ans)
	if ans != 0 {
		return true
	}
	return false
}
func main() {

}
