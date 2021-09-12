package main

import (
	"bufio"
	"fmt"
	"os"
)

var c [3000005]int

func lowbit(x int) int {
	return x & -x
}
func add(x, val int) {
	for i := x; i <= 1000006; i += lowbit(i) {
		c[i] += val
	}
}
func query(x int) (res int) {
	for i := x; i >= 1; i -= lowbit(i) {
		res += c[i]
	}
	return
}
func main() {
	in := bufio.NewReader(os.Stdin)
	out := os.Stdout
	var n int
	fmt.Fscan(in, &n)
	for i := 1; i <= n; i++ {
		c[i] = 0
	}
	var a [3000005]int
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i])
	}
	mp1 := map[int]int{}
	var f [3000005]int
	for i := 1; i <= n; i++ {
		mp1[a[i]]++
		f[i] = mp1[a[i]]
	}
	mp2 := map[int]int{}
	var l [3000005]int
	for i := n; i >= 1; i-- {
		mp2[a[i]]++
		l[i] = mp2[a[i]]
	}
	ans := int64(0)
	for i := n; i >= 1; i-- {
		//	fmt.Println(f[i]-1,query(f[i]-1))
		ans += int64(query(f[i] - 1))
		add(l[i], 1)
	}
	fmt.Fprintln(out, ans)
}
