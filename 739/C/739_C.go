package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := os.Stdout
	var t int
	fmt.Fscan(in, &t)
	for c := 1; c <= t; c++ {
		var k int
		fmt.Fscan(in, &k)
		now := 1
		cur := 1
		r := 1
		c := 1
		for k > now {
			now += cur + 2
			cur = cur + 2
			r++
		}
		nums := now - k
		flag := false
		for i := 1; i <= nums; i++ {
			if flag == false {
				c++
				if c == r {
					flag = true
				}
			} else if flag {
				r--
			}
		}
		fmt.Fprintln(out, r, c)
	}
}
