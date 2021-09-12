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
	cnt := 0
	now := 0
	lol := [30005]int{}
	for cnt != 1000 {
		now++
		if now%10 == 3 || now%3 == 0 {
			continue
		} else {
			cnt++
			lol[cnt] = now
		}
	}
	fmt.Fscan(in, &t)
	for c := 1; c <= t; c++ {
		var k int
		fmt.Fscan(in, &k)
		fmt.Fprintln(out, lol[k])
	}
}
