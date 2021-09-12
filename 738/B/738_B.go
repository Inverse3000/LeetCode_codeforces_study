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
		var n int
		fmt.Fscan(in, &n)
		var temp string
		cnt := 0
		fmt.Fscan(in, &temp)
		s := []byte(temp)
		for i := 0; i < n; i++ {
			if s[i] == 'R' || s[i] == 'B' {
				cnt++
			}
		}
		s = append(s, 0)
		s = append(s, 0)
		for i := n - 1; i >= 0; i-- {
			s[i+1] = s[i]
		}
		s[0] = 0
		if cnt == 0 {
			for i := 1; i <= n; i++ {
				if i%2 == 1 {
					s[i] = 'B'
				} else {
					s[i] = 'R'
				}
			}
			fmt.Fprintln(out, string(s[1:len(s)-1]))
			continue
		}
		for cnt != len(s)-2 {
			for i := 1; i <= n; i++ {
				if s[i-1] != 'R' && s[i-1] != 'B' && s[i+1] != 'R' && s[i+1] != 'B' {
					continue
				} else if s[i] == '?' {
					if (s[i-1] == 'R' && s[i+1] == 'B') || (s[i-1] == 'B' && s[i+1] == 'R') {
						s[i] = 'R'
					}
					if s[i-1] == 'R' && s[i+1] != 'R' && s[i+1] != 'B' {
						s[i] = 'B'
					}
					if s[i-1] == 'B' && s[i+1] != 'R' && s[i+1] != 'B' {
						s[i] = 'R'
					}
					if s[i+1] == 'R' && s[i-1] != 'R' && s[i-1] != 'B' {
						s[i] = 'B'
					}
					if s[i+1] == 'B' && s[i-1] != 'R' && s[i-1] != 'B' {
						s[i] = 'R'
					}
					if s[i-1] == 'B' && s[i+1] == 'B' {
						s[i] = 'R'
					}
					if s[i-1] == 'R' && s[i+1] == 'R' {
						s[i] = 'B'
					}
					cnt++
				}
			}
		}
		fmt.Fprintln(out, string(s[1:len(s)-1]))
	}
}
