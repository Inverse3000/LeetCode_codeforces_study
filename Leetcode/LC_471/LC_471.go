package main

import (
	"strconv"
	"strings"
)

var f [205][205]string

func encode(s string) string {
	f = [205][205]string{}
	for l := 1; l <= len(s); l++ {
		for i := 0; i+l-1 < len(s); i++ {
			j := i + l - 1
			f[i][j] = sub(s, i, j)
			if l > 4 {
				for k := i; k < j; k++ {
					split := f[i][k] + f[k+1][j]
					if len(f[i][j]) > len(split) {
						f[i][j] = split
					}
				}
			}
		}
	}
	return f[0][len(s)-1]
}
func sub(s string, i, j int) string {
	if len(s) < 5 {
		return s
	}
	subs := s[i : j+1]
	p := strings.Index((s + s)[1:], subs)
	if p != len(s)-1 {
		cnt := len(s) / p
		return strconv.Itoa(cnt) + "[" + f[i][i+p] + "]"
	}
	return s
}
func main() {

}
