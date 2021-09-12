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
	out := os.Stdout
	var t int
	fmt.Fscan(in, &t)
	for c := 1; c <= t; c++ {
		var n int64
		fmt.Fscan(in, &n)
		var pow [100]int64
		pow[0] = 1
		for i := 1; i <= 62; i++ {
			pow[i] = pow[i-1] * 2
		}
		ret := 9999
		for i := 0; i <= 62; i++ {
			d := n
			lst := -1
			ans := 0
			d = n
			sizen := 0
			for d > 0 {
				sizen++
				d /= 10
			}
			sizet := 0
			d = pow[i]
			for d > 0 {
				sizet++
				d /= 10
			}
			p := [305]int64{}
			tempn := [305]int64{}
			d = n
			cnt := sizen
			for d > 0 {
				tempn[cnt] = d % 10
				cnt--
				d /= 10
			}
			d = pow[i]
			cnt = sizet
			for d > 0 {
				p[cnt] = d % 10
				cnt--
				d /= 10
			}
			for k := 1; k <= sizet; k++ {
				index := -1
				for j := 1; j <= sizen; j++ {
					if p[k] == tempn[j] && j > lst {
						index = j
						ans++
						break
					}
				}
				if index == -1 {
					break
				}
				lst = index
			}
			ret = min(ret, sizen-ans+sizet-ans)
		}
		fmt.Fprintln(out, ret)
	}
}
