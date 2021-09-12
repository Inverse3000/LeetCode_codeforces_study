package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := os.Stdout
	var n, m int
	fmt.Fscan(in, &n, &m)
	var b []byte
	fmt.Fscan(in, &b)
	var f1 []byte
	f1 = []byte("abc")

	var cst [7][300005]int
	for i := 1; i <= 6; i++ {
		if i == 1 {
			f1 = []byte("abc")
		}
		if i == 2 {
			f1 = []byte("cba")
		}
		if i == 3 {
			f1 = []byte("bca")
		}
		if i == 4 {
			f1 = []byte("cab")
		}
		if i == 5 {
			f1 = []byte("bac")
		}
		if i == 6 {
			f1 = []byte("acb")
		}
		for j := 0; j < n; j++ {
			v := 0
			if f1[j%3] != b[j] {
				v = 1
			}
			cst[i][j+1] = cst[i][j] + v
		}
	}
	for i := 1; i <= m; i++ {
		var l, r int
		fmt.Fscan(in, &l, &r)
		ans := 999999999
		for j := 1; j <= 6; j++ {
			if ans > cst[j][r]-cst[j][l-1] {
				ans = cst[j][r] - cst[j][l-1]
			}
		}
		fmt.Fprintln(out, ans)
	}
}
