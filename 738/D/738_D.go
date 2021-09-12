package main

import (
	"bufio"
	"fmt"
	"os"
)

var f1 [3005]int
var f2 [3005]int

func findf1(x int) int {
	if x != f1[x] {
		f1[x] = findf1(f1[x])
	}
	return f1[x]
}
func findf2(x int) int {
	if x != f2[x] {
		f2[x] = findf2(f2[x])
	}
	return f2[x]
}
func main() {
	in := bufio.NewReader(os.Stdin)
	out := os.Stdout
	var n, m1, m2 int
	fmt.Fscan(in, &n, &m1, &m2)
	var e1 [3005][3005]int
	var e2 [3005][3005]int
	for i := 1; i <= n; i++ {
		f1[i] = i
		f2[i] = i
	}
	for i := 1; i <= m1; i++ {
		var u, v int
		fmt.Fscan(in, &u, &v)
		e1[u][v] = 1
		e1[v][u] = 1
		f1u := findf1(u)
		f1v := findf1(v)
		if f1u != f1v {
			f1[f1u] = f1v
		}
	}
	for i := 1; i <= m2; i++ {
		var u, v int
		fmt.Fscan(in, &u, &v)
		e2[u][v] = 1
		e2[v][u] = 1
		f2u := findf2(u)
		f2v := findf2(v)
		if f2u != f2v {
			f2[f2u] = f2v
		}
	}
	ans := 0
	cool := [][]int{}
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			if e1[i][j] != 0 || e2[i][j] != 0 {
				continue
			}
			if findf1(i) != findf1(j) && findf2(j) != findf2(i) {
				f1i := findf1(i)
				f1j := findf1(j)
				f1[f1i] = f1j
				f2i := findf2(i)
				f2j := findf2(j)
				f2[f2i] = f2j
				ans++
				e1[i][j] = 1
				e2[i][j] = 1
				e1[j][i] = 1
				e2[j][i] = 1
				cool = append(cool, []int{i, j})
			}
		}
	}
	fmt.Fprintln(out, ans)
	for i := 1; i <= ans; i++ {
		fmt.Fprintln(out, cool[i-1][0], cool[i-1][1])
	}
}
